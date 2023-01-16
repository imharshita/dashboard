/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package externalcluster

import (
	"context"
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"

	clusterv1alpha1 "github.com/kubermatic/machine-controller/pkg/apis/cluster/v1alpha1"
	apiv2 "k8c.io/dashboard/v2/pkg/api/v2"
	handlercommon "k8c.io/dashboard/v2/pkg/handler/common"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/provider"
	kubeonev1beta2 "k8c.io/kubeone/pkg/apis/kubeone/v1beta2"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"
	ksemver "k8c.io/kubermatic/v2/pkg/semver"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func importKubeOneCluster(ctx context.Context, userInfoGetter func(ctx context.Context, projectID string) (*provider.UserInfo, error), project *kubermaticv1.Project, cloud *apiv2.ExternalClusterCloudSpec, clusterProvider provider.ExternalClusterProvider, privilegedClusterProvider provider.PrivilegedExternalClusterProvider) (*kubermaticv1.ExternalCluster, error) {
	kubeOneCluster, err := DecodeManifestFromKubeOneReq(cloud.KubeOne.Manifest)
	if err != nil {
		return nil, err
	}

	isImported := resources.ExternalClusterIsImportedTrue
	newCluster := genExternalCluster(kubeOneCluster.Name, project.Name, isImported)

	version, err := ksemver.NewSemver(kubeOneCluster.Versions.Kubernetes)
	if err != nil {
		return nil, err
	}
	newCluster.Spec.Version = *version

	if kubeOneCluster.ContainerRuntime.Docker != nil {
		newCluster.Spec.ContainerRuntime = resources.ContainerRuntimeDocker
	} else if kubeOneCluster.ContainerRuntime.Containerd != nil {
		newCluster.Spec.ContainerRuntime = resources.ContainerRuntimeContainerd
	}

	newCluster.Spec.CloudSpec = kubermaticv1.ExternalClusterCloudSpec{
		KubeOne: &kubermaticv1.ExternalClusterKubeOneCloudSpec{},
	}
	newCluster = setKubeOneProvider(kubeOneCluster, *newCluster)

	kubermaticNamespace := resources.KubermaticNamespace
	err = clusterProvider.CreateOrUpdateKubeOneSSHSecret(ctx, kubermaticNamespace, cloud.KubeOne.SSHKey, newCluster)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}
	err = clusterProvider.CreateOrUpdateKubeOneManifestSecret(ctx, kubermaticNamespace, cloud.KubeOne.Manifest, newCluster)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}

	err = clusterProvider.CreateOrUpdateKubeOneCredentialSecret(ctx, kubermaticNamespace, *cloud.KubeOne.CloudSpec, newCluster)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}

	newCluster.Status.Condition.Phase = kubermaticv1.ExternalClusterPhaseProvisioning
	return createNewCluster(ctx, userInfoGetter, clusterProvider, privilegedClusterProvider, newCluster, project)
}

func patchKubeOneCluster(ctx context.Context,
	existingCluster *kubermaticv1.ExternalCluster,
	patchData json.RawMessage,
	secretKeySelector provider.SecretKeySelectorValueFunc,
	clusterProvider provider.ExternalClusterProvider,
	masterClient ctrlruntimeclient.Client) (*apiv2.ExternalCluster, error) {
	operation := existingCluster.Status.Condition.Phase
	if operation == kubermaticv1.ExternalClusterPhaseReconciling {
		return nil, utilerrors.NewBadRequest("Operation is not allowed: Another operation: (%s) is in progress, please wait for it to finish before starting a new operation.", operation)
	}

	version, err := clusterProvider.GetVersion(ctx, masterClient, existingCluster)
	if err != nil {
		return nil, err
	}
	currentVersion := *version
	currentContainerRuntime := existingCluster.Spec.ContainerRuntime

	existingClusterJSON, err := json.Marshal(existingCluster)
	if err != nil {
		return nil, utilerrors.NewBadRequest("cannot decode existing kubeone cluster: %v", err)
	}

	patchedClusterJSON, err := jsonpatch.MergePatch(existingClusterJSON, patchData)
	if err != nil {
		return nil, utilerrors.NewBadRequest("cannot patch kubeone cluster: %v", err)
	}

	var patchedCluster *kubermaticv1.ExternalCluster
	err = json.Unmarshal(patchedClusterJSON, &patchedCluster)
	if err != nil {
		return nil, utilerrors.NewBadRequest("cannot decode patched settings: %v", err)
	}

	desiredVersion := patchedCluster.Spec.Version
	desiredContainerRuntime := patchedCluster.Spec.ContainerRuntime
	switch {
	case currentVersion != desiredVersion && currentContainerRuntime != desiredContainerRuntime:
		return nil, utilerrors.NewBadRequest("Operation not supported: cannot change version and containerRuntime at the same time")
	case currentVersion != desiredVersion:
		mdList, err := getKubeOneMachineDeployments(ctx, masterClient, existingCluster, clusterProvider)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		mdsVersion := make(map[string]string)
		for _, md := range mdList.Items {
			mdsVersion[md.Name] = md.Spec.Template.Spec.Versions.Kubelet
		}
		for mdName, mdVersion := range mdsVersion {
			mdVersion, err := ksemver.NewSemver(mdVersion)
			if err != nil {
				return nil, err
			}
			if currentVersion.Semver().Minor()-mdVersion.Semver().Minor() >= 1 {
				return nil, fmt.Errorf("MachineDeployment %s must be updated to match cluster version %v before updating cluster version", mdName, currentVersion)
			}
		}
	case currentContainerRuntime != desiredContainerRuntime:
		// currently only migration to containerd is supported
		if !sets.NewString("containerd").Has(desiredContainerRuntime) {
			return nil, fmt.Errorf("Operation not supported: Only migration from docker to containerd is supported, desiredContainerRuntime: %s", desiredContainerRuntime)
		}
	}

	if err := masterClient.Update(ctx, patchedCluster); err != nil {
		return nil, fmt.Errorf("failed to update kubeone external cluster %s: %w", existingCluster.Name, err)
	}

	return convertClusterToAPI(patchedCluster), nil
}

func getKubeOneMachineDeployment(ctx context.Context, masterClient ctrlruntimeclient.Client, mdName string, cluster *kubermaticv1.ExternalCluster, clusterProvider provider.ExternalClusterProvider) (*clusterv1alpha1.MachineDeployment, error) {
	machineDeployment := &clusterv1alpha1.MachineDeployment{}
	userClusterClient, err := clusterProvider.GetClient(ctx, masterClient, cluster)
	if err != nil {
		return nil, err
	}
	if err := userClusterClient.Get(ctx, types.NamespacedName{Name: mdName, Namespace: metav1.NamespaceSystem}, machineDeployment); err != nil && !meta.IsNoMatchError(err) {
		return nil, fmt.Errorf("failed to get MachineDeployment: %w", err)
	}
	return machineDeployment, nil
}

func getKubeOneMachineDeployments(ctx context.Context, masterClient ctrlruntimeclient.Client, cluster *kubermaticv1.ExternalCluster, clusterProvider provider.ExternalClusterProvider) (*clusterv1alpha1.MachineDeploymentList, error) {
	mdList := &clusterv1alpha1.MachineDeploymentList{}
	userClusterClient, err := clusterProvider.GetClient(ctx, masterClient, cluster)
	if err != nil {
		return nil, err
	}
	if err := userClusterClient.List(ctx, mdList); err != nil {
		return nil, fmt.Errorf("failed to list MachineDeployment: %w", err)
	}
	return mdList, nil
}

func patchKubeOneMachineDeployment(ctx context.Context, masterClient ctrlruntimeclient.Client, machineDeployment *clusterv1alpha1.MachineDeployment, oldmd, newmd *apiv2.ExternalClusterMachineDeployment, cluster *kubermaticv1.ExternalCluster, clusterProvider provider.ExternalClusterProvider) (*apiv2.ExternalClusterMachineDeployment, error) {
	currentVersion := oldmd.NodeDeployment.Spec.Template.Versions.Kubelet
	desiredVersion := newmd.NodeDeployment.Spec.Template.Versions.Kubelet
	if desiredVersion != currentVersion {
		machineDeployment.Spec.Template.Spec.Versions.Kubelet = desiredVersion
		userClusterClient, err := clusterProvider.GetClient(ctx, masterClient, cluster)
		if err != nil {
			return nil, err
		}
		if err := userClusterClient.Update(ctx, machineDeployment); err != nil && !meta.IsNoMatchError(err) {
			return nil, fmt.Errorf("failed to update MachineDeployment: %w", err)
		}
		return newmd, nil
	}

	currentReplicas := oldmd.NodeDeployment.Spec.Replicas
	desiredReplicas := newmd.NodeDeployment.Spec.Replicas
	if desiredReplicas != currentReplicas {
		machineDeployment.Spec.Replicas = &desiredReplicas
		userClusterClient, err := clusterProvider.GetClient(ctx, masterClient, cluster)
		if err != nil {
			return nil, err
		}
		if err := userClusterClient.Update(ctx, machineDeployment); err != nil && !meta.IsNoMatchError(err) {
			return nil, fmt.Errorf("failed to update MachineDeployment: %w", err)
		}
		return newmd, nil
	}

	return oldmd, nil
}

func getKubeOneAPIMachineDeployment(ctx context.Context,
	masterClient ctrlruntimeclient.Client,
	mdName string,
	cluster *kubermaticv1.ExternalCluster,
	clusterProvider provider.ExternalClusterProvider) (*apiv2.ExternalClusterMachineDeployment, error) {
	md, err := getKubeOneMachineDeployment(ctx, masterClient, mdName, cluster, clusterProvider)
	if err != nil {
		return nil, err
	}
	nd, err := handlercommon.OutputMachineDeployment(md)
	if err != nil {
		return nil, err
	}

	return &apiv2.ExternalClusterMachineDeployment{NodeDeployment: *nd}, nil
}

func setKubeOneProvider(kubeOneCluster *kubeonev1beta2.KubeOneCluster, newCluster kubermaticv1.ExternalCluster) *kubermaticv1.ExternalCluster {
	switch {
	case kubeOneCluster.CloudProvider.AWS != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneAWS
	case kubeOneCluster.CloudProvider.GCE != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneGCP
	case kubeOneCluster.CloudProvider.Azure != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneAzure
	case kubeOneCluster.CloudProvider.DigitalOcean != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneDigitalOcean
	case kubeOneCluster.CloudProvider.Hetzner != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneHetzner
	case kubeOneCluster.CloudProvider.Nutanix != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneNutanix
	case kubeOneCluster.CloudProvider.Openstack != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneOpenStack
	case kubeOneCluster.CloudProvider.EquinixMetal != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneEquinix
	case kubeOneCluster.CloudProvider.Vsphere != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneVSphere
	case kubeOneCluster.CloudProvider.VMwareCloudDirector != nil:
		newCluster.Spec.CloudSpec.KubeOne.ProviderName = resources.KubeOneVMwareCloudDirector
	}

	return &newCluster
}

func getKubeOneAPIMachineDeployments(ctx context.Context,
	masterClient ctrlruntimeclient.Client,
	cluster *kubermaticv1.ExternalCluster,
	clusterProvider provider.ExternalClusterProvider) ([]apiv2.ExternalClusterMachineDeployment, error) {
	mdList, err := getKubeOneMachineDeployments(ctx, masterClient, cluster, clusterProvider)
	nodeDeployments := make([]apiv2.ExternalClusterMachineDeployment, 0, len(mdList.Items))
	if err != nil {
		return nil, err
	}

	for _, md := range mdList.Items {
		nd, err := handlercommon.OutputMachineDeployment(&md)
		if err != nil {
			return nil, fmt.Errorf("failed to output machine deployment %s: %w", md.Name, err)
		}
		nodeDeployments = append(nodeDeployments, apiv2.ExternalClusterMachineDeployment{NodeDeployment: *nd})
	}

	return nodeDeployments, nil
}

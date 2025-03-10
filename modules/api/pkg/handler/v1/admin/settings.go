/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package admin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/go-kit/kit/endpoint"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/provider"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"
)

// KubermaticSettingsEndpoint returns global settings.
func KubermaticSettingsEndpoint(settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		globalSettings, err := settingsProvider.GetGlobalSettings(ctx)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return apiv1.GlobalSettings(globalSettings.Spec), nil
	}
}

// KubermaticCustomLinksEndpoint returns custom links.
func KubermaticCustomLinksEndpoint(settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		globalSettings, err := settingsProvider.GetGlobalSettings(ctx)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return apiv1.GlobalCustomLinks(globalSettings.Spec.CustomLinks), nil
	}
}

// UpdateKubermaticSettingsEndpoint updates global settings.
func UpdateKubermaticSettingsEndpoint(userInfoGetter provider.UserInfoGetter, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(patchKubermaticSettingsReq)
		userInfo, err := userInfoGetter(ctx, "")
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		existingGlobalSettings, err := settingsProvider.GetGlobalSettings(ctx)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		existingGlobalSettingsSpecJSON, err := json.Marshal(existingGlobalSettings.Spec)
		if err != nil {
			return nil, utilerrors.NewBadRequest("cannot decode existing settings: %v", err)
		}

		patchedGlobalSettingsSpecJSON, err := jsonpatch.MergePatch(existingGlobalSettingsSpecJSON, req.Patch)
		if err != nil {
			return nil, utilerrors.NewBadRequest("cannot patch global settings: %v", err)
		}
		var patchedGlobalSettingsSpec *kubermaticv1.SettingSpec
		err = json.Unmarshal(patchedGlobalSettingsSpecJSON, &patchedGlobalSettingsSpec)
		if err != nil {
			return nil, utilerrors.NewBadRequest("cannot decode patched settings: %v", err)
		}

		existingGlobalSettings.Spec = *patchedGlobalSettingsSpec
		globalSettings, err := settingsProvider.UpdateGlobalSettings(ctx, userInfo, existingGlobalSettings)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return apiv1.GlobalSettings(globalSettings.Spec), nil
	}
}

// patchKubermaticSettingsReq defines HTTP request for patchKubermaticSettings endpoint
// swagger:parameters patchKubermaticSettings
type patchKubermaticSettingsReq struct {
	// in: body
	Patch json.RawMessage
}

func DecodePatchKubermaticSettingsReq(c context.Context, r *http.Request) (interface{}, error) {
	var req patchKubermaticSettingsReq
	var err error

	if req.Patch, err = io.ReadAll(r.Body); err != nil {
		return nil, err
	}

	return req, nil
}

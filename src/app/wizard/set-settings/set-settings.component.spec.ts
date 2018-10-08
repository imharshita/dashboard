import { SharedModule } from '../../shared/shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatTabsModule } from '@angular/material';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { SetSettingsComponent } from './set-settings.component';
import { WizardService } from '../../core/services/wizard/wizard.service';
import { fakeDigitaloceanCluster } from '../../testing/fake-data/cluster.fake';
import { ClusterProviderSettingsComponent } from './provider-settings/provider-settings.component';
import { AddNodeComponent } from '../../add-node/add-node.component';
import { ClusterSSHKeysComponent } from './ssh-keys/cluster-ssh-keys.component';
import { DigitaloceanClusterSettingsComponent } from './provider-settings/digitalocean/digitalocean.component';
import { BringyourownClusterSettingsComponent } from './provider-settings/bringyourown/bringyourown.component';
import { AWSClusterSettingsComponent } from './provider-settings/aws/aws.component';
import { OpenstackClusterSettingsComponent } from './provider-settings/openstack/openstack.component';
import { AzureClusterSettingsComponent } from './provider-settings/azure/azure.component';
import { OpenstackAddNodeComponent } from '../../add-node/openstack-add-node/openstack-add-node.component';
import { OpenstackOptionsComponent } from '../../add-node/openstack-add-node/openstack-options/openstack-options.component';
import { DigitaloceanAddNodeComponent } from '../../add-node/digitalocean-add-node/digitalocean-add-node.component';
import { DigitaloceanOptionsComponent } from '../../add-node/digitalocean-add-node/digitalocean-options/digitalocean-options.component';
import { AwsAddNodeComponent } from '../../add-node/aws-add-node/aws-add-node.component';
import { AzureAddNodeComponent } from '../../add-node/azure-add-node/azure-add-node.component';
import { AddNodeService } from '../../core/services/add-node/add-node.service';
import { fakeDigitaloceanSizes } from '../../testing/fake-data/addNodeModal.fake';
import { asyncData } from '../../testing/services/api-mock.service';
import { ApiService, DatacenterService } from '../../core/services';
import { fakeSSHKeys } from '../../testing/fake-data/sshkey.fake';
import { HetznerClusterSettingsComponent } from './provider-settings/hetzner/hetzner.component';
import { VSphereClusterSettingsComponent } from './provider-settings/vsphere/vsphere.component';
import Spy = jasmine.Spy;
import { HetznerAddNodeComponent } from '../../add-node/hetzner-add-node/hetzner-add-node.component';
import { VSphereAddNodeComponent } from '../../add-node/vsphere-add-node/vsphere-add-node.component';
import { VSphereOptionsComponent } from '../../add-node/vsphere-add-node/vsphere-options/vsphere-options.component';
import { nodeDataFake } from '../../testing/fake-data/node.fake';
import { DatacenterMockService } from '../../testing/services/datacenter-mock.service';

describe('SetSettingsComponent', () => {
  let fixture: ComponentFixture<SetSettingsComponent>;
  let component: SetSettingsComponent;
  let getDigitaloceanSizesSpy: Spy;
  let getDigitaloceanSizesForWizardSpy: Spy;
  let getSSHKeysSpy: Spy;

  beforeEach(async(() => {
    const apiMock = jasmine.createSpyObj('ApiService', ['getDigitaloceanSizes', 'getDigitaloceanSizesForWizard', 'getSSHKeys']);
    getDigitaloceanSizesSpy = apiMock.getDigitaloceanSizes.and.returnValue(asyncData(fakeDigitaloceanSizes));
    getDigitaloceanSizesForWizardSpy = apiMock.getDigitaloceanSizesForWizard.and.returnValue(asyncData(fakeDigitaloceanSizes));
    getSSHKeysSpy = apiMock.getSSHKeys.and.returnValue(asyncData(fakeSSHKeys));

    TestBed.configureTestingModule({
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule,
        MatTabsModule,
      ],
      declarations: [
        SetSettingsComponent,
        ClusterSSHKeysComponent,
        ClusterProviderSettingsComponent,
        DigitaloceanClusterSettingsComponent,
        AWSClusterSettingsComponent,
        OpenstackClusterSettingsComponent,
        BringyourownClusterSettingsComponent,
        HetznerClusterSettingsComponent,
        VSphereClusterSettingsComponent,
        AzureClusterSettingsComponent,
        AddNodeComponent,
        OpenstackAddNodeComponent,
        OpenstackOptionsComponent,
        AwsAddNodeComponent,
        DigitaloceanAddNodeComponent,
        DigitaloceanOptionsComponent,
        HetznerAddNodeComponent,
        VSphereAddNodeComponent,
        VSphereOptionsComponent,
        AzureAddNodeComponent,
      ],
      providers: [
        AddNodeService,
        WizardService,
        { provide: DatacenterService, useClass: DatacenterMockService },
        { provide: ApiService, useValue: apiMock },
      ],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SetSettingsComponent);
    component = fixture.componentInstance;
    component.cluster = fakeDigitaloceanCluster;
    component.clusterSSHKeys = [];
    component.nodeData = nodeDataFake;
    fixture.detectChanges();
  });

  it('should create the set-settings cmp', () => {
    expect(component).toBeTruthy();
  });
});

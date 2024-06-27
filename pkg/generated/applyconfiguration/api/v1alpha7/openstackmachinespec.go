/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha7

// OpenStackMachineSpecApplyConfiguration represents an declarative configuration of the OpenStackMachineSpec type for use
// with apply.
type OpenStackMachineSpecApplyConfiguration struct {
	ProviderID             *string                                       `json:"providerID,omitempty"`
	InstanceID             *string                                       `json:"instanceID,omitempty"`
	CloudName              *string                                       `json:"cloudName,omitempty"`
	Flavor                 *string                                       `json:"flavor,omitempty"`
	Image                  *string                                       `json:"image,omitempty"`
	ImageUUID              *string                                       `json:"imageUUID,omitempty"`
	SSHKeyName             *string                                       `json:"sshKeyName,omitempty"`
	Ports                  []PortOptsApplyConfiguration                  `json:"ports,omitempty"`
	FloatingIP             *string                                       `json:"floatingIP,omitempty"`
	SecurityGroups         []SecurityGroupFilterApplyConfiguration       `json:"securityGroups,omitempty"`
	Trunk                  *bool                                         `json:"trunk,omitempty"`
	Tags                   []string                                      `json:"tags,omitempty"`
	ServerMetadata         map[string]string                             `json:"serverMetadata,omitempty"`
	ConfigDrive            *bool                                         `json:"configDrive,omitempty"`
	RootVolume             *RootVolumeApplyConfiguration                 `json:"rootVolume,omitempty"`
	AdditionalBlockDevices []AdditionalBlockDeviceApplyConfiguration     `json:"additionalBlockDevices,omitempty"`
	ServerGroupID          *string                                       `json:"serverGroupID,omitempty"`
	IdentityRef            *OpenStackIdentityReferenceApplyConfiguration `json:"identityRef,omitempty"`
}

// OpenStackMachineSpecApplyConfiguration constructs an declarative configuration of the OpenStackMachineSpec type for use with
// apply.
func OpenStackMachineSpec() *OpenStackMachineSpecApplyConfiguration {
	return &OpenStackMachineSpecApplyConfiguration{}
}

// WithProviderID sets the ProviderID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProviderID field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithProviderID(value string) *OpenStackMachineSpecApplyConfiguration {
	b.ProviderID = &value
	return b
}

// WithInstanceID sets the InstanceID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InstanceID field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithInstanceID(value string) *OpenStackMachineSpecApplyConfiguration {
	b.InstanceID = &value
	return b
}

// WithCloudName sets the CloudName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CloudName field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithCloudName(value string) *OpenStackMachineSpecApplyConfiguration {
	b.CloudName = &value
	return b
}

// WithFlavor sets the Flavor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Flavor field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithFlavor(value string) *OpenStackMachineSpecApplyConfiguration {
	b.Flavor = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithImage(value string) *OpenStackMachineSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithImageUUID sets the ImageUUID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageUUID field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithImageUUID(value string) *OpenStackMachineSpecApplyConfiguration {
	b.ImageUUID = &value
	return b
}

// WithSSHKeyName sets the SSHKeyName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SSHKeyName field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithSSHKeyName(value string) *OpenStackMachineSpecApplyConfiguration {
	b.SSHKeyName = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *OpenStackMachineSpecApplyConfiguration) WithPorts(values ...*PortOptsApplyConfiguration) *OpenStackMachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithFloatingIP sets the FloatingIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FloatingIP field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithFloatingIP(value string) *OpenStackMachineSpecApplyConfiguration {
	b.FloatingIP = &value
	return b
}

// WithSecurityGroups adds the given value to the SecurityGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SecurityGroups field.
func (b *OpenStackMachineSpecApplyConfiguration) WithSecurityGroups(values ...*SecurityGroupFilterApplyConfiguration) *OpenStackMachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSecurityGroups")
		}
		b.SecurityGroups = append(b.SecurityGroups, *values[i])
	}
	return b
}

// WithTrunk sets the Trunk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Trunk field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithTrunk(value bool) *OpenStackMachineSpecApplyConfiguration {
	b.Trunk = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *OpenStackMachineSpecApplyConfiguration) WithTags(values ...string) *OpenStackMachineSpecApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithServerMetadata puts the entries into the ServerMetadata field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ServerMetadata field,
// overwriting an existing map entries in ServerMetadata field with the same key.
func (b *OpenStackMachineSpecApplyConfiguration) WithServerMetadata(entries map[string]string) *OpenStackMachineSpecApplyConfiguration {
	if b.ServerMetadata == nil && len(entries) > 0 {
		b.ServerMetadata = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ServerMetadata[k] = v
	}
	return b
}

// WithConfigDrive sets the ConfigDrive field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigDrive field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithConfigDrive(value bool) *OpenStackMachineSpecApplyConfiguration {
	b.ConfigDrive = &value
	return b
}

// WithRootVolume sets the RootVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RootVolume field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithRootVolume(value *RootVolumeApplyConfiguration) *OpenStackMachineSpecApplyConfiguration {
	b.RootVolume = value
	return b
}

// WithAdditionalBlockDevices adds the given value to the AdditionalBlockDevices field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalBlockDevices field.
func (b *OpenStackMachineSpecApplyConfiguration) WithAdditionalBlockDevices(values ...*AdditionalBlockDeviceApplyConfiguration) *OpenStackMachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdditionalBlockDevices")
		}
		b.AdditionalBlockDevices = append(b.AdditionalBlockDevices, *values[i])
	}
	return b
}

// WithServerGroupID sets the ServerGroupID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerGroupID field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithServerGroupID(value string) *OpenStackMachineSpecApplyConfiguration {
	b.ServerGroupID = &value
	return b
}

// WithIdentityRef sets the IdentityRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IdentityRef field is set to the value of the last call.
func (b *OpenStackMachineSpecApplyConfiguration) WithIdentityRef(value *OpenStackIdentityReferenceApplyConfiguration) *OpenStackMachineSpecApplyConfiguration {
	b.IdentityRef = value
	return b
}

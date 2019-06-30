/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/metal3-io/machine-remediation-request-operator/pkg/apis/machineremediationrequest/v1alpha1"
)

// FakeMachineRemediationRequests implements MachineRemediationRequestInterface
type FakeMachineRemediationRequests struct {
	Fake *FakeMachineremediationrequestV1alpha1
	ns   string
}

var machineremediationrequestsResource = schema.GroupVersionResource{Group: "machineremediationrequest.openshift.io", Version: "v1alpha1", Resource: "machineremediationrequests"}

var machineremediationrequestsKind = schema.GroupVersionKind{Group: "machineremediationrequest.openshift.io", Version: "v1alpha1", Kind: "MachineRemediationRequest"}

// Get takes name of the machineRemediationRequest, and returns the corresponding machineRemediationRequest object, and an error if there is any.
func (c *FakeMachineRemediationRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineRemediationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machineremediationrequestsResource, c.ns, name), &v1alpha1.MachineRemediationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediationRequest), err
}

// List takes label and field selectors, and returns the list of MachineRemediationRequests that match those selectors.
func (c *FakeMachineRemediationRequests) List(opts v1.ListOptions) (result *v1alpha1.MachineRemediationRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machineremediationrequestsResource, machineremediationrequestsKind, c.ns, opts), &v1alpha1.MachineRemediationRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineRemediationRequestList{ListMeta: obj.(*v1alpha1.MachineRemediationRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.MachineRemediationRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineRemediationRequests.
func (c *FakeMachineRemediationRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machineremediationrequestsResource, c.ns, opts))

}

// Create takes the representation of a machineRemediationRequest and creates it.  Returns the server's representation of the machineRemediationRequest, and an error, if there is any.
func (c *FakeMachineRemediationRequests) Create(machineRemediationRequest *v1alpha1.MachineRemediationRequest) (result *v1alpha1.MachineRemediationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machineremediationrequestsResource, c.ns, machineRemediationRequest), &v1alpha1.MachineRemediationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediationRequest), err
}

// Update takes the representation of a machineRemediationRequest and updates it. Returns the server's representation of the machineRemediationRequest, and an error, if there is any.
func (c *FakeMachineRemediationRequests) Update(machineRemediationRequest *v1alpha1.MachineRemediationRequest) (result *v1alpha1.MachineRemediationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machineremediationrequestsResource, c.ns, machineRemediationRequest), &v1alpha1.MachineRemediationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediationRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineRemediationRequests) UpdateStatus(machineRemediationRequest *v1alpha1.MachineRemediationRequest) (*v1alpha1.MachineRemediationRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machineremediationrequestsResource, "status", c.ns, machineRemediationRequest), &v1alpha1.MachineRemediationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediationRequest), err
}

// Delete takes name of the machineRemediationRequest and deletes it. Returns an error if one occurs.
func (c *FakeMachineRemediationRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machineremediationrequestsResource, c.ns, name), &v1alpha1.MachineRemediationRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineRemediationRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machineremediationrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineRemediationRequestList{})
	return err
}

// Patch applies the patch and returns the patched machineRemediationRequest.
func (c *FakeMachineRemediationRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineRemediationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machineremediationrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MachineRemediationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineRemediationRequest), err
}

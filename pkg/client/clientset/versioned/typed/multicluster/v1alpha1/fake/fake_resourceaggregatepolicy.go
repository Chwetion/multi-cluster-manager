/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceAggregatePolicies implements ResourceAggregatePolicyInterface
type FakeResourceAggregatePolicies struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var resourceaggregatepoliciesResource = schema.GroupVersionResource{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Resource: "resourceaggregatepolicies"}

var resourceaggregatepoliciesKind = schema.GroupVersionKind{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Kind: "ResourceAggregatePolicy"}

// Get takes name of the resourceAggregatePolicy, and returns the corresponding resourceAggregatePolicy object, and an error if there is any.
func (c *FakeResourceAggregatePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceAggregatePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourceaggregatepoliciesResource, c.ns, name), &v1alpha1.ResourceAggregatePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceAggregatePolicy), err
}

// List takes label and field selectors, and returns the list of ResourceAggregatePolicies that match those selectors.
func (c *FakeResourceAggregatePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceAggregatePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourceaggregatepoliciesResource, resourceaggregatepoliciesKind, c.ns, opts), &v1alpha1.ResourceAggregatePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceAggregatePolicyList{ListMeta: obj.(*v1alpha1.ResourceAggregatePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceAggregatePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceAggregatePolicies.
func (c *FakeResourceAggregatePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourceaggregatepoliciesResource, c.ns, opts))

}

// Create takes the representation of a resourceAggregatePolicy and creates it.  Returns the server's representation of the resourceAggregatePolicy, and an error, if there is any.
func (c *FakeResourceAggregatePolicies) Create(ctx context.Context, resourceAggregatePolicy *v1alpha1.ResourceAggregatePolicy, opts v1.CreateOptions) (result *v1alpha1.ResourceAggregatePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourceaggregatepoliciesResource, c.ns, resourceAggregatePolicy), &v1alpha1.ResourceAggregatePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceAggregatePolicy), err
}

// Update takes the representation of a resourceAggregatePolicy and updates it. Returns the server's representation of the resourceAggregatePolicy, and an error, if there is any.
func (c *FakeResourceAggregatePolicies) Update(ctx context.Context, resourceAggregatePolicy *v1alpha1.ResourceAggregatePolicy, opts v1.UpdateOptions) (result *v1alpha1.ResourceAggregatePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourceaggregatepoliciesResource, c.ns, resourceAggregatePolicy), &v1alpha1.ResourceAggregatePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceAggregatePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceAggregatePolicies) UpdateStatus(ctx context.Context, resourceAggregatePolicy *v1alpha1.ResourceAggregatePolicy, opts v1.UpdateOptions) (*v1alpha1.ResourceAggregatePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourceaggregatepoliciesResource, "status", c.ns, resourceAggregatePolicy), &v1alpha1.ResourceAggregatePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceAggregatePolicy), err
}

// Delete takes name of the resourceAggregatePolicy and deletes it. Returns an error if one occurs.
func (c *FakeResourceAggregatePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourceaggregatepoliciesResource, c.ns, name), &v1alpha1.ResourceAggregatePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceAggregatePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourceaggregatepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceAggregatePolicyList{})
	return err
}

// Patch applies the patch and returns the patched resourceAggregatePolicy.
func (c *FakeResourceAggregatePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceAggregatePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourceaggregatepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceAggregatePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceAggregatePolicy), err
}
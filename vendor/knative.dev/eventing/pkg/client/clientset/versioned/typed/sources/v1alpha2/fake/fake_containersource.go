/*
Copyright 2021 The Knative Authors

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "knative.dev/eventing/pkg/apis/sources/v1alpha2"
)

// FakeContainerSources implements ContainerSourceInterface
type FakeContainerSources struct {
	Fake *FakeSourcesV1alpha2
	ns   string
}

var containersourcesResource = schema.GroupVersionResource{Group: "sources.knative.dev", Version: "v1alpha2", Resource: "containersources"}

var containersourcesKind = schema.GroupVersionKind{Group: "sources.knative.dev", Version: "v1alpha2", Kind: "ContainerSource"}

// Get takes name of the containerSource, and returns the corresponding containerSource object, and an error if there is any.
func (c *FakeContainerSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ContainerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containersourcesResource, c.ns, name), &v1alpha2.ContainerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ContainerSource), err
}

// List takes label and field selectors, and returns the list of ContainerSources that match those selectors.
func (c *FakeContainerSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ContainerSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(containersourcesResource, containersourcesKind, c.ns, opts), &v1alpha2.ContainerSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ContainerSourceList{ListMeta: obj.(*v1alpha2.ContainerSourceList).ListMeta}
	for _, item := range obj.(*v1alpha2.ContainerSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerSources.
func (c *FakeContainerSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containersourcesResource, c.ns, opts))

}

// Create takes the representation of a containerSource and creates it.  Returns the server's representation of the containerSource, and an error, if there is any.
func (c *FakeContainerSources) Create(ctx context.Context, containerSource *v1alpha2.ContainerSource, opts v1.CreateOptions) (result *v1alpha2.ContainerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(containersourcesResource, c.ns, containerSource), &v1alpha2.ContainerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ContainerSource), err
}

// Update takes the representation of a containerSource and updates it. Returns the server's representation of the containerSource, and an error, if there is any.
func (c *FakeContainerSources) Update(ctx context.Context, containerSource *v1alpha2.ContainerSource, opts v1.UpdateOptions) (result *v1alpha2.ContainerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(containersourcesResource, c.ns, containerSource), &v1alpha2.ContainerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ContainerSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerSources) UpdateStatus(ctx context.Context, containerSource *v1alpha2.ContainerSource, opts v1.UpdateOptions) (*v1alpha2.ContainerSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(containersourcesResource, "status", c.ns, containerSource), &v1alpha2.ContainerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ContainerSource), err
}

// Delete takes name of the containerSource and deletes it. Returns an error if one occurs.
func (c *FakeContainerSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(containersourcesResource, c.ns, name), &v1alpha2.ContainerSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(containersourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ContainerSourceList{})
	return err
}

// Patch applies the patch and returns the patched containerSource.
func (c *FakeContainerSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ContainerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(containersourcesResource, c.ns, name, pt, data, subresources...), &v1alpha2.ContainerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ContainerSource), err
}

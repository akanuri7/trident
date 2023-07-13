// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentActionSnapshotRestoresGetter has a method to return a TridentActionSnapshotRestoreInterface.
// A group's client should implement this interface.
type TridentActionSnapshotRestoresGetter interface {
	TridentActionSnapshotRestores(namespace string) TridentActionSnapshotRestoreInterface
}

// TridentActionSnapshotRestoreInterface has methods to work with TridentActionSnapshotRestore resources.
type TridentActionSnapshotRestoreInterface interface {
	Create(ctx context.Context, tridentActionSnapshotRestore *v1.TridentActionSnapshotRestore, opts metav1.CreateOptions) (*v1.TridentActionSnapshotRestore, error)
	Update(ctx context.Context, tridentActionSnapshotRestore *v1.TridentActionSnapshotRestore, opts metav1.UpdateOptions) (*v1.TridentActionSnapshotRestore, error)
	UpdateStatus(ctx context.Context, tridentActionSnapshotRestore *v1.TridentActionSnapshotRestore, opts metav1.UpdateOptions) (*v1.TridentActionSnapshotRestore, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentActionSnapshotRestore, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentActionSnapshotRestoreList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentActionSnapshotRestore, err error)
	TridentActionSnapshotRestoreExpansion
}

// tridentActionSnapshotRestores implements TridentActionSnapshotRestoreInterface
type tridentActionSnapshotRestores struct {
	client rest.Interface
	ns     string
}

// newTridentActionSnapshotRestores returns a TridentActionSnapshotRestores
func newTridentActionSnapshotRestores(c *TridentV1Client, namespace string) *tridentActionSnapshotRestores {
	return &tridentActionSnapshotRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentActionSnapshotRestore, and returns the corresponding tridentActionSnapshotRestore object, and an error if there is any.
func (c *tridentActionSnapshotRestores) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentActionSnapshotRestore, err error) {
	result = &v1.TridentActionSnapshotRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentActionSnapshotRestores that match those selectors.
func (c *tridentActionSnapshotRestores) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentActionSnapshotRestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentActionSnapshotRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentActionSnapshotRestores.
func (c *tridentActionSnapshotRestores) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentActionSnapshotRestore and creates it.  Returns the server's representation of the tridentActionSnapshotRestore, and an error, if there is any.
func (c *tridentActionSnapshotRestores) Create(ctx context.Context, tridentActionSnapshotRestore *v1.TridentActionSnapshotRestore, opts metav1.CreateOptions) (result *v1.TridentActionSnapshotRestore, err error) {
	result = &v1.TridentActionSnapshotRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentActionSnapshotRestore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentActionSnapshotRestore and updates it. Returns the server's representation of the tridentActionSnapshotRestore, and an error, if there is any.
func (c *tridentActionSnapshotRestores) Update(ctx context.Context, tridentActionSnapshotRestore *v1.TridentActionSnapshotRestore, opts metav1.UpdateOptions) (result *v1.TridentActionSnapshotRestore, err error) {
	result = &v1.TridentActionSnapshotRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		Name(tridentActionSnapshotRestore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentActionSnapshotRestore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tridentActionSnapshotRestores) UpdateStatus(ctx context.Context, tridentActionSnapshotRestore *v1.TridentActionSnapshotRestore, opts metav1.UpdateOptions) (result *v1.TridentActionSnapshotRestore, err error) {
	result = &v1.TridentActionSnapshotRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		Name(tridentActionSnapshotRestore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentActionSnapshotRestore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentActionSnapshotRestore and deletes it. Returns an error if one occurs.
func (c *tridentActionSnapshotRestores) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentActionSnapshotRestores) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentActionSnapshotRestore.
func (c *tridentActionSnapshotRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentActionSnapshotRestore, err error) {
	result = &v1.TridentActionSnapshotRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentactionsnapshotrestores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	secondaryschedulerv1 "github.com/openshift/secondary-scheduler-operator/pkg/apis/secondaryscheduler/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecondarySchedulers implements SecondarySchedulerInterface
type FakeSecondarySchedulers struct {
	Fake *FakeSecondaryschedulersV1
	ns   string
}

var secondaryschedulersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "secondaryschedulers"}

var secondaryschedulersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "SecondaryScheduler"}

// Get takes name of the secondaryScheduler, and returns the corresponding secondaryScheduler object, and an error if there is any.
func (c *FakeSecondarySchedulers) Get(ctx context.Context, name string, options v1.GetOptions) (result *secondaryschedulerv1.SecondaryScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secondaryschedulersResource, c.ns, name), &secondaryschedulerv1.SecondaryScheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*secondaryschedulerv1.SecondaryScheduler), err
}

// List takes label and field selectors, and returns the list of SecondarySchedulers that match those selectors.
func (c *FakeSecondarySchedulers) List(ctx context.Context, opts v1.ListOptions) (result *secondaryschedulerv1.SecondarySchedulerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secondaryschedulersResource, secondaryschedulersKind, c.ns, opts), &secondaryschedulerv1.SecondarySchedulerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &secondaryschedulerv1.SecondarySchedulerList{ListMeta: obj.(*secondaryschedulerv1.SecondarySchedulerList).ListMeta}
	for _, item := range obj.(*secondaryschedulerv1.SecondarySchedulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secondarySchedulers.
func (c *FakeSecondarySchedulers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secondaryschedulersResource, c.ns, opts))

}

// Create takes the representation of a secondaryScheduler and creates it.  Returns the server's representation of the secondaryScheduler, and an error, if there is any.
func (c *FakeSecondarySchedulers) Create(ctx context.Context, secondaryScheduler *secondaryschedulerv1.SecondaryScheduler, opts v1.CreateOptions) (result *secondaryschedulerv1.SecondaryScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secondaryschedulersResource, c.ns, secondaryScheduler), &secondaryschedulerv1.SecondaryScheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*secondaryschedulerv1.SecondaryScheduler), err
}

// Update takes the representation of a secondaryScheduler and updates it. Returns the server's representation of the secondaryScheduler, and an error, if there is any.
func (c *FakeSecondarySchedulers) Update(ctx context.Context, secondaryScheduler *secondaryschedulerv1.SecondaryScheduler, opts v1.UpdateOptions) (result *secondaryschedulerv1.SecondaryScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secondaryschedulersResource, c.ns, secondaryScheduler), &secondaryschedulerv1.SecondaryScheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*secondaryschedulerv1.SecondaryScheduler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecondarySchedulers) UpdateStatus(ctx context.Context, secondaryScheduler *secondaryschedulerv1.SecondaryScheduler, opts v1.UpdateOptions) (*secondaryschedulerv1.SecondaryScheduler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secondaryschedulersResource, "status", c.ns, secondaryScheduler), &secondaryschedulerv1.SecondaryScheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*secondaryschedulerv1.SecondaryScheduler), err
}

// Delete takes name of the secondaryScheduler and deletes it. Returns an error if one occurs.
func (c *FakeSecondarySchedulers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secondaryschedulersResource, c.ns, name), &secondaryschedulerv1.SecondaryScheduler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecondarySchedulers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secondaryschedulersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &secondaryschedulerv1.SecondarySchedulerList{})
	return err
}

// Patch applies the patch and returns the patched secondaryScheduler.
func (c *FakeSecondarySchedulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *secondaryschedulerv1.SecondaryScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secondaryschedulersResource, c.ns, name, pt, data, subresources...), &secondaryschedulerv1.SecondaryScheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*secondaryschedulerv1.SecondaryScheduler), err
}

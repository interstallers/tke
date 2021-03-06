/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
	v1 "tkestack.io/tke/api/platform/v1"
)

// TappControllersGetter has a method to return a TappControllerInterface.
// A group's client should implement this interface.
type TappControllersGetter interface {
	TappControllers() TappControllerInterface
}

// TappControllerInterface has methods to work with TappController resources.
type TappControllerInterface interface {
	Create(*v1.TappController) (*v1.TappController, error)
	Update(*v1.TappController) (*v1.TappController, error)
	UpdateStatus(*v1.TappController) (*v1.TappController, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.TappController, error)
	List(opts metav1.ListOptions) (*v1.TappControllerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TappController, err error)
	TappControllerExpansion
}

// tappControllers implements TappControllerInterface
type tappControllers struct {
	client rest.Interface
}

// newTappControllers returns a TappControllers
func newTappControllers(c *PlatformV1Client) *tappControllers {
	return &tappControllers{
		client: c.RESTClient(),
	}
}

// Get takes name of the tappController, and returns the corresponding tappController object, and an error if there is any.
func (c *tappControllers) Get(name string, options metav1.GetOptions) (result *v1.TappController, err error) {
	result = &v1.TappController{}
	err = c.client.Get().
		Resource("tappcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TappControllers that match those selectors.
func (c *tappControllers) List(opts metav1.ListOptions) (result *v1.TappControllerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TappControllerList{}
	err = c.client.Get().
		Resource("tappcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tappControllers.
func (c *tappControllers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tappcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a tappController and creates it.  Returns the server's representation of the tappController, and an error, if there is any.
func (c *tappControllers) Create(tappController *v1.TappController) (result *v1.TappController, err error) {
	result = &v1.TappController{}
	err = c.client.Post().
		Resource("tappcontrollers").
		Body(tappController).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tappController and updates it. Returns the server's representation of the tappController, and an error, if there is any.
func (c *tappControllers) Update(tappController *v1.TappController) (result *v1.TappController, err error) {
	result = &v1.TappController{}
	err = c.client.Put().
		Resource("tappcontrollers").
		Name(tappController.Name).
		Body(tappController).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *tappControllers) UpdateStatus(tappController *v1.TappController) (result *v1.TappController, err error) {
	result = &v1.TappController{}
	err = c.client.Put().
		Resource("tappcontrollers").
		Name(tappController.Name).
		SubResource("status").
		Body(tappController).
		Do().
		Into(result)
	return
}

// Delete takes name of the tappController and deletes it. Returns an error if one occurs.
func (c *tappControllers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tappcontrollers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched tappController.
func (c *tappControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TappController, err error) {
	result = &v1.TappController{}
	err = c.client.Patch(pt).
		Resource("tappcontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

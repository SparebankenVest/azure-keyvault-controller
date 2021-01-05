/*
Copyright Sparebanken Vest

Based on the Kubernetes controller example at
https://github.com/kubernetes/sample-controller

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

package v1

import (
	"context"
	"time"

	v1 "github.com/SparebankenVest/azure-key-vault-to-kubernetes/pkg/k8s/apis/azurekeyvault/v1"
	scheme "github.com/SparebankenVest/azure-key-vault-to-kubernetes/pkg/k8s/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureKeyVaultSecretsGetter has a method to return a AzureKeyVaultSecretInterface.
// A group's client should implement this interface.
type AzureKeyVaultSecretsGetter interface {
	AzureKeyVaultSecrets(namespace string) AzureKeyVaultSecretInterface
}

// AzureKeyVaultSecretInterface has methods to work with AzureKeyVaultSecret resources.
type AzureKeyVaultSecretInterface interface {
	Create(ctx context.Context, azureKeyVaultSecret *v1.AzureKeyVaultSecret, opts metav1.CreateOptions) (*v1.AzureKeyVaultSecret, error)
	Update(ctx context.Context, azureKeyVaultSecret *v1.AzureKeyVaultSecret, opts metav1.UpdateOptions) (*v1.AzureKeyVaultSecret, error)
	UpdateStatus(ctx context.Context, azureKeyVaultSecret *v1.AzureKeyVaultSecret, opts metav1.UpdateOptions) (*v1.AzureKeyVaultSecret, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AzureKeyVaultSecret, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AzureKeyVaultSecretList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AzureKeyVaultSecret, err error)
	AzureKeyVaultSecretExpansion
}

// azureKeyVaultSecrets implements AzureKeyVaultSecretInterface
type azureKeyVaultSecrets struct {
	client rest.Interface
	ns     string
}

// newAzureKeyVaultSecrets returns a AzureKeyVaultSecrets
func newAzureKeyVaultSecrets(c *KeyvaultV1Client, namespace string) *azureKeyVaultSecrets {
	return &azureKeyVaultSecrets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureKeyVaultSecret, and returns the corresponding azureKeyVaultSecret object, and an error if there is any.
func (c *azureKeyVaultSecrets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AzureKeyVaultSecret, err error) {
	result = &v1.AzureKeyVaultSecret{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureKeyVaultSecrets that match those selectors.
func (c *azureKeyVaultSecrets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AzureKeyVaultSecretList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AzureKeyVaultSecretList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureKeyVaultSecrets.
func (c *azureKeyVaultSecrets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azureKeyVaultSecret and creates it.  Returns the server's representation of the azureKeyVaultSecret, and an error, if there is any.
func (c *azureKeyVaultSecrets) Create(ctx context.Context, azureKeyVaultSecret *v1.AzureKeyVaultSecret, opts metav1.CreateOptions) (result *v1.AzureKeyVaultSecret, err error) {
	result = &v1.AzureKeyVaultSecret{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureKeyVaultSecret).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azureKeyVaultSecret and updates it. Returns the server's representation of the azureKeyVaultSecret, and an error, if there is any.
func (c *azureKeyVaultSecrets) Update(ctx context.Context, azureKeyVaultSecret *v1.AzureKeyVaultSecret, opts metav1.UpdateOptions) (result *v1.AzureKeyVaultSecret, err error) {
	result = &v1.AzureKeyVaultSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		Name(azureKeyVaultSecret.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureKeyVaultSecret).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *azureKeyVaultSecrets) UpdateStatus(ctx context.Context, azureKeyVaultSecret *v1.AzureKeyVaultSecret, opts metav1.UpdateOptions) (result *v1.AzureKeyVaultSecret, err error) {
	result = &v1.AzureKeyVaultSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		Name(azureKeyVaultSecret.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureKeyVaultSecret).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azureKeyVaultSecret and deletes it. Returns an error if one occurs.
func (c *azureKeyVaultSecrets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureKeyVaultSecrets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azureKeyVaultSecret.
func (c *azureKeyVaultSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AzureKeyVaultSecret, err error) {
	result = &v1.AzureKeyVaultSecret{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azurekeyvaultsecrets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

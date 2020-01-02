// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVault) DeepCopyInto(out *AzureKeyVault) {
	*out = *in
	out.Object = in.Object
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVault.
func (in *AzureKeyVault) DeepCopy() *AzureKeyVault {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultObject) DeepCopyInto(out *AzureKeyVaultObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultObject.
func (in *AzureKeyVaultObject) DeepCopy() *AzureKeyVaultObject {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultOutput) DeepCopyInto(out *AzureKeyVaultOutput) {
	*out = *in
	out.Secret = in.Secret
	if in.Transforms != nil {
		in, out := &in.Transforms, &out.Transforms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultOutput.
func (in *AzureKeyVaultOutput) DeepCopy() *AzureKeyVaultOutput {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultOutputSecret) DeepCopyInto(out *AzureKeyVaultOutputSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultOutputSecret.
func (in *AzureKeyVaultOutputSecret) DeepCopy() *AzureKeyVaultOutputSecret {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultOutputSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultSecret) DeepCopyInto(out *AzureKeyVaultSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultSecret.
func (in *AzureKeyVaultSecret) DeepCopy() *AzureKeyVaultSecret {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureKeyVaultSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultSecretList) DeepCopyInto(out *AzureKeyVaultSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureKeyVaultSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultSecretList.
func (in *AzureKeyVaultSecretList) DeepCopy() *AzureKeyVaultSecretList {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureKeyVaultSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultSecretSpec) DeepCopyInto(out *AzureKeyVaultSecretSpec) {
	*out = *in
	out.Vault = in.Vault
	in.Output.DeepCopyInto(&out.Output)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultSecretSpec.
func (in *AzureKeyVaultSecretSpec) DeepCopy() *AzureKeyVaultSecretSpec {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultSecretStatus) DeepCopyInto(out *AzureKeyVaultSecretStatus) {
	*out = *in
	in.LastAzureUpdate.DeepCopyInto(&out.LastAzureUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultSecretStatus.
func (in *AzureKeyVaultSecretStatus) DeepCopy() *AzureKeyVaultSecretStatus {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultSecretStatus)
	in.DeepCopyInto(out)
	return out
}

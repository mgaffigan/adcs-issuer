//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsIssuer) DeepCopyInto(out *AdcsIssuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsIssuer.
func (in *AdcsIssuer) DeepCopy() *AdcsIssuer {
	if in == nil {
		return nil
	}
	out := new(AdcsIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdcsIssuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsIssuerList) DeepCopyInto(out *AdcsIssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdcsIssuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsIssuerList.
func (in *AdcsIssuerList) DeepCopy() *AdcsIssuerList {
	if in == nil {
		return nil
	}
	out := new(AdcsIssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdcsIssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsIssuerSpec) DeepCopyInto(out *AdcsIssuerSpec) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsIssuerSpec.
func (in *AdcsIssuerSpec) DeepCopy() *AdcsIssuerSpec {
	if in == nil {
		return nil
	}
	out := new(AdcsIssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsIssuerStatus) DeepCopyInto(out *AdcsIssuerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsIssuerStatus.
func (in *AdcsIssuerStatus) DeepCopy() *AdcsIssuerStatus {
	if in == nil {
		return nil
	}
	out := new(AdcsIssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsRequest) DeepCopyInto(out *AdcsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsRequest.
func (in *AdcsRequest) DeepCopy() *AdcsRequest {
	if in == nil {
		return nil
	}
	out := new(AdcsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdcsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsRequestList) DeepCopyInto(out *AdcsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdcsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsRequestList.
func (in *AdcsRequestList) DeepCopy() *AdcsRequestList {
	if in == nil {
		return nil
	}
	out := new(AdcsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdcsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsRequestSpec) DeepCopyInto(out *AdcsRequestSpec) {
	*out = *in
	if in.CSRPEM != nil {
		in, out := &in.CSRPEM, &out.CSRPEM
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	out.IssuerRef = in.IssuerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsRequestSpec.
func (in *AdcsRequestSpec) DeepCopy() *AdcsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(AdcsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdcsRequestStatus) DeepCopyInto(out *AdcsRequestStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdcsRequestStatus.
func (in *AdcsRequestStatus) DeepCopy() *AdcsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(AdcsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdcsIssuer) DeepCopyInto(out *ClusterAdcsIssuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdcsIssuer.
func (in *ClusterAdcsIssuer) DeepCopy() *ClusterAdcsIssuer {
	if in == nil {
		return nil
	}
	out := new(ClusterAdcsIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAdcsIssuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdcsIssuerList) DeepCopyInto(out *ClusterAdcsIssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAdcsIssuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdcsIssuerList.
func (in *ClusterAdcsIssuerList) DeepCopy() *ClusterAdcsIssuerList {
	if in == nil {
		return nil
	}
	out := new(ClusterAdcsIssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAdcsIssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdcsIssuerSpec) DeepCopyInto(out *ClusterAdcsIssuerSpec) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdcsIssuerSpec.
func (in *ClusterAdcsIssuerSpec) DeepCopy() *ClusterAdcsIssuerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAdcsIssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdcsIssuerStatus) DeepCopyInto(out *ClusterAdcsIssuerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdcsIssuerStatus.
func (in *ClusterAdcsIssuerStatus) DeepCopy() *ClusterAdcsIssuerStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAdcsIssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Tekton Authors

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

package config

import (
	pod "github.com/tektoncd/pipeline/pkg/apis/pipeline/pod"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactBucket) DeepCopyInto(out *ArtifactBucket) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactBucket.
func (in *ArtifactBucket) DeepCopy() *ArtifactBucket {
	if in == nil {
		return nil
	}
	out := new(ArtifactBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactPVC) DeepCopyInto(out *ArtifactPVC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactPVC.
func (in *ArtifactPVC) DeepCopy() *ArtifactPVC {
	if in == nil {
		return nil
	}
	out := new(ArtifactPVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Defaults) DeepCopyInto(out *Defaults) {
	*out = *in
	if in.DefaultPodTemplate != nil {
		in, out := &in.DefaultPodTemplate, &out.DefaultPodTemplate
		*out = new(pod.Template)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultAAPodTemplate != nil {
		in, out := &in.DefaultAAPodTemplate, &out.DefaultAAPodTemplate
		*out = new(pod.AffinityAssistantTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Defaults.
func (in *Defaults) DeepCopy() *Defaults {
	if in == nil {
		return nil
	}
	out := new(Defaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlags) DeepCopyInto(out *FeatureFlags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlags.
func (in *FeatureFlags) DeepCopy() *FeatureFlags {
	if in == nil {
		return nil
	}
	out := new(FeatureFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metrics) DeepCopyInto(out *Metrics) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metrics.
func (in *Metrics) DeepCopy() *Metrics {
	if in == nil {
		return nil
	}
	out := new(Metrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustedResources) DeepCopyInto(out *TrustedResources) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustedResources.
func (in *TrustedResources) DeepCopy() *TrustedResources {
	if in == nil {
		return nil
	}
	out := new(TrustedResources)
	in.DeepCopyInto(out)
	return out
}

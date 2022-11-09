//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionLimitSpec) DeepCopyInto(out *IngestionLimitSpec) {
	*out = *in
	if in.IngestionBurstSizeBytes != nil {
		in, out := &in.IngestionBurstSizeBytes, &out.IngestionBurstSizeBytes
		*out = new(int)
		**out = **in
	}
	if in.IngestionRateLimitBytes != nil {
		in, out := &in.IngestionRateLimitBytes, &out.IngestionRateLimitBytes
		*out = new(int)
		**out = **in
	}
	if in.MaxBytesPerTrace != nil {
		in, out := &in.MaxBytesPerTrace, &out.MaxBytesPerTrace
		*out = new(int)
		**out = **in
	}
	if in.MaxTracesPerUser != nil {
		in, out := &in.MaxTracesPerUser, &out.MaxTracesPerUser
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionLimitSpec.
func (in *IngestionLimitSpec) DeepCopy() *IngestionLimitSpec {
	if in == nil {
		return nil
	}
	out := new(IngestionLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerQuerySpec) DeepCopyInto(out *JaegerQuerySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerQuerySpec.
func (in *JaegerQuerySpec) DeepCopy() *JaegerQuerySpec {
	if in == nil {
		return nil
	}
	out := new(JaegerQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitSpec) DeepCopyInto(out *LimitSpec) {
	*out = *in
	if in.PerTenant != nil {
		in, out := &in.PerTenant, &out.PerTenant
		*out = make(map[string]RateLimitSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Global.DeepCopyInto(&out.Global)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitSpec.
func (in *LimitSpec) DeepCopy() *LimitSpec {
	if in == nil {
		return nil
	}
	out := new(LimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Microservices) DeepCopyInto(out *Microservices) {
	*out = *in
	out.Status = in.Status
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Microservices.
func (in *Microservices) DeepCopy() *Microservices {
	if in == nil {
		return nil
	}
	out := new(Microservices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Microservices) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroservicesList) DeepCopyInto(out *MicroservicesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Microservices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroservicesList.
func (in *MicroservicesList) DeepCopy() *MicroservicesList {
	if in == nil {
		return nil
	}
	out := new(MicroservicesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicroservicesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroservicesSpec) DeepCopyInto(out *MicroservicesSpec) {
	*out = *in
	in.Components.DeepCopyInto(&out.Components)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Storage.DeepCopyInto(&out.Storage)
	in.Retention.DeepCopyInto(&out.Retention)
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	in.LimitSpec.DeepCopyInto(&out.LimitSpec)
	out.StorageSize = in.StorageSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroservicesSpec.
func (in *MicroservicesSpec) DeepCopy() *MicroservicesSpec {
	if in == nil {
		return nil
	}
	out := new(MicroservicesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroservicesStatus) DeepCopyInto(out *MicroservicesStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroservicesStatus.
func (in *MicroservicesStatus) DeepCopy() *MicroservicesStatus {
	if in == nil {
		return nil
	}
	out := new(MicroservicesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageSpec) DeepCopyInto(out *ObjectStorageSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(ObjectStorageTLSSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageSpec.
func (in *ObjectStorageSpec) DeepCopy() *ObjectStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageTLSSpec) DeepCopyInto(out *ObjectStorageTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageTLSSpec.
func (in *ObjectStorageTLSSpec) DeepCopy() *ObjectStorageTLSSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryLimit) DeepCopyInto(out *QueryLimit) {
	*out = *in
	if in.MaxBytesPerTagValues != nil {
		in, out := &in.MaxBytesPerTagValues, &out.MaxBytesPerTagValues
		*out = new(int)
		**out = **in
	}
	if in.MaxSearchBytesPerTrace != nil {
		in, out := &in.MaxSearchBytesPerTrace, &out.MaxSearchBytesPerTrace
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryLimit.
func (in *QueryLimit) DeepCopy() *QueryLimit {
	if in == nil {
		return nil
	}
	out := new(QueryLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitSpec) DeepCopyInto(out *RateLimitSpec) {
	*out = *in
	in.Ingestion.DeepCopyInto(&out.Ingestion)
	in.Query.DeepCopyInto(&out.Query)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitSpec.
func (in *RateLimitSpec) DeepCopy() *RateLimitSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionConfig) DeepCopyInto(out *RetentionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionConfig.
func (in *RetentionConfig) DeepCopy() *RetentionConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionSpec) DeepCopyInto(out *RetentionSpec) {
	*out = *in
	if in.PerTenant != nil {
		in, out := &in.PerTenant, &out.PerTenant
		*out = make(map[string]RetentionConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Global = in.Global
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionSpec.
func (in *RetentionSpec) DeepCopy() *RetentionSpec {
	if in == nil {
		return nil
	}
	out := new(RetentionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoComponentSpec) DeepCopyInto(out *TempoComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoComponentSpec.
func (in *TempoComponentSpec) DeepCopy() *TempoComponentSpec {
	if in == nil {
		return nil
	}
	out := new(TempoComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoComponentsSpec) DeepCopyInto(out *TempoComponentsSpec) {
	*out = *in
	if in.Distributor != nil {
		in, out := &in.Distributor, &out.Distributor
		*out = new(TempoComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Ingester != nil {
		in, out := &in.Ingester, &out.Ingester
		*out = new(TempoComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Compactor != nil {
		in, out := &in.Compactor, &out.Compactor
		*out = new(TempoComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Querier != nil {
		in, out := &in.Querier, &out.Querier
		*out = new(TempoComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryFrontend != nil {
		in, out := &in.QueryFrontend, &out.QueryFrontend
		*out = new(TempoQueryFrontendSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoComponentsSpec.
func (in *TempoComponentsSpec) DeepCopy() *TempoComponentsSpec {
	if in == nil {
		return nil
	}
	out := new(TempoComponentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoQueryFrontendSpec) DeepCopyInto(out *TempoQueryFrontendSpec) {
	*out = *in
	in.TempoComponentSpec.DeepCopyInto(&out.TempoComponentSpec)
	out.JaegerQuery = in.JaegerQuery
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoQueryFrontendSpec.
func (in *TempoQueryFrontendSpec) DeepCopy() *TempoQueryFrontendSpec {
	if in == nil {
		return nil
	}
	out := new(TempoQueryFrontendSpec)
	in.DeepCopyInto(out)
	return out
}

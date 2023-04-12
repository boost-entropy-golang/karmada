//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSettings) DeepCopyInto(out *CommonSettings) {
	*out = *in
	out.Image = in.Image
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSettings.
func (in *CommonSettings) DeepCopy() *CommonSettings {
	if in == nil {
		return nil
	}
	out := new(CommonSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalEtcd)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalEtcd)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEtcd) DeepCopyInto(out *ExternalEtcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEtcd.
func (in *ExternalEtcd) DeepCopy() *ExternalEtcd {
	if in == nil {
		return nil
	}
	out := new(ExternalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostCluster) DeepCopyInto(out *HostCluster) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalSecretReference)
		**out = **in
	}
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(Networking)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostCluster.
func (in *HostCluster) DeepCopy() *HostCluster {
	if in == nil {
		return nil
	}
	out := new(HostCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistry) DeepCopyInto(out *ImageRegistry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistry.
func (in *ImageRegistry) DeepCopy() *ImageRegistry {
	if in == nil {
		return nil
	}
	out := new(ImageRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Karmada) DeepCopyInto(out *Karmada) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Karmada.
func (in *Karmada) DeepCopy() *Karmada {
	if in == nil {
		return nil
	}
	out := new(Karmada)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Karmada) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaAPIServer) DeepCopyInto(out *KarmadaAPIServer) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.ServiceSubnet != nil {
		in, out := &in.ServiceSubnet, &out.ServiceSubnet
		*out = new(string)
		**out = **in
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CertSANs != nil {
		in, out := &in.CertSANs, &out.CertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaAPIServer.
func (in *KarmadaAPIServer) DeepCopy() *KarmadaAPIServer {
	if in == nil {
		return nil
	}
	out := new(KarmadaAPIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaAggregratedAPIServer) DeepCopyInto(out *KarmadaAggregratedAPIServer) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CertSANs != nil {
		in, out := &in.CertSANs, &out.CertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaAggregratedAPIServer.
func (in *KarmadaAggregratedAPIServer) DeepCopy() *KarmadaAggregratedAPIServer {
	if in == nil {
		return nil
	}
	out := new(KarmadaAggregratedAPIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaComponents) DeepCopyInto(out *KarmadaComponents) {
	*out = *in
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(Etcd)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaAPIServer != nil {
		in, out := &in.KarmadaAPIServer, &out.KarmadaAPIServer
		*out = new(KarmadaAPIServer)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaAggregratedAPIServer != nil {
		in, out := &in.KarmadaAggregratedAPIServer, &out.KarmadaAggregratedAPIServer
		*out = new(KarmadaAggregratedAPIServer)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeControllerManager != nil {
		in, out := &in.KubeControllerManager, &out.KubeControllerManager
		*out = new(KubeControllerManager)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaControllerManager != nil {
		in, out := &in.KarmadaControllerManager, &out.KarmadaControllerManager
		*out = new(KarmadaControllerManager)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaScheduler != nil {
		in, out := &in.KarmadaScheduler, &out.KarmadaScheduler
		*out = new(KarmadaScheduler)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaWebhook != nil {
		in, out := &in.KarmadaWebhook, &out.KarmadaWebhook
		*out = new(KarmadaWebhook)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaDescheduler != nil {
		in, out := &in.KarmadaDescheduler, &out.KarmadaDescheduler
		*out = new(KarmadaDescheduler)
		(*in).DeepCopyInto(*out)
	}
	if in.KarmadaSearch != nil {
		in, out := &in.KarmadaSearch, &out.KarmadaSearch
		*out = new(KarmadaSearch)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaComponents.
func (in *KarmadaComponents) DeepCopy() *KarmadaComponents {
	if in == nil {
		return nil
	}
	out := new(KarmadaComponents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaControllerManager) DeepCopyInto(out *KarmadaControllerManager) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaControllerManager.
func (in *KarmadaControllerManager) DeepCopy() *KarmadaControllerManager {
	if in == nil {
		return nil
	}
	out := new(KarmadaControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaDescheduler) DeepCopyInto(out *KarmadaDescheduler) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaDescheduler.
func (in *KarmadaDescheduler) DeepCopy() *KarmadaDescheduler {
	if in == nil {
		return nil
	}
	out := new(KarmadaDescheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaList) DeepCopyInto(out *KarmadaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Karmada, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaList.
func (in *KarmadaList) DeepCopy() *KarmadaList {
	if in == nil {
		return nil
	}
	out := new(KarmadaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KarmadaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaScheduler) DeepCopyInto(out *KarmadaScheduler) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaScheduler.
func (in *KarmadaScheduler) DeepCopy() *KarmadaScheduler {
	if in == nil {
		return nil
	}
	out := new(KarmadaScheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaSearch) DeepCopyInto(out *KarmadaSearch) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaSearch.
func (in *KarmadaSearch) DeepCopy() *KarmadaSearch {
	if in == nil {
		return nil
	}
	out := new(KarmadaSearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaSpec) DeepCopyInto(out *KarmadaSpec) {
	*out = *in
	if in.HostCluster != nil {
		in, out := &in.HostCluster, &out.HostCluster
		*out = new(HostCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateRegistry != nil {
		in, out := &in.PrivateRegistry, &out.PrivateRegistry
		*out = new(ImageRegistry)
		**out = **in
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = new(KarmadaComponents)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaSpec.
func (in *KarmadaSpec) DeepCopy() *KarmadaSpec {
	if in == nil {
		return nil
	}
	out := new(KarmadaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaStatus) DeepCopyInto(out *KarmadaStatus) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalSecretReference)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaStatus.
func (in *KarmadaStatus) DeepCopy() *KarmadaStatus {
	if in == nil {
		return nil
	}
	out := new(KarmadaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaWebhook) DeepCopyInto(out *KarmadaWebhook) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaWebhook.
func (in *KarmadaWebhook) DeepCopy() *KarmadaWebhook {
	if in == nil {
		return nil
	}
	out := new(KarmadaWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManager) DeepCopyInto(out *KubeControllerManager) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManager.
func (in *KubeControllerManager) DeepCopy() *KubeControllerManager {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalEtcd) DeepCopyInto(out *LocalEtcd) {
	*out = *in
	in.CommonSettings.DeepCopyInto(&out.CommonSettings)
	if in.VolumeData != nil {
		in, out := &in.VolumeData, &out.VolumeData
		*out = new(VolumeData)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerCertSANs != nil {
		in, out := &in.ServerCertSANs, &out.ServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PeerCertSANs != nil {
		in, out := &in.PeerCertSANs, &out.PeerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalEtcd.
func (in *LocalEtcd) DeepCopy() *LocalEtcd {
	if in == nil {
		return nil
	}
	out := new(LocalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSecretReference) DeepCopyInto(out *LocalSecretReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSecretReference.
func (in *LocalSecretReference) DeepCopy() *LocalSecretReference {
	if in == nil {
		return nil
	}
	out := new(LocalSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	if in.DNSDomain != nil {
		in, out := &in.DNSDomain, &out.DNSDomain
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeData) DeepCopyInto(out *VolumeData) {
	*out = *in
	if in.VolumeClaim != nil {
		in, out := &in.VolumeClaim, &out.VolumeClaim
		*out = new(corev1.PersistentVolumeClaimTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(corev1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(corev1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeData.
func (in *VolumeData) DeepCopy() *VolumeData {
	if in == nil {
		return nil
	}
	out := new(VolumeData)
	in.DeepCopyInto(out)
	return out
}

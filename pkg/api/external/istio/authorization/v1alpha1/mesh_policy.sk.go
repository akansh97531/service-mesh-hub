// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewMeshPolicy(namespace, name string) *MeshPolicy {
	return &MeshPolicy{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *MeshPolicy) SetStatus(status core.Status) {
	r.Status = status
}

func (r *MeshPolicy) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *MeshPolicy) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Targets,
		r.Peers,
		r.PeerIsOptional,
		r.Origins,
		r.OriginIsOptional,
		r.PrincipalBinding,
	)
}

type MeshPolicyList []*MeshPolicy
type MeshpoliciesByNamespace map[string]MeshPolicyList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list MeshPolicyList) Find(namespace, name string) (*MeshPolicy, error) {
	for _, meshPolicy := range list {
		if meshPolicy.Metadata.Name == name {
			if namespace == "" || meshPolicy.Metadata.Namespace == namespace {
				return meshPolicy, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find meshPolicy %v.%v", namespace, name)
}

func (list MeshPolicyList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, meshPolicy := range list {
		ress = append(ress, meshPolicy)
	}
	return ress
}

func (list MeshPolicyList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, meshPolicy := range list {
		ress = append(ress, meshPolicy)
	}
	return ress
}

func (list MeshPolicyList) Names() []string {
	var names []string
	for _, meshPolicy := range list {
		names = append(names, meshPolicy.Metadata.Name)
	}
	return names
}

func (list MeshPolicyList) NamespacesDotNames() []string {
	var names []string
	for _, meshPolicy := range list {
		names = append(names, meshPolicy.Metadata.Namespace+"."+meshPolicy.Metadata.Name)
	}
	return names
}

func (list MeshPolicyList) Sort() MeshPolicyList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list MeshPolicyList) Clone() MeshPolicyList {
	var meshPolicyList MeshPolicyList
	for _, meshPolicy := range list {
		meshPolicyList = append(meshPolicyList, proto.Clone(meshPolicy).(*MeshPolicy))
	}
	return meshPolicyList
}

func (list MeshPolicyList) Each(f func(element *MeshPolicy)) {
	for _, meshPolicy := range list {
		f(meshPolicy)
	}
}

func (list MeshPolicyList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *MeshPolicy) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (list MeshPolicyList) ByNamespace() MeshpoliciesByNamespace {
	byNamespace := make(MeshpoliciesByNamespace)
	for _, meshPolicy := range list {
		byNamespace.Add(meshPolicy)
	}
	return byNamespace
}

func (byNamespace MeshpoliciesByNamespace) Add(meshPolicy ...*MeshPolicy) {
	for _, item := range meshPolicy {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace MeshpoliciesByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace MeshpoliciesByNamespace) List() MeshPolicyList {
	var list MeshPolicyList
	for _, meshPolicyList := range byNamespace {
		list = append(list, meshPolicyList...)
	}
	return list.Sort()
}

func (byNamespace MeshpoliciesByNamespace) Clone() MeshpoliciesByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &MeshPolicy{}

// Kubernetes Adapter for MeshPolicy

func (o *MeshPolicy) GetObjectKind() schema.ObjectKind {
	t := MeshPolicyCrd.TypeMeta()
	return &t
}

func (o *MeshPolicy) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*MeshPolicy)
}

var MeshPolicyCrd = crd.NewCrd("istio.authentication.v1alpha1",
	"meshpolicies",
	"istio.authentication.v1alpha1",
	"v1alpha1",
	"MeshPolicy",
	"meshpolicy",
	true,
	&MeshPolicy{})
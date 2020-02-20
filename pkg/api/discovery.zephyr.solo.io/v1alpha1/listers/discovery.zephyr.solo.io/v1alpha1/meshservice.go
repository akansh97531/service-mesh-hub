/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MeshServiceLister helps list MeshServices.
type MeshServiceLister interface {
	// List lists all MeshServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MeshService, err error)
	// MeshServices returns an object that can list and get MeshServices.
	MeshServices(namespace string) MeshServiceNamespaceLister
	MeshServiceListerExpansion
}

// meshServiceLister implements the MeshServiceLister interface.
type meshServiceLister struct {
	indexer cache.Indexer
}

// NewMeshServiceLister returns a new MeshServiceLister.
func NewMeshServiceLister(indexer cache.Indexer) MeshServiceLister {
	return &meshServiceLister{indexer: indexer}
}

// List lists all MeshServices in the indexer.
func (s *meshServiceLister) List(selector labels.Selector) (ret []*v1alpha1.MeshService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MeshService))
	})
	return ret, err
}

// MeshServices returns an object that can list and get MeshServices.
func (s *meshServiceLister) MeshServices(namespace string) MeshServiceNamespaceLister {
	return meshServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MeshServiceNamespaceLister helps list and get MeshServices.
type MeshServiceNamespaceLister interface {
	// List lists all MeshServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MeshService, err error)
	// Get retrieves the MeshService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MeshService, error)
	MeshServiceNamespaceListerExpansion
}

// meshServiceNamespaceLister implements the MeshServiceNamespaceLister
// interface.
type meshServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MeshServices in the indexer for a given namespace.
func (s meshServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MeshService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MeshService))
	})
	return ret, err
}

// Get retrieves the MeshService from the indexer for a given namespace and name.
func (s meshServiceNamespaceLister) Get(name string) (*v1alpha1.MeshService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("meshservice"), name)
	}
	return obj.(*v1alpha1.MeshService), nil
}
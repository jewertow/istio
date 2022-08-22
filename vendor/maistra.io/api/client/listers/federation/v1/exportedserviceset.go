// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "maistra.io/api/federation/v1"
)

// ExportedServiceSetLister helps list ExportedServiceSets.
// All objects returned here must be treated as read-only.
type ExportedServiceSetLister interface {
	// List lists all ExportedServiceSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ExportedServiceSet, err error)
	// ExportedServiceSets returns an object that can list and get ExportedServiceSets.
	ExportedServiceSets(namespace string) ExportedServiceSetNamespaceLister
	ExportedServiceSetListerExpansion
}

// exportedServiceSetLister implements the ExportedServiceSetLister interface.
type exportedServiceSetLister struct {
	indexer cache.Indexer
}

// NewExportedServiceSetLister returns a new ExportedServiceSetLister.
func NewExportedServiceSetLister(indexer cache.Indexer) ExportedServiceSetLister {
	return &exportedServiceSetLister{indexer: indexer}
}

// List lists all ExportedServiceSets in the indexer.
func (s *exportedServiceSetLister) List(selector labels.Selector) (ret []*v1.ExportedServiceSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ExportedServiceSet))
	})
	return ret, err
}

// ExportedServiceSets returns an object that can list and get ExportedServiceSets.
func (s *exportedServiceSetLister) ExportedServiceSets(namespace string) ExportedServiceSetNamespaceLister {
	return exportedServiceSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExportedServiceSetNamespaceLister helps list and get ExportedServiceSets.
// All objects returned here must be treated as read-only.
type ExportedServiceSetNamespaceLister interface {
	// List lists all ExportedServiceSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ExportedServiceSet, err error)
	// Get retrieves the ExportedServiceSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ExportedServiceSet, error)
	ExportedServiceSetNamespaceListerExpansion
}

// exportedServiceSetNamespaceLister implements the ExportedServiceSetNamespaceLister
// interface.
type exportedServiceSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExportedServiceSets in the indexer for a given namespace.
func (s exportedServiceSetNamespaceLister) List(selector labels.Selector) (ret []*v1.ExportedServiceSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ExportedServiceSet))
	})
	return ret, err
}

// Get retrieves the ExportedServiceSet from the indexer for a given namespace and name.
func (s exportedServiceSetNamespaceLister) Get(name string) (*v1.ExportedServiceSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("exportedserviceset"), name)
	}
	return obj.(*v1.ExportedServiceSet), nil
}

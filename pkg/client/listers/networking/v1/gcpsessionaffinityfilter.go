/*
* Copyright 2024 Google LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     https://www.apache.org/licenses/LICENSE-2.0
*
*     Unless required by applicable law or agreed to in writing, software
*     distributed under the License is distributed on an "AS IS" BASIS,
*     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*     See the License for the specific language governing permissions and
*     limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/GoogleCloudPlatform/gke-gateway-api/apis/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GCPSessionAffinityFilterLister helps list GCPSessionAffinityFilters.
// All objects returned here must be treated as read-only.
type GCPSessionAffinityFilterLister interface {
	// List lists all GCPSessionAffinityFilters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GCPSessionAffinityFilter, err error)
	// GCPSessionAffinityFilters returns an object that can list and get GCPSessionAffinityFilters.
	GCPSessionAffinityFilters(namespace string) GCPSessionAffinityFilterNamespaceLister
	GCPSessionAffinityFilterListerExpansion
}

// gCPSessionAffinityFilterLister implements the GCPSessionAffinityFilterLister interface.
type gCPSessionAffinityFilterLister struct {
	indexer cache.Indexer
}

// NewGCPSessionAffinityFilterLister returns a new GCPSessionAffinityFilterLister.
func NewGCPSessionAffinityFilterLister(indexer cache.Indexer) GCPSessionAffinityFilterLister {
	return &gCPSessionAffinityFilterLister{indexer: indexer}
}

// List lists all GCPSessionAffinityFilters in the indexer.
func (s *gCPSessionAffinityFilterLister) List(selector labels.Selector) (ret []*v1.GCPSessionAffinityFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GCPSessionAffinityFilter))
	})
	return ret, err
}

// GCPSessionAffinityFilters returns an object that can list and get GCPSessionAffinityFilters.
func (s *gCPSessionAffinityFilterLister) GCPSessionAffinityFilters(namespace string) GCPSessionAffinityFilterNamespaceLister {
	return gCPSessionAffinityFilterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GCPSessionAffinityFilterNamespaceLister helps list and get GCPSessionAffinityFilters.
// All objects returned here must be treated as read-only.
type GCPSessionAffinityFilterNamespaceLister interface {
	// List lists all GCPSessionAffinityFilters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GCPSessionAffinityFilter, err error)
	// Get retrieves the GCPSessionAffinityFilter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GCPSessionAffinityFilter, error)
	GCPSessionAffinityFilterNamespaceListerExpansion
}

// gCPSessionAffinityFilterNamespaceLister implements the GCPSessionAffinityFilterNamespaceLister
// interface.
type gCPSessionAffinityFilterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GCPSessionAffinityFilters in the indexer for a given namespace.
func (s gCPSessionAffinityFilterNamespaceLister) List(selector labels.Selector) (ret []*v1.GCPSessionAffinityFilter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GCPSessionAffinityFilter))
	})
	return ret, err
}

// Get retrieves the GCPSessionAffinityFilter from the indexer for a given namespace and name.
func (s gCPSessionAffinityFilterNamespaceLister) Get(name string) (*v1.GCPSessionAffinityFilter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gcpsessionaffinityfilter"), name)
	}
	return obj.(*v1.GCPSessionAffinityFilter), nil
}
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

// GCPBackendPolicyLister helps list GCPBackendPolicies.
// All objects returned here must be treated as read-only.
type GCPBackendPolicyLister interface {
	// List lists all GCPBackendPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GCPBackendPolicy, err error)
	// GCPBackendPolicies returns an object that can list and get GCPBackendPolicies.
	GCPBackendPolicies(namespace string) GCPBackendPolicyNamespaceLister
	GCPBackendPolicyListerExpansion
}

// gCPBackendPolicyLister implements the GCPBackendPolicyLister interface.
type gCPBackendPolicyLister struct {
	indexer cache.Indexer
}

// NewGCPBackendPolicyLister returns a new GCPBackendPolicyLister.
func NewGCPBackendPolicyLister(indexer cache.Indexer) GCPBackendPolicyLister {
	return &gCPBackendPolicyLister{indexer: indexer}
}

// List lists all GCPBackendPolicies in the indexer.
func (s *gCPBackendPolicyLister) List(selector labels.Selector) (ret []*v1.GCPBackendPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GCPBackendPolicy))
	})
	return ret, err
}

// GCPBackendPolicies returns an object that can list and get GCPBackendPolicies.
func (s *gCPBackendPolicyLister) GCPBackendPolicies(namespace string) GCPBackendPolicyNamespaceLister {
	return gCPBackendPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GCPBackendPolicyNamespaceLister helps list and get GCPBackendPolicies.
// All objects returned here must be treated as read-only.
type GCPBackendPolicyNamespaceLister interface {
	// List lists all GCPBackendPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GCPBackendPolicy, err error)
	// Get retrieves the GCPBackendPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GCPBackendPolicy, error)
	GCPBackendPolicyNamespaceListerExpansion
}

// gCPBackendPolicyNamespaceLister implements the GCPBackendPolicyNamespaceLister
// interface.
type gCPBackendPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GCPBackendPolicies in the indexer for a given namespace.
func (s gCPBackendPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.GCPBackendPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GCPBackendPolicy))
	})
	return ret, err
}

// Get retrieves the GCPBackendPolicy from the indexer for a given namespace and name.
func (s gCPBackendPolicyNamespaceLister) Get(name string) (*v1.GCPBackendPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gcpbackendpolicy"), name)
	}
	return obj.(*v1.GCPBackendPolicy), nil
}
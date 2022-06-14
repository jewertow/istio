/*
Copyright Red Hat, Inc.

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

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha3

import (
	internalinterfaces "github.com/maistra/xns-informer/pkg/generated/istio/internalinterfaces"
	informers "github.com/maistra/xns-informer/pkg/informers"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DestinationRules returns a DestinationRuleInformer.
	DestinationRules() DestinationRuleInformer
	// EnvoyFilters returns a EnvoyFilterInformer.
	EnvoyFilters() EnvoyFilterInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// ServiceEntries returns a ServiceEntryInformer.
	ServiceEntries() ServiceEntryInformer
	// Sidecars returns a SidecarInformer.
	Sidecars() SidecarInformer
	// VirtualServices returns a VirtualServiceInformer.
	VirtualServices() VirtualServiceInformer
	// WorkloadEntries returns a WorkloadEntryInformer.
	WorkloadEntries() WorkloadEntryInformer
	// WorkloadGroups returns a WorkloadGroupInformer.
	WorkloadGroups() WorkloadGroupInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespaces       informers.NamespaceSet
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespaces informers.NamespaceSet, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespaces: namespaces, tweakListOptions: tweakListOptions}
}

// DestinationRules returns a DestinationRuleInformer.
func (v *version) DestinationRules() DestinationRuleInformer {
	return &destinationRuleInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// EnvoyFilters returns a EnvoyFilterInformer.
func (v *version) EnvoyFilters() EnvoyFilterInformer {
	return &envoyFilterInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ServiceEntries returns a ServiceEntryInformer.
func (v *version) ServiceEntries() ServiceEntryInformer {
	return &serviceEntryInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Sidecars returns a SidecarInformer.
func (v *version) Sidecars() SidecarInformer {
	return &sidecarInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// VirtualServices returns a VirtualServiceInformer.
func (v *version) VirtualServices() VirtualServiceInformer {
	return &virtualServiceInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// WorkloadEntries returns a WorkloadEntryInformer.
func (v *version) WorkloadEntries() WorkloadEntryInformer {
	return &workloadEntryInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// WorkloadGroups returns a WorkloadGroupInformer.
func (v *version) WorkloadGroups() WorkloadGroupInformer {
	return &workloadGroupInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

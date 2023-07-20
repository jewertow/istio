//go:build integ
// +build integ

// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicemesh

import (
	"testing"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/tests/integration/servicemesh/maistra"
)

var istioNamespace namespace.Instance

func TestMain(m *testing.M) {
	// do not change order of setup functions
	// nolint: staticcheck
	framework.
		NewSuite(m).
		RequireMaxClusters(1).
		Setup(maistra.ApplyServiceMeshCRDs).
		Setup(namespace.Setup(&istioNamespace, namespace.Config{Prefix: "istio-system"})).
		Setup(maistra.Install(namespace.Future(&istioNamespace))).
		// We cannot apply restricted RBAC before the control plane installation, because the operator always applies
		// the default RBAC, so we have to remove it and apply after the installation.
		Setup(maistra.ApplyRestrictedRBAC(namespace.Future(&istioNamespace))).
		// We cannot disable webhooks in maistra.Install(), because then we would need maistra/istio-operator
		// to properly patch CA bundles in the webhooks. To avoid that problem we restart Istio with disabled webhooks
		// and without roles for managing webhooks once they are already created and patched.
		Setup(maistra.DisableWebhooksAndRestart(namespace.Future(&istioNamespace))).
		Run()
}

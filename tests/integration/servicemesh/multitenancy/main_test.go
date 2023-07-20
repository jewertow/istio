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

package multitenancy

import (
	"context"
	"fmt"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/util/retry"
	"istio.io/istio/tests/integration/servicemesh/maistra"
)

var (
	istioNs1 namespace.Instance
	istioNs2 namespace.Instance
)

func TestMain(m *testing.M) {
	// do not change order of setup functions
	// nolint: staticcheck
	framework.
		NewSuite(m).
		RequireMaxClusters(1).
		Setup(maistra.ApplyServiceMeshCRDs).
		Setup(namespace.Setup(&istioNs1, namespace.Config{Prefix: "istio-system"})).
		Setup(namespace.Setup(&istioNs2, namespace.Config{Prefix: "istio-system"})).
		Setup(maistra.Install(namespace.Future(&istioNs1))).
		Setup(maistra.Install(namespace.Future(&istioNs2))).
		// We cannot apply restricted RBAC before the control plane installation, because the operator always applies
		// the default RBAC, so we have to remove it and apply after the installation.
		Setup(maistra.RemoveDefaultRBAC).
		Setup(maistra.ApplyRestrictedRBAC(namespace.Future(&istioNs1))).
		Setup(maistra.ApplyRestrictedRBAC(namespace.Future(&istioNs2))).
		// We cannot disable webhooks in maistra.Install(), because then we would need maistra/istio-operator
		// to properly patch CA bundles in the webhooks. To avoid that problem we restart Istio with disabled webhooks
		// and without roles for managing webhooks once they are already created and patched.
		Setup(maistra.DisableWebhooksAndRestart(namespace.Future(&istioNs1))).
		Setup(maistra.DisableWebhooksAndRestart(namespace.Future(&istioNs2))).
		Run()
}

func TestMultiTenancy(t *testing.T) {
	framework.NewTest(t).Run(func(ctx framework.TestContext) {
		kubeClient := ctx.Clusters().Default().Kube()
		for _, ns := range []string{istioNs1.Name(), istioNs2.Name()} {
			if err := retry.UntilSuccess(func() error {
				isitod, err := kubeClient.AppsV1().Deployments(ns).Get(context.TODO(), "istiod-"+ns, v1.GetOptions{})
				if err != nil {
					return fmt.Errorf("failed to get istiod/%s: %v", ns, err)
				}
				if isitod.Status.ReadyReplicas != isitod.Status.Replicas {
					return fmt.Errorf("istiod not ready: %d of %d ready", isitod.Status.ReadyReplicas, isitod.Status.Replicas)
				}
				return nil
			}); err != nil {
				ctx.Errorf("istiod/%s not ready: %v", ns, err)
			}
		}
	})
}

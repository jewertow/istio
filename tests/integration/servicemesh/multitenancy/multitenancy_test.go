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
	"istio.io/istio/pkg/test/framework/components/istio"
	"istio.io/istio/tests/integration/servicemesh"
)

var i istio.Instance

func TestMain(m *testing.M) {
	framework.
		NewSuite(m).
		RequireSingleCluster().
		Setup(istio.Setup(&i, nil)).
		Run()
}

func TestMultiTenancy(t *testing.T) {
	framework.NewTest(t).
		Run(func(ctx framework.TestContext) {
			cluster := ctx.Clusters().Default()
			bookinfoNamespace := servicemesh.CreateNamespace(ctx, cluster, "bookinfo")
			sleepNamespace := servicemesh.CreateNamespace(ctx, cluster, "sleep")
			configureMemberRollNameInIstiod(ctx, cluster)
			createServiceMeshMemberRoll(ctx, cluster, bookinfoNamespace)
			servicemesh.InstallBookinfo(ctx, cluster, bookinfoNamespace)
			servicemesh.InstallSleep(ctx, cluster, sleepNamespace)
			verifyIstioProxyConfig(ctx, cluster, bookinfoNamespace)
		})
}

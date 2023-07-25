//go:build integ
// +build integ

//
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

package maistra

import (
	"fmt"
	"strconv"

	"istio.io/istio/pkg/test/framework/components/echo"
	"istio.io/istio/pkg/test/framework/components/echo/common/ports"
	"istio.io/istio/pkg/test/framework/components/echo/deployment"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/framework/resource"
)

func DeployEchos(apps *echo.Instances, namespaces ...namespace.Getter) func(t resource.Context) error {
	return func(t resource.Context) error {
		echoBuilder := deployment.New(t).WithClusters(t.Clusters()...)
		for i, ns := range namespaces {
			echoBuilder = echoBuilder.WithConfig(echo.Config{
				Service:   fmt.Sprintf("app-%d", i+1),
				Namespace: ns.Get(),
				Ports:     ports.All(),
				Subsets: []echo.SubsetConfig{
					{
						Annotations: map[echo.Annotation]*echo.AnnotationValue{
							echo.SidecarInject: {
								Value: strconv.FormatBool(false),
							},
						},
					},
				},
			})
		}
		var err error
		*apps, err = echoBuilder.Build()
		return err
	}
}

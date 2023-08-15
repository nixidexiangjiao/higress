// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"testing"

	"github.com/alibaba/higress/test/e2e/conformance/utils/http"
	"github.com/alibaba/higress/test/e2e/conformance/utils/suite"
)

func init() {
	HigressConformanceTests = append(HigressConformanceTests, HTTPRouteConsulHttpBin)
}

var HTTPRouteConsulHttpBin = suite.ConformanceTest{
	ShortName:   "HTTPRouteConsulHttpBin",
	Description: "The Ingress in the higress-conformance-infra namespace uses the consul service registry.",
	Manifests:   []string{"tests/httproute-consul-httpbin.yaml"},
	Test: func(t *testing.T, suite *suite.ConformanceTestSuite) {
		testcases := []http.Assertion{
			{
				Request: http.AssertionRequest{
					ActualRequest: http.Request{
						Host:   "foo.com",
						Path:   "/ping",
						Method: "GET",
					},
				},
				Response: http.AssertionResponse{
					ExpectedResponseNoRequest: true,
					ExpectedResponse: http.Response{
						StatusCode: 200,
					},
				},
			},
		}
		t.Run("HTTPRoute Consul HttpBin", func(t *testing.T) {
			for _, testcase := range testcases {
				http.MakeRequestAndExpectEventuallyConsistentResponse(t, suite.RoundTripper, suite.TimeoutConfig, suite.GatewayAddress, testcase)
			}
		})
	},
}
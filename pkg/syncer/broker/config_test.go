/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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
package broker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/wangyd1988/admiral-inspur/pkg/syncer/broker"
)

var _ = Describe("Broker configuration", func() {
	When("an environment variable is requested for a valid setting", func() {
		It("should return the correct value", func() {
			Expect(broker.EnvironmentVariable("Insecure")).To(Equal("BROKER_K8S_INSECURE"))
		})
	})

	When("an environment variable is requested for an invalid setting", func() {
		It("should panic", func() {
			Expect(func() { broker.EnvironmentVariable("impossible setting") }).To(Panic())
		})
	})
})

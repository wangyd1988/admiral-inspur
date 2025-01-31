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

package broker

import (
	"github.com/wangyd1988/admiral-inspur/pkg/federate"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/dynamic"
)

func NewFederator(dynClient dynamic.Interface, restMapper meta.RESTMapper, targetNamespace,
	localClusterID string, keepMetadataField ...string,
) federate.Federator {
	return federate.NewCreateOrUpdateFederator(dynClient, restMapper, targetNamespace, localClusterID, keepMetadataField...)
}

// Copyright 2018 The Operator-SDK Authors
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

package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestBoolAnnotation(t *testing.T) {
	tests := []struct {
		input       map[string]interface{}
		expectedVal bool
		fallback    bool
		expectedOut string
		name        string
	}{
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "True",
			},
			expectedVal: true,
			name:        "base case true",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "False",
			},
			expectedVal: false,
			name:        "base case false",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "1",
			},
			expectedVal: true,
			name:        "true as int",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "0",
			},
			expectedVal: false,
			name:        "false as int",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/wrong-annotation": "true",
			},
			expectedVal: false,
			name:        "annotation not set",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "invalid",
			},
			expectedVal: false,
			name:        "invalid value",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "false",
			},
			fallback:    true,
			expectedVal: false,
			name:        "false annotation fallback true",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/upgrade-force": "",
			},
			fallback:    true,
			expectedVal: true,
			name:        "empty annotation fallback true",
		},
		{
			input: map[string]interface{}{
				"helm.sdk.operatorframework.io/invalid": "",
			},
			fallback:    true,
			expectedVal: true,
			name:        "invalid annotation fallback true",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedVal, hasBoolAnnotation(annotations(test.input), "helm.sdk.operatorframework.io/upgrade-force", test.fallback), test.name)
	}
}

func annotations(m map[string]interface{}) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"annotations": m,
			},
		},
	}
}

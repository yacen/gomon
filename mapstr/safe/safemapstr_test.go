// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package safe

import (
	"github.com/yacen/gomon/mapstr"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	m := mapstr.MapStr{
		"subMap": mapstr.MapStr{
			"a": 1,
		},
	}

	// Add new value to the top-level.
	err := Put(m, "a", "ok")
	assert.NoError(t, err)
	assert.Equal(t, mapstr.MapStr{"a": "ok", "subMap": mapstr.MapStr{"a": 1}}, m)

	// Add new value to subMap.
	err = Put(m, "subMap.b", 2)
	assert.NoError(t, err)
	assert.Equal(t, mapstr.MapStr{"a": "ok", "subMap": mapstr.MapStr{"a": 1, "b": 2}}, m)

	// Overwrite a value in subMap.
	err = Put(m, "subMap.a", 2)
	assert.NoError(t, err)
	assert.Equal(t, mapstr.MapStr{"a": "ok", "subMap": mapstr.MapStr{"a": 2, "b": 2}}, m)

	// Add value to map that does not exist.
	m = mapstr.MapStr{}
	err = Put(m, "subMap.newMap.a", 1)
	assert.NoError(t, err)
	assert.Equal(t, mapstr.MapStr{"subMap": mapstr.MapStr{"newMap": mapstr.MapStr{"a": 1}}}, m)
}

func TestPutRenames(t *testing.T) {
	assert := assert.New(t)

	a := mapstr.MapStr{}
	Put(a, "com.docker.swarm.task", "x")
	Put(a, "com.docker.swarm.task.id", 1)
	Put(a, "com.docker.swarm.task.name", "foobar")
	assert.Equal(mapstr.MapStr{"com": mapstr.MapStr{"docker": mapstr.MapStr{"swarm": mapstr.MapStr{
		"task": mapstr.MapStr{
			"id":    1,
			"name":  "foobar",
			"value": "x",
		}}}}}, a)

	// order is not important:
	b := mapstr.MapStr{}
	Put(b, "com.docker.swarm.task.id", 1)
	Put(b, "com.docker.swarm.task.name", "foobar")
	Put(b, "com.docker.swarm.task", "x")
	assert.Equal(mapstr.MapStr{"com": mapstr.MapStr{"docker": mapstr.MapStr{"swarm": mapstr.MapStr{
		"task": mapstr.MapStr{
			"id":    1,
			"name":  "foobar",
			"value": "x",
		}}}}}, b)
}

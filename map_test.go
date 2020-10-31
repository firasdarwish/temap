/*
 * Copyright (c) 2020 Firas M. Darwish ( https://firas.dev.sy )
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package temap

import (
	"testing"
	"time"
)

var tmap = New(time.Second * 10)
var expiresAt = time.Now().Add(time.Minute)

func BenchmarkTimedMap_SetTemporary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmap.SetTemporary("some key", "some value", expiresAt)
	}
}

func BenchmarkTimedMap_SetPermanent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmap.SetPermanent("some other key", "some other value")
	}
}

func BenchmarkTimedMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, ok := tmap.Get("some key")
		if !ok {
			b.Fail()
		}
	}
}

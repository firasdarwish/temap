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
	"sync"
	"time"
)

const (
	ElementDoesntExist = -1
	ElementPermanent   = 0
)

type TimedMap struct {
	tmap map[string]*element
	mu   *sync.RWMutex

	cleanerInterval   time.Duration
	cleanerTicker     *time.Ticker
	stopCleanerTicker chan bool
	stoppedCleaner    bool
}

func New(interval time.Duration) *TimedMap {
	t := &TimedMap{
		tmap:              map[string]*element{},
		mu:                &sync.RWMutex{},
		cleanerInterval:   interval,
		cleanerTicker:     time.NewTicker(interval),
		stopCleanerTicker: make(chan bool),
		stoppedCleaner:    true,
	}

	t.StartCleaner()

	return t
}

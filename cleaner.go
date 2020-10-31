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

import "time"

func (t *timedMap) StartCleaner() {
	// already running
	if t.stoppedCleaner == false {
		return
	}

	go func() {
		for {
			select {
			case <-t.cleanerTicker.C:
				t.clean()
			case <-t.stopCleanerTicker:
				//stop cleaner
				break
			}
		}
	}()
}

func (t *timedMap) StopCleaner() {
	// already stopped
	if t.stoppedCleaner {
		return
	}

	// stop the ticker
	t.cleanerTicker.Stop()

	// stop the cleaner
	go func() {
		t.stopCleanerTicker <- true
		return
	}()
}

func (t *timedMap) RestartCleanerWithInterval(interval time.Duration) {
	// stop the cleaner
	t.StopCleaner()

	// set new interval
	t.cleanerInterval = interval

	// set new ticker
	if t.cleanerTicker == nil {
		t.cleanerTicker = time.NewTicker(t.cleanerInterval)
	} else {
		t.cleanerTicker.Reset(interval)
	}

	// restart the cleaner
	t.StartCleaner()
}

func (t *timedMap) clean() {
	// skip cleaning session if map is empty
	if len(t.tmap) == 0 {
		return
	}

	// current time in unix timestamp (nanoseconds)
	now := time.Now().UnixNano()

	t.mu.Lock()
	defer t.mu.Unlock()

	for k, v := range t.tmap {
		// ExpiresAt == 0  => permanent element
		if v.ExpiresAt != ElementPermanent && now >= v.ExpiresAt {
			delete(t.tmap, k)
		}
	}
}

func (t *timedMap) CleanNow() {
	t.clean()
}

// Copyright 2016 Joel Scoble and The JoeFriday authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loadavg

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	inf, err := Get()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	checkLoad("ticker", inf, t)
	t.Logf("%#v\n", inf)
}

func TestTicker(t *testing.T) {
	tkr, err := NewTicker(time.Millisecond)
	if err != nil {
		t.Error(err)
		return
	}
	tk := tkr.(*Ticker)
	for i := 0; i < 5; i++ {
		select {
		case <-tk.Done:
			break
		case v, ok := <-tk.Data:
			if !ok {
				break
			}
			checkLoad("ticker", v, t)
		case err := <-tk.Errs:
			t.Errorf("unexpected error: %s", err)
		}
	}
	tk.Stop()
	tk.Close()
}

func checkLoad(n string, la LoadAvg, t *testing.T) {
	if la.Timestamp == 0 {
		t.Errorf("%s: expected Timestamp to be a non-zero value; got 0", n)
	}
	if la.Minute == 0 {
		t.Errorf("%s: expected Minute to be a non-zero value; got 0", n)
	}
	if la.Five == 0 {
		t.Errorf("%s: expected Five to be a non-zero value; got 0", n)
	}
	if la.Fifteen == 0 {
		t.Errorf("%s: expected Fifteen to be a non-zero value; got 0", n)
	}
	if la.Running == 0 {
		t.Errorf("%s: expected Running to be a non-zero value; got 0", n)
	}
	if la.Total == 0 {
		t.Errorf("%s: expected Total to be a non-zero value; got 0", n)
	}
	if la.PID == 0 {
		t.Errorf("%s: expected PID to be a non-zero value; got 0", n)
	}
}

func BenchmarkGet(b *testing.B) {
	var la LoadAvg
	p, err := NewProfiler()
	if err != nil {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		la, _ = p.Get()
	}
	_ = la
}

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

package version

import (
	"testing"

	v "github.com/mohae/joefriday/system/version"
)

func TestSerializeDeserialize(t *testing.T) {
	p, err := Get()
	if err != nil {
		t.Errorf("Get(): got %s, want nil", err)
		return
	}
	k, err := v.Get()
	if err != nil {
		t.Errorf("version.Get(): got %s, want nil", err)
		return
	}
	kD := Deserialize(p)
	if k.OS != kD.OS {
		t.Errorf("OS: got %s; want %s", kD.OS, k.OS)
	}
	if k.Version != kD.Version {
		t.Errorf("Version: got %s; want %s", kD.Version, k.Version)
	}
	if k.CompileUser != kD.CompileUser {
		t.Errorf("CompileUser: got %s; want %s", kD.CompileUser, k.CompileUser)
	}
	if k.GCC != kD.GCC {
		t.Errorf("GCC: got %s; want %s", kD.GCC, k.GCC)
	}
	if k.OSGCC != kD.OSGCC {
		t.Errorf("Version: got %s; want %s", kD.OSGCC, k.OSGCC)
	}
	if k.Type != kD.Type {
		t.Errorf("Version: got %s; want %s", kD.Type, k.Type)
	}
	if k.CompileDate != kD.CompileDate {
		t.Errorf("CompileDate: got %s; want %s", kD.CompileDate, k.CompileDate)
	}
	if k.Arch != kD.Arch {
		t.Errorf("Arch: got %s; want %s", kD.Arch, k.Arch)
	}
}

func BenchmarkGet(b *testing.B) {
	var tmp []byte
	p, _ := NewProfiler()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = p.Get()
	}
	_ = tmp
}

func BenchmarkSerialize(b *testing.B) {
	var tmp []byte
	p, _ := NewProfiler()
	inf, _ := p.Profiler.Get()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = Serialize(inf)
	}
	_ = tmp
}

func BenchmarkDeserialize(b *testing.B) {
	var k *v.Kernel
	p, _ := NewProfiler()
	tmp, _ := p.Get()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k = Deserialize(tmp)
	}
	_ = k
}

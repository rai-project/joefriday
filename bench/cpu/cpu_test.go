package cpu

import (
	"encoding/json"
	"testing"

	"github.com/DataDog/gohai/cpu"
	joefacts "github.com/mohae/joefriday/cpu/facts"
	joestats "github.com/mohae/joefriday/cpu/stats"
	gopsutilcpu "github.com/shirou/gopsutil/cpu"
)

func BenchmarkJoeFridayGetFacts(b *testing.B) {
	var fct *joefacts.Facts
	b.StopTimer()
	p, _ := joefacts.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fct, _ = p.Get()
	}
	_ = fct
}

func BenchmarkJoeFridayGetStats(b *testing.B) {
	var st *joestats.Stats
	b.StopTimer()
	p, _ := joestats.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		st, _ = p.Get()
	}
	_ = st
}

func BenchmarkDataDogGohaiCPU(b *testing.B) {
	type Collector interface {
		Name() string
		Collect() (interface{}, error)
	}
	var collector = &cpu.Cpu{}
	var c interface{}
	for i := 0; i < b.N; i++ {
		c, _ = collector.Collect()
	}
	_ = c
}

func BenchmarkShirouGopsutilInfoStat(b *testing.B) {
	var st []gopsutilcpu.InfoStat
	for i := 0; i < b.N; i++ {
		st, _ = gopsutilcpu.Info()
	}
	_ = st
}

func BenchmarkShirouGopsutilTimeStat(b *testing.B) {
	var st []gopsutilcpu.TimesStat
	for i := 0; i < b.N; i++ {
		st, _ = gopsutilcpu.Times(true)
	}
	_ = st
}

// These tests exist to print out the data that is collected; not to test the
// methods themselves.  Run with the -v flag.
func TestJoeFridayGetFacts(t *testing.T) {
	prof, err := joefacts.NewProfiler()
	if err != nil {
		t.Error(err)
		return
	}
	fct, err := prof.Get()
	if err != nil {
		t.Error(err)
		return
	}
	p, err := json.MarshalIndent(fct, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s\n", string(p))
}

func TestJoeFridayGetStats(t *testing.T) {
	prof, err := joestats.NewProfiler()
	if err != nil {
		t.Error(err)
		return
	}
	st, err := prof.Get()
	if err != nil {
		t.Error(err)
		return
	}
	p, err := json.MarshalIndent(st, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s\n", string(p))
}

func TestDataDogGohaiCPU(t *testing.T) {
	type Collector interface {
		Name() string
		Collect() (interface{}, error)
	}
	var collector = &cpu.Cpu{}
	c, err := collector.Collect()
	if err != nil {
		t.Error(err)
		return
	}
	p, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s\n", string(p))
}

func TestShirouGopsutilInfoStat(t *testing.T) {
	st, err := gopsutilcpu.Info()
	if err != nil {
		t.Error(err)
		return
	}
	p, err := json.MarshalIndent(st, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s\n", string(p))
}

func TestShirouGopsutilTimeStat(t *testing.T) {
	st, err := gopsutilcpu.Times(true)
	if err != nil {
		t.Error(err)
		return
	}
	p, err := json.MarshalIndent(st, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s\n", string(p))
}

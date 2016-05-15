package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mohae/benchutil"
)

const (
	Flat = "FlatBuffers"
	JSON = "JSON"
)

// flags
var (
	output         string
	format         string
	section        bool
	sectionHeaders bool
	nameSections   bool
)

func init() {
	flag.StringVar(&output, "output", "stdout", "output destination")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")
	flag.StringVar(&format, "format", "txt", "format of output")
	flag.StringVar(&format, "f", "txt", "format of output")
	flag.BoolVar(&nameSections, "namesections", false, "use group as section name: some restrictions apply")
	flag.BoolVar(&nameSections, "n", false, "use group as section name: some restrictions apply")
	flag.BoolVar(&section, "sections", false, "don't separate groups of tests into sections")
	flag.BoolVar(&section, "s", false, "don't separate groups of tests into sections")
	flag.BoolVar(&sectionHeaders, "sectionheader", false, "if there are sections, add a section header row")
	flag.BoolVar(&sectionHeaders, "h", false, "if there are sections, add a section header row")
}

func main() {
	flag.Parse()
	done := make(chan struct{})
	go benchutil.Dot(done)

	// set the output
	var w io.Writer
	var err error
	switch output {
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer w.(*os.File).Close()
	}
	// get the benchmark for the desired format
	// process the output
	var bench benchutil.Benchmarker
	switch format {
	case "csv":
		bench = benchutil.NewCSVBench(w)
	case "md":
		bench = benchutil.NewMDBench(w)
		bench.(*benchutil.MDBench).GroupAsSectionName = nameSections
	default:
		bench = benchutil.NewStringBench(w)
	}
	bench.SectionPerGroup(section)
	bench.SectionHeaders(sectionHeaders)
	// CPU
	runCPUBenchmarks(bench)

	// Disk
	runDiskBenchmarks(bench)

	// Memory
	runMemBenchmarks(bench)

	// Network
	runNetBenchmarks(bench)

	// Platform
	runPlatformBenchmarks(bench)

	fmt.Println("\ngenerating output...\n")
	err = bench.Out()
	if err != nil {
		fmt.Printf("error generating output: %s\n", err)
	}
}

func runCPUBenchmarks(bench benchutil.Benchmarker) {
	b := CPUGetFacts()
	bench.Add(b)

	b = CPUGetFactsFB()
	bench.Add(b)

	b = CPUFactsSerializeFB()
	bench.Add(b)

	b = CPUFactsDeserializeFB()
	bench.Add(b)

	b = CPUGetFactsJSON()
	bench.Add(b)

	b = CPUFactsSerializeJSON()
	bench.Add(b)

	b = CPUFactsDeserializeJSON()
	bench.Add(b)

	b = CPUGetStats()
	bench.Add(b)

	b = CPUGetStatsFB()
	bench.Add(b)

	b = CPUStatsSerializeFB()
	bench.Add(b)

	b = CPUStatsDeserializeFB()
	bench.Add(b)

	b = CPUGetStatsJSON()
	bench.Add(b)

	b = CPUStatsSerializeJSON()
	bench.Add(b)

	b = CPUStatsDeserializeJSON()
	bench.Add(b)

	b = CPUGetUtilization()
	bench.Add(b)

	b = CPUGetUtilizationFB()
	bench.Add(b)

	b = CPUUtilizationSerializeFB()
	bench.Add(b)

	b = CPUUtilizationDeserializeFB()
	bench.Add(b)

	b = CPUGetUtilizationJSON()
	bench.Add(b)

	b = CPUUtilizationSerializeJSON()
	bench.Add(b)

	b = CPUUtilizationDeserializeJSON()
	bench.Add(b)
}

func runDiskBenchmarks(bench benchutil.Benchmarker) {
	b := DiskGetStats()
	bench.Add(b)

	b = DiskGetStatsJSON()
	bench.Add(b)

	b = DiskStatsSerializeJSON()
	bench.Add(b)

	b = DiskStatsDeserializeJSON()
	bench.Add(b)

	b = DiskGetUsage()
	bench.Add(b)

	b = DiskGetUsageFB()
	bench.Add(b)

	b = DiskUsageSerializeFB()
	bench.Add(b)

	b = DiskUsageDeserializeFB()
	bench.Add(b)

	b = DiskGetUsageJSON()
	bench.Add(b)

	b = DiskUsageSerializeJSON()
	bench.Add(b)

	b = DiskUsageDeserializeJSON()
	bench.Add(b)
}

func runMemBenchmarks(bench benchutil.Benchmarker) {
	b := MemInfoGet()
	bench.Add(b)

	b = MemInfoGetFB()
	bench.Add(b)

	b = MemInfoSerializeFB()
	bench.Add(b)

	b = MemInfoDeserializeFB()
	bench.Add(b)

	b = MemInfoGetSON()
	bench.Add(b)

	b = MemInfoSerializeJSON()
	bench.Add(b)

	b = MemInfoDeserializeJSON()
	bench.Add(b)
}

func runNetBenchmarks(bench benchutil.Benchmarker) {
	b := NetInfoGet()
	bench.Add(b)

	b = NetInfoGetFB()
	bench.Add(b)

	b = NetInfoSerializeFB()
	bench.Add(b)

	b = NetInfoDeserializeFB()
	bench.Add(b)

	b = NetInfoGetSON()
	bench.Add(b)

	b = NetInfoSerializeJSON()
	bench.Add(b)

	b = NetInfoDeserializeJSON()
	bench.Add(b)

	b = NetGetUsage()
	bench.Add(b)

	b = NetGetUsageFB()
	bench.Add(b)

	b = NetUsageSerializeFB()
	bench.Add(b)

	b = NetUsageDeserializeFB()
	bench.Add(b)

	b = NetGetUsageJSON()
	bench.Add(b)

	b = NetUsageSerializeJSON()
	bench.Add(b)

	b = NetUsageDeserializeJSON()
	bench.Add(b)
}

func runPlatformBenchmarks(bench benchutil.Benchmarker) {
	b := PlatformKernelGet()
	bench.Add(b)

	b = PlatformKernelGetFB()
	bench.Add(b)

	b = PlatformKernelSerializeFB()
	bench.Add(b)

	b = PlatformKernelDeserializeFB()
	bench.Add(b)

	b = PlatformKernelGetJSON()
	bench.Add(b)

	b = PlatformKernelSerializeJSON()
	bench.Add(b)

	b = PlatformKernelDeserializeJSON()
	bench.Add(b)

	b = PlatformLoadAvgGet()
	bench.Add(b)

	b = PlatformLoadAvgGetFB()
	bench.Add(b)

	b = PlatformLoadAvgSerializeFB()
	bench.Add(b)

	b = PlatformLoadAvgDeserializeFB()
	bench.Add(b)

	b = PlatformLoadAvgGetJSON()
	bench.Add(b)

	b = PlatformLoadAvgSerializeJSON()
	bench.Add(b)

	b = PlatformLoadAvgDeserializeJSON()
	bench.Add(b)

	b = PlatformReleaseGet()
	bench.Add(b)

	b = PlatformReleaseGetFB()
	bench.Add(b)

	b = PlatformReleaseSerializeFB()
	bench.Add(b)

	b = PlatformReleaseDeserializeFB()
	bench.Add(b)

	b = PlatformReleaseGetJSON()
	bench.Add(b)

	b = PlatformReleaseSerializeJSON()
	bench.Add(b)

	b = PlatformReleaseDeserializeJSON()
	bench.Add(b)

	b = PlatformUptimeGet()
	bench.Add(b)

	b = PlatformUptimeGetFB()
	bench.Add(b)

	b = PlatformUptimeSerializeFB()
	bench.Add(b)

	b = PlatformUptimeDeserializeFB()
	bench.Add(b)

	b = PlatformUptimeGetJSON()
	bench.Add(b)

	b = PlatformUptimeSerializeJSON()
	bench.Add(b)

	b = PlatformUptimeDeserializeJSON()
	bench.Add(b)
}

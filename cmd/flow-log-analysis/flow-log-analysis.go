package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kushal2705/flow-log-analysis/internal/output"
	"github.com/kushal2705/flow-log-analysis/internal/parser"
)

func main() {
	lookupFile := flag.String("lookup", "", "Path to the lookup table CSV file")
	logFile := flag.String("log", "", "Path to the flow log file")
	outputDir := flag.String("output", "output", "Directory to store output files")
	flag.Parse()

	if *lookupFile == "" || *logFile == "" {
		log.Fatal("Both lookup file and log file must be specified")
	}

	lookupTable, err := parser.ReadLookupTable(*lookupFile)
	if err != nil {
		log.Fatalf("Error reading lookup table: %v", err)
	}

	tagCounts, portProtocolCounts, untaggedCount, err := parser.ParseFlowLogs(*logFile, lookupTable)
	if err != nil {
		log.Fatalf("Error parsing flow logs: %v", err)
	}

	err = output.WriteOutput(tagCounts, portProtocolCounts, untaggedCount, *outputDir)
	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}

	fmt.Println("Analysis complete. Output files written to:", *outputDir)
}

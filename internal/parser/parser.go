package parser

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

// key value map
type LookupTable map[string]string

func ReadLookupTable(filepath string) (LookupTable, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	// this will ensure that the file is close on the exit of this func
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	lookupTable := make(LookupTable)
	for _, record := range records[1:] { // Skip header
		// preparing the key with destport and protocol
		key := strings.ToLower(record[0] + "," + record[1])
		// storing tag value for the prepared key
		lookupTable[key] = strings.ToLower(record[2])
	}

	return lookupTable, nil
}

func ParseFlowLogs(logFilepath string, lookupTable LookupTable) (map[string]int, map[string]int, int, error) {
	file, err := os.Open(logFilepath)
	if err != nil {
		return nil, nil, 0, err
	}
	defer file.Close()

	tagCounts := make(map[string]int)
	portProtocolCounts := make(map[string]int)
	untaggedCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 14 {
			continue // Skip malformed lines
		}
		// all fields are referenced from https://docs.aws.amazon.com/vpc/latest/userguide/flow-log-records.html
		if fields[0] != "2" {
			continue // version 2 only supported
		}
		dstPort := fields[6]
		protocol := protocolNumberToName(fields[8])

		//prepare key for lookup
		key := strings.ToLower(dstPort + "," + protocol)
		// fetch the tag and accordingly increment the count
		if tag, ok := lookupTable[key]; ok {
			tagCounts[tag]++
		} else {
			untaggedCount++
		}

		portProtocolCounts[key]++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, 0, err
	}

	return tagCounts, portProtocolCounts, untaggedCount, nil
}

func protocolNumberToName(number string) string {
	// reference https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	switch number {
	case "1":
		return "ICMP"
	case "5":
		return "ST"
	case "6":
		return "TCP"
	case "7":
		return "CBT"
	case "8":
		return "EGP"
	case "10":
		return "BBN-RCC-MON"
	case "12":
		return "PUP"
	case "15":
		return "XNET"
	case "17":
		return "UDP"
	case "18":
		return "MUX"
	case "25":
		return "LEAF-1"
	default:
		return number
	}
}

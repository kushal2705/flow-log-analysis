package output

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func WriteOutput(tagCounts map[string]int, portProtocolCounts map[string]int, untaggedCount int, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	// This func will write Tagged and untagged counts to tag_counts.txt
	if err := writeTagUntaggedCounts(tagCounts, untaggedCount, outputDir); err != nil {
		return err
	}

	if err := writePortProtocolCounts(portProtocolCounts, outputDir); err != nil {
		return err
	}

	return nil
}

func writeTagUntaggedCounts(tagCounts map[string]int, untaggedCount int, outputDir string) error {
	file, err := os.Create(filepath.Join(outputDir, "tag_counts.txt"))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"Tag", "Count"}); err != nil {
		return err
	}

	var rows [][]string
	// Writing Tagged and Untagged count
	for tag, count := range tagCounts {
		rows = append(rows, []string{tag, strconv.Itoa(count)})
	}
	rows = append(rows, []string{"Untagged", strconv.Itoa(untaggedCount)})

	// sort the slice based on
	//sort.Slice(rows, func(i, j int) bool {
	//	return rows[i][0] < rows[j][0]
	//})

	return writer.WriteAll(rows)
}

func writePortProtocolCounts(portProtocolCounts map[string]int, outputDir string) error {
	file, err := os.Create(filepath.Join(outputDir, "port_protocol_counts.txt"))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"Port", "Protocol", "Count"}); err != nil {
		return err
	}

	var rows [][]string
	for key, count := range portProtocolCounts {
		parts := strings.Split(key, ",") //check this
		rows = append(rows, []string{parts[0], parts[1], strconv.Itoa(count)})
	}

	/*sort.Slice(rows, func(i, j int) bool {
		if rows[i][0] == rows[j][0] {
			return rows[i][1] < rows[j][1]
		}
		return rows[i][0] < rows[j][0]
	})*/

	return writer.WriteAll(rows)
}

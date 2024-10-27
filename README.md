# Flow Log Analysis

## Purpose
This project is designed to analyze AWS VPC Flow Logs. It parses flow log data and maps each log entry to a tag based on a provided lookup table. The program processes log files and generates statistical output about the analyzed logs.

## Directory Structure

flow-log-analysis/
├── cmd/
│ └── flow-log-analysis/
│ └── flow-log-analysis.go
├── internal/
│ ├── parser/
│ │ └── parser.go
│ └── output/
│ └── output.go
├── data/
│ ├── lookup.csv
│ └── flowlog.txt
├── go.mod
└── README.md

## Libraries Used
This project uses only Go standard libraries:
- `bufio`: For efficient file reading
- `encoding/csv`: For CSV file operations
- `flag`: For command-line flag parsing
- `fmt`: For formatted I/O
- `os`: For file system operations
- `path/filepath`: For file path manipulations
- `sort`: For sorting slices
- `strconv`: For string conversions
- `strings`: For string operations

## Assumptions
1. Each flow log entry should have 14 fields, which is the default log format. Entries with fewer fields will be skipped. 
    - Reference : https://docs.aws.amazon.com/vpc/latest/userguide/flow-log-records.html
2. Only Version 2 log entries will be parsed. The program checks the first field of each log entry to ensure it's "2".
3. The lookup table is a CSV file with columns: dstport, protocol, and tag.

## Installation and Setup

### Installing Go
1. Visit the [official Go download page](https://golang.org/dl/).
2. Download the installer for your operating system.
3. Follow the installation instructions for your OS.
4. Verify the installation by opening a terminal and running: `go version`

### Building and Running the Program
1. Clone this repository:
git clone https://github.com/yourusername/flow-log-analysis.git
cd flow-log-analysis

2. Build the program:
go build -o flow-analyzer ./cmd/flow-log-analysis

3. Run the program:
./flow-analyzer -lookup ./data/lookup.txt -log ./data/flowlog.txt -output .

## Output
The program generates two CSV files in the specified output directory:
1. `tag_counts.txt`: Contains the count of each tag found in the log file.
2. `port_protocol_counts.txt`: Contains the count of each unique port/protocol combination.

## Testing and Validation
The following tests and validations were performed:

1. Log entry with fewer than 14 fields:
- Added a log entry with only 13 fields.
- Result: The entry was skipped, and a warning was logged.

2. Log entry with version 3 instead of version 2:
- Added a log entry starting with "3" instead of "2".
- Result: The entry was skipped, and a warning was logged.

3. Empty log file:
- Tested with an empty log file.
- Result: The program completed successfully with empty output files.

4. Case sensitivity:
- Tested with mixed case in log entries and lookup table.
- Result: Matching was performed case-insensitively as required.

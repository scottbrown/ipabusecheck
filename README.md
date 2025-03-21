# ipabusecheck

![go](https://img.shields.io/badge/language-go-blue)
![security](https://img.shields.io/badge/category-security-blue)
[![.github/workflows/ci.yaml](https://github.com/scottbrown/ipabusecheck/actions/workflows/ci.yaml/badge.svg)](https://github.com/scottbrown/ipabusecheck/actions/workflows/ci.yaml)
[![SAST](https://github.com/scottbrown/ipabusecheck/actions/workflows/sast.yaml/badge.svg)](https://github.com/scottbrown/ipabusecheck/actions/workflows/sast.yaml)
![version](https://img.shields.io/badge/version-0.1-blue)

A command-line tool that checks IP addresses against the AbuseIPDB API to retrieve abuse scores and reports.

## Overview

`ipabusecheck` allows you to quickly check a list of IP addresses against the [AbuseIPDB](https://www.abuseipdb.com/) database to identify potentially malicious IP addresses. The tool processes a list of IP addresses from a file, queries the AbuseIPDB API for each one, and outputs the results to a CSV file.

## Features

- Batch processing of multiple IP addresses
- CSV output format for easy analysis
- Progress bar for tracking query status
- Secure API key handling via environment variables

## Installation

### Pre-built Binaries

Download the latest release from the [Releases page](https://github.com/scottbrown/ipabusecheck/releases).

### Build from Source

Requirements:
- Go 1.20+

```bash
# Clone the repository
git clone https://github.com/scottbrown/ipabusecheck.git
cd ipabusecheck

# Build the application
make build

# The binary will be available in .build/ipabusecheck
```

## Usage

First, set your AbuseIPDB API key as an environment variable:

```bash
export ABUSEIPDB_API_KEY="your_api_key_here"
```

Then run the tool with your input file:

```bash
ipabusecheck --input ips.txt --output results.csv
```

### Command-line Options

```
Usage:
  ipabusecheck [flags]

Flags:
  -i, --input string    The list of IP addresses to check (one per line)
  -o, --output string   The file to write the data to
  -s, --silent          Do not print statistics, just output to the file
```

## Output Format

The output CSV file contains the following columns:

1. IP Address
2. Total Reports - Number of abuse reports for the IP
3. Confidence Score - AbuseIPDB confidence score (0-100)

Example output:
```csv
192.168.1.1,0,0
8.8.8.8,0,0
1.2.3.4,27,90
```

## Development

### Project Structure

```
.
├── cmd/                  # Command-line application code
│   ├── flags.go          # CLI flag definitions
│   ├── main.go           # Application entry point
│   └── rootCmd.go        # Main command implementation
├── fixtures/             # Test fixtures with sample API responses
├── .github/workflows/    # CI/CD pipeline configurations
├── checker.go            # Core IP checking functionality
├── checker_test.go       # Tests for checker
├── constants.go          # Application constants
├── go.mod                # Go module definition
└── Makefile              # Build tasks
```

### Running Tests

```bash
make test
```

### Security Checks

```bash
make check  # Runs all security checks
make sast   # Runs static analysis security testing
make vet    # Runs Go's built-in code analyzer
make vuln   # Runs vulnerability checking
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

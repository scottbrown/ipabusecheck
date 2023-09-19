package main

import (
	"bufio"
  "encoding/csv"
	"os"
  "strconv"

	"github.com/scottbrown/ipabusecheck"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipabusecheck",
	Short: "Checks for IP Abuse scores from Abuse IP DB",
	Long:  "Provides the number of abuse reports for each IP address in a given list, without stdout containing the ones with the most abuse. Set your API key in an environment variable named ABUSEIPDB_API_KEY.",
	RunE:  handleRoot,
}

func handleRoot(cmd *cobra.Command, args []string) error {
	checker := ipabusecheck.Checker{
		ApiKey: os.Getenv("ABUSEIPDB_API_KEY"),
	}

	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var ipAddresses []string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		ipAddress := scanner.Text()

		ipAddresses = append(ipAddresses, ipAddress)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	var reports []ipabusecheck.Report
	bar := progressbar.Default(int64(len(ipAddresses)))
	for _, ipAddress := range ipAddresses {
		bar.Add(1)
		report, err := checker.Check(ipAddress)
		if err != nil {
			return err
		}

		reports = append(reports, report)
	}

	outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	for _, report := range reports {
    iTotalReports := strconv.FormatInt(report.TotalReports, 10)
    iConfidenceScore := strconv.FormatInt(report.ConfidenceScore, 10)
		if err := writer.Write([]string{report.IPAddress, iTotalReports, iConfidenceScore,}); err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

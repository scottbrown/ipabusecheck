package ipabusecheck

import (
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
)

type Report struct {
	IPAddress       string
	TotalReports    int64
	ConfidenceScore int64
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Checker struct {
	Client HttpClient
	ApiKey string
}

func (c Checker) Check(ipAddress string) (Report, error) {
	url := fmt.Sprintf("%s?ipAddress=%s", apiEndpoint, ipAddress)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Report{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Key", os.Getenv(envApiKey))

	resp, err := client.Do(req)
	if err != nil {
		return Report{}, err
	}
	/* #nosec G307 */
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Report{}, fmt.Errorf("API request failed. Status code: %d, message: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Report{}, err
	}

	totalReports, err := jsonparser.GetInt(body, dataKey, totalReportsKey)
	if err != nil {
		return Report{}, err
	}

	abuseConfidenceScore, err := jsonparser.GetInt(body, dataKey, abuseConfidenceScoreKey)
	if err != nil {
		return Report{}, err
	}

	report := Report{}
	report.IPAddress = ipAddress
	report.TotalReports = totalReports
	report.ConfidenceScore = abuseConfidenceScore

	return report, nil
}

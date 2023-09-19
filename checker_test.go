package ipabusecheck

import (
	"embed"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

//go:embed fixtures/*.json
var f embed.FS

type mockSuccessHttpClient struct{}

func (m *mockSuccessHttpClient) Do(req *http.Request) (*http.Response, error) {
	mockResponse, _ := f.ReadFile("success.json")
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(string(mockResponse))),
	}, nil
}

type mockFailureHttpClient struct{}

func (m *mockFailureHttpClient) Do(req *http.Request) (*http.Response, error) {
	mockResponse, _ := f.ReadFile("failure.json")
	return &http.Response{
		StatusCode: 400,
		Body:       ioutil.NopCloser(strings.NewReader(string(mockResponse))),
	}, nil
}

func TestSuccess(t *testing.T) {
	c := Checker{
		Client: &mockSuccessHttpClient{},
	}

	report, err := c.Check("1.2.3.4")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if report.TotalReports != 1 {
		t.Errorf("Expected %d but got %d", 1, report.TotalReports)
	}

	if report.ConfidenceScore != 100 {
		t.Errorf("Expected %d but got %d", 100, report.ConfidenceScore)
	}

	if report.IPAddress != "1.2.3.4" {
		t.Errorf("Expected %s but got %s", "1.2.3.4", report.IPAddress)
	}
}

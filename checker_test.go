package ipabusecheck

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

//go:embed fixtures/*.json
var f embed.FS

type mockHttpClient struct {
	fixtureFilename  string
	mockResponseCode int
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mockResponse, _ := f.ReadFile(fmt.Sprintf("fixtures/%s.json", m.fixtureFilename))
	return &http.Response{
		StatusCode: m.mockResponseCode,
		Body:       io.NopCloser(strings.NewReader(string(mockResponse))),
	}, nil
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name            string
		checker         Checker
		ipAddress       string
		expectedReports int64
		expectedScore   int64
		expectedError   error
	}{
		{
			name:            "success",
			checker:         Checker{Client: &mockHttpClient{mockResponseCode: 200, fixtureFilename: "success"}},
			ipAddress:       "1.2.3.4",
			expectedReports: 1,
			expectedScore:   100,
			expectedError:   nil,
		},
		{
			name:            "failure",
			checker:         Checker{Client: &mockHttpClient{mockResponseCode: 400, fixtureFilename: "failure"}},
			ipAddress:       "",
			expectedReports: 0,
			expectedScore:   0,
			expectedError:   errors.New("API request failed. Status code: 400, message: Bad Request"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report, err := test.checker.Check(test.ipAddress)

			if err != nil && test.expectedError != nil {
				if err.Error() != test.expectedError.Error() {
					t.Errorf("Expected %v but got %v", test.expectedError, err)
				}
			} else if err != test.expectedError {
				t.Errorf("Expected %v but got %v", test.expectedError, err)
			}

			if report.TotalReports != test.expectedReports {
				t.Errorf("Expected %d but got %d", test.expectedReports, report.TotalReports)
			}

			if report.ConfidenceScore != test.expectedScore {
				t.Errorf("Expected %v but got %v", test.expectedScore, report.ConfidenceScore)
			}

			if report.IPAddress != test.ipAddress {
				t.Errorf("Expected %s but got %s", test.ipAddress, report.IPAddress)
			}
		})
	}
}

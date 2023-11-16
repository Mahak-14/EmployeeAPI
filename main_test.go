package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/gofr/request"
)

func Test_Integration(t *testing.T) {
	go main()

	time.Sleep(5 * time.Second)

	employeeCreateBody := []byte(`{"id": 1, "name": "test", "dept": "test-dept"}`)
	employeeUpdateBody := []byte(`{"id": 1, "name": "test-name", "dept": "test-dept"}`)

	successResp := `{"data":{"ID":1,"name":"test","dept":"test-dept"}}`
	successUpdateResp := `{"data":{"ID":1,"name":"test-name","dept":"test-dept"}}`

	testCases := []struct {
		desc          string
		method        string
		endpoint      string
		body          []byte
		expStatusCode int
		expResp       string
	}{
		{"Create employee", http.MethodPost, "/employee", employeeCreateBody, http.StatusCreated,
			successResp},
		{"Get employee", http.MethodGet, "/employee/1", nil, http.StatusOK, successResp},
		{"Update employee", http.MethodPut, "/employee/1", employeeUpdateBody, http.StatusOK,
			successUpdateResp},
		{"Delete employee", http.MethodDelete, "/employee/1", nil, http.StatusNoContent, ``},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			req, _ := request.NewMock(tc.method, "http://localhost:9000"+tc.endpoint, bytes.NewBuffer(tc.body))
			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error occurred in calling api: %v", err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error while reading response: %v", err)
			}

			respBody := strings.TrimSpace(string(body))

			assert.Equal(t, tc.expStatusCode, resp.StatusCode, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expResp, respBody, "Test [%d] failed", i+1)

			resp.Body.Close()
		})
	}
}

package github

import (
	"errors"
	"net/http"
	"testing"
)

type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestRequestFailureInvalidURL(t *testing.T) {
	invalidBaseURL := "://invalid_url"
	client := &Client{
		BaseURL: invalidBaseURL,
	}
	_, err := client.Request(http.MethodGet, "/path")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestRequestFailureInvalidPath(t *testing.T) {
	client := &Client{}

	invalidPath := "://"

	_, err := client.Request(http.MethodGet, invalidPath)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestRequestFailureInvalidMethod(t *testing.T) {
	client := &Client{}

	invalidMethod := "INVALID METHOD"
	_, err := client.Request(invalidMethod, "/path")

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestRequestSuccess(t *testing.T) {
	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
			}, nil
		},
	}

	client := &Client{
		httpClient: mockClient,
	}

	_, err := client.Request(http.MethodGet, "/path")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRequestFailure(t *testing.T) {
	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("error")
		},
	}

	client := &Client{
		httpClient: mockClient,
	}

	_, err := client.Request(http.MethodGet, "/path")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

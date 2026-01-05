package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


type Client struct{
	BaseURL string
	HttpClient *http.Client
}


func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HttpClient: &http.Client{
			Timeout: 60 * time.Second, // Set timeout 
		},
	}
}


func (c *Client) Generate(ctx context.Context, req GenerateRequest)(*GenerateResponse, error){
	// Marshalling: Converting our Go struct to JSON
	jsonData, err := json.Marshal(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshall request: %w", err)
	}

	// creating http request with context (allows cancellation)
	httpReq,err := http.NewRequestWithContext(ctx,"POST",c.BaseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Failed to create http request: %w", err)
	}

	httpReq.Header.Set("Content-Type","application/json")

	res, err := c.HttpClient.Do(httpReq)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-200 response: %d", res.StatusCode)
	}

	// unmarshalling: Converting JSON response to Go struct
	var result GenerateResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode response: %w", err)
	}

	return &result, nil
}
package resty

import (
	"net/http"
	"time"
)

// Client represents a client for making HTTP requests
 type Client struct {
	Timeout time.Duration
 }

// Do sends an HTTP request and returns an HTTP response
 func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	req = req.WithContext(http.NewRequest(req.Context(), req, nil))
	req = req.WithContext(http.TimeoutContext(req.Context(), c.Timeout))
	return http.DefaultClient.Do(req)
 }

// NewClient creates a new client with the specified timeout
 func NewClient(timeout time.Duration) *Client {
	return &Client{Timeout: timeout}
 }
package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DoHttp(config AxiosRequestConfig) AxiosResponse {
	resp, _ := doRequest(config)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)
	headers := make(map[string]string)
	for key, value := range resp.Header {
		headers[key] = strings.Join(value, " ")
	}
	return AxiosResponse{Status: resp.StatusCode, StatusText: resp.Status, Headers: headers, Config: config, Data: bodyStr}
}

type AxiosRequestConfig struct {
	URL     string            `json:"url,omitempty"`
	Method  string            `json:"method,omitempty"`
	BaseURL string            `json:"baseURL,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Auth    map[string]string `json:"auth,omitempty"`
	Params  interface{}       `json:"params,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
}

type AxiosResponse struct {
	Status     int                `json:"status"`
	StatusText string             `json:"statusText"`
	Headers    map[string]string  `json:"headers"`
	Config     AxiosRequestConfig `json:"config"`
	Data       interface{}        `json:"data"`
}

func doRequest(config AxiosRequestConfig) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Build URL with parameters if any
	fullUrl := config.BaseURL + config.URL
	if config.Params != nil {
		params := url.Values{}
		for key, value := range config.Params.(map[string]interface{}) {
			if val, ok := value.(string); ok {
				params.Add(key, val)
			}
		}
		fullUrl += "?" + params.Encode()
	}

	// Ceate request
	req, err := http.NewRequest(config.Method, fullUrl, nil)
	if err != nil {
		return nil, err
	}
	if len(config.Auth) > 0 {
		baseAuth := config.Auth["username"] + ":" + config.Auth["password"]
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(baseAuth)))
	}
	if len(config.Headers) > 0 {
		for key, value := range config.Headers {
			req.Header.Add(key, value)
		}
	}
	// Set request body if any
	if config.Data != nil {
		data, err := json.Marshal(config.Data)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(data))
		req.Body = io.NopCloser(bytes.NewBuffer(data))
	}
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

package openhab

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Client for openHAB. It's using openHAB REST API internally.
type Client struct {
	config     *Config
	httpClient *http.Client
	token      string
}

func New(config *Config) *Client {

	config = config.Clone()

	// API token takes precedence over user/password
	if config.APIToken != "" {
		config.User = config.APIToken
		config.Password = ""
	}

	return &Client{
		config:     config,
		httpClient: &http.Client{},
	}
}

func (t *Client) GetThings(ctx context.Context) (*Things, error) {
	resource := "/rest/things"
	params := url.Values{}
	params.Add("staticDataOnly", "false")
	u, _ := url.ParseRequestURI(t.config.API)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	result := &Things{}
	err := t.getJSON(ctx, urlStr, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *Client) auth(ctx context.Context) error {

	if t.token != "" {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.config.API, http.NoBody)
	if err != nil {
		return err
	}

	req.SetBasicAuth(t.config.User, t.config.Password)

	// checkov:skip=CKV_SECRET_6: This is not a secret
	token := req.Header.Get("Authorization")

	if token == "" {
		return fmt.Errorf("received empty token")
	}

	t.token = token

	return nil
}

func (c *Client) getJSON(ctx context.Context, url string, result interface{}) error {
	resp, err := c.get(ctx, url, "application/json")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (t *Client) get(ctx context.Context, url, contentType string) (*http.Response, error) {

	err := t.auth(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Authorization", t.token)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode >= 400 {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return resp, fmt.Errorf("not found")
		default:
			return resp, errors.New(resp.Status)
		}
	}

	return resp, nil
}

// func (c *Client) getString(ctx context.Context, url string) (string, error) {
// 	resp, err := c.get(ctx, url, "text/plain")
// 	if resp != nil {
// 		defer resp.Body.Close()
// 	}
// 	if err != nil {
// 		return "", err
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(body), nil
// }

// func (c *Client) postString(ctx context.Context, url, value string) error {
// 	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+url, strings.NewReader(value))
// 	if err != nil {
// 		return err
// 	}
// 	req.SetBasicAuth(c.user, c.password)
// 	req.Header.Set("Content-Type", "text/plain")

// 	resp, err := c.client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	// we don't expect any body in the response
// 	resp.Body.Close()

// 	if resp.StatusCode >= http.StatusBadRequest {
// 		switch resp.StatusCode {
// 		case http.StatusNotFound:
// 			return ErrorNotFound
// 		default:
// 			return errors.New(resp.Status)
// 		}
// 	}

// 	return nil
// }

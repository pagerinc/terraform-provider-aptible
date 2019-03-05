package aptible

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/hashicorp/terraform/helper/logging"
)

const (
	DefaultPostAppCreateDelay = int64(5)
	AptibleAPIURL             = "https://api.aptible.com"
)

type Client struct {
	APIKey      string
	DebugHTTP   bool
	Email       string
	Environment string
	Headers     http.Header
	URL         string
}

func (c Client) String() string {
	return fmt.Sprintf("{APIKey:xxx Email:%s URL:%s Headers:xxx DebugHTTP:%t}",
		c.Email, c.URL, c.DebugHTTP)
}

func NewClient() *Client {
	client := &Client{
		Headers: make(http.Header),
		URL:     AptibleAPIURL,
	}
	if logging.IsDebugOrHigher() {
		client.DebugHTTP = true
	}
	return client
}

func (c *Client) Login() error {
	args := fmt.Sprintf("/usr/local/bin/aptible login --lifetime 10m --email '%s' --password '%s'", c.Email, c.APIKey)
	return exec.Command("sh", "-c", args).Run()
}

func (c *Client) GetApp(id int) ([]byte, error) {
	args := fmt.Sprintf("/usr/local/bin/aptible inspect %s/apps/%d", c.URL, id)
	return exec.Command("sh", "-c", args).Output()
}

func (c *Client) GetAppByHandle(handle string) (*App, error) {
	args := fmt.Sprintf("/usr/local/bin/aptible apps --environment %s", c.Environment)
	blob, err := exec.Command("sh", "-c", args).Output()
	if err != nil {
		return nil, err
	}

	apps, err := NewAppList(blob)
	if err != nil {
		return nil, err
	}
	for i := range apps {
		if apps[i].Handle == handle {
			return &apps[i], nil
		}
	}

	return nil, nil
}

package adpexpert

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	DefaultBaseURL = "https://expert.brasil.adp.com/"
)

type Client struct {
	client    *resty.Client
	sessionID string
	contextID string

	Debug    bool
	BaseURL  string
	LoginURL string
}

func (c *Client) newRequest() *resty.Request {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(errors.Wrap(err, "could not create a new request"))
	}

	if c.client == nil {
		c.client = resty.New().
			SetDebug(c.Debug).
			SetBaseURL(coalesce(c.BaseURL, DefaultBaseURL)).
			SetCookieJar(jar).
			SetRedirectPolicy(resty.RedirectPolicyFunc(func(r *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}))
	}
	return c.client.NewRequest()
}

func (c *Client) Clear() {
	c.client = nil
}

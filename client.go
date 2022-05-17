package adpexpert

type Client struct {
	LoginURL string
}

const (
	BaseURL = "https://expert.brasil.adp.com/"
)

func (c Client) getLoginURL() string {
	if c.LoginURL != "" {
		return c.LoginURL
	}
	return DefaultLoginURL
}

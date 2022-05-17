package adpexpert

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func (c *Client) Login(username, password string) error {
	target := coalesce(c.BaseURL, DefaultBaseURL)
	resp, err := c.newRequest().
		SetFormData(map[string]string{
			"USER":     username,
			"PASSWORD": password,
			"TARGET":   target,
		}).
		Post("ipclogin/1/loginform.fcc")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 302 {
		return fmt.Errorf("login failed: invalid status: %s", resp.Status())
	}
	location, err := resp.RawResponse.Location()
	if location == nil || location.String() != target {
		return fmt.Errorf("login failed: invalid redirect location: %s", location)
	}

	resp, err = c.newRequest().
		SetDoNotParseResponse(true).
		Get("/expert/v4/?lp=true")
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromResponse(resp.RawResponse)
	if err != nil {
		return errors.Wrap(err, "failed to parse home page")
	}

	sessionID, exists := doc.Find("input#newexpert_sessionid").First().Attr("value")
	if !exists {
		return fmt.Errorf("failed to find session ID in the home page")
	}
	c.sessionID = sessionID
	return nil
}

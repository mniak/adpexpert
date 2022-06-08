package adpexpert

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/mniak/adpexpert/models"
	"github.com/pkg/errors"
)

func (c *Client) ensureLoggedIn() error {
	if c.sessionID == "" {
		return errors.New("not logged in")
	}
	return nil
}

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
	if !resp.IsSuccess() {
		return fmt.Errorf("login failed: home page returned status %s", resp.Status())
	}

	doc, err := goquery.NewDocumentFromResponse(resp.RawResponse)
	if err != nil {
		return errors.Wrap(err, "failed to parse home page")
	}

	sessionID, exists := doc.Find("input#newexpert_sessionid").First().Attr("value")
	if !exists {
		return fmt.Errorf("login failed: could not find session ID in the home page")
	}
	c.sessionID = sessionID

	resp, err = c.newRequest().
		SetHeader("newexpert_sessionid", sessionID).
		SetResult(models.Context{}).
		Get("/expert/api/contextselection/default/usercontext/tree")
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("login failed: user context request returned status %s", resp.Status())
	}

	ctx := resp.Result().(*models.Context)
	if len(ctx.Contexts) == 0 {
		return errors.New("login failed: user context response has 0 items")
	}
	ctx0 := ctx.Contexts[0]
	c.contextID = ctx0.ContextID
	return nil
}

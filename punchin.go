package adpexpert

import (
	"errors"
	"fmt"
)

func (c *Client) PunchIn() error {
	if c.sessionID == "" {
		return errors.New("not logged in")
	}

	resp, err := c.newRequest().
		SetHeader("newexpert_sessionid", c.sessionID).
		SetBody(map[string]any{
			"punchType":      "SPDesktop",
			"punchLatitude":  nil,
			"punchLongitude": nil,
			"punchAction":    nil,
		}).
		// POST, but for testing I am still using Get in order to avoid punching in many times
		Get("/expert/api/punch/punchin?lp=true")
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("punch-in failed with invalid status: %s", resp.Status())
	}
	return err
}

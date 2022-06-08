package adpexpert

import (
	"errors"
	"fmt"
)

func (c *Client) GetTimecard() error {
	if c.sessionID == "" {
		return errors.New("not logged in")
	}

	resp, err := c.newRequest().
		SetHeader("newexpert_sessionid", c.sessionID).
		SetHeader("serviceplace-context-contextid", c.contextID).
		SetBody(map[string]interface{}{
			"punchType":      "SPDesktop",
			"punchLatitude":  nil,
			"punchLongitude": nil,
			"punchAction":    nil,
		}).
		Post("/expert/api/punch/punchin?lp=true")
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("punch-in failed with invalid status: %s", resp.Status())
	}
	return err
}

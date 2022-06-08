package adpexpert

import (
	"fmt"

	"github.com/mniak/adpexpert/models"
)

func (c *Client) PunchIn() error {
	if err := c.ensureLoggedIn(); err != nil {
		return err
	}

	resp, err := c.newRequest().
		SetHeader("newexpert_sessionid", c.sessionID).
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

type LastPunchesInfo models.PunchesResponse

func (c *Client) GetLastPunches() (*LastPunchesInfo, error) {
	if err := c.ensureLoggedIn(); err != nil {
		return nil, err
	}

	resp, err := c.newRequest().
		SetHeader("newexpert_sessionid", c.sessionID).
		SetBody(map[string]interface{}{
			"punchType":      "SPDesktop",
			"punchLatitude":  nil,
			"punchLongitude": nil,
			"punchAction":    nil,
		}).
		SetResult(LastPunchesInfo{}).
		Get("/expert/api/punch/punchin?lp=true")
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("punch-in failed with invalid status: %s", resp.Status())
	}
	result := resp.Result().(*LastPunchesInfo)
	return result, nil
}

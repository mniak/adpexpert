package adpexpert

import (
	"errors"
	"fmt"

	"github.com/mniak/adpexpert/models"
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
	if c.sessionID == "" {
		return nil, errors.New("not logged in")
	}

	resp, err := c.newRequest().
		SetHeader("newexpert_sessionid", c.sessionID).
		SetBody(map[string]any{
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

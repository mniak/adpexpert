package adpexpert

import (
	"errors"
	"fmt"

	"github.com/mniak/adpexpert/models"
)

func (c *Client) GetTimecard(year, month int) (*models.Timecard, error) {
	if c.sessionID == "" {
		return nil, errors.New("not logged in")
	}

	resp, err := c.newRequest().
		SetHeader("newexpert_sessionid", c.sessionID).
		SetHeader("serviceplace-context-contextid", c.contextID).
		SetResult(&models.Timecard{}).
		Get(fmt.Sprintf("/expert/api/timesheet/time-card/reference/%d/%02d", year, month))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to load timecard invalid status: %s", resp.Status())
	}

	return resp.Result().(*models.Timecard), nil
}

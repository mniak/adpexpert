package models

import "time"

type PunchesResponse struct {
	PunchDisabled             bool      `json:"punchDisabled"`
	EnableEntranceExitButtons bool      `json:"enableEntranceExitButtons"`
	ServerTime                time.Time `json:"serverTime"`
	LastPunches               []Punch   `json:"lastPunches"`
}

type Punch struct {
	ID             string      `json:"_id"`
	OrgOID         string      `json:"orgoid"`
	AssociateOID   string      `json:"associateoid"`
	EmployeeKey    string      `json:"employeeKey"`
	Status         string      `json:"status"`
	PunchDateTime  time.Time   `json:"punchDateTime"`
	PunchTimezone  string      `json:"punchTimezone"`
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
	V              int64       `json:"__v"`
	PunchType      *string     `json:"punchType,omitempty"`
	PunchLatitude  interface{} `json:"punchLatitude"`
	PunchLongitude interface{} `json:"punchLongitude"`
	PunchAction    interface{} `json:"punchAction"`
}

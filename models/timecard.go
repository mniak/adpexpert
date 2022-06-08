package models

type Timecard struct {
	Employee    TimecardEmployee  `json:"employee"`
	Reference   TimecardReference `json:"reference"`
	Timetable   []TimetableEntry  `json:"timetable"`
	Validation  Validation        `json:"validation"`
	TimeEnabled bool              `json:"timeEnabled"`
}

type TimecardEmployee struct {
	Ooid             string `json:"ooid"`
	Aoid             string `json:"aoid"`
	EmployeeKey      string `json:"employeeKey"`
	EmployeeID       int64  `json:"employeeId"`
	EmployeeName     string `json:"employeeName"`
	CompanyID        int64  `json:"companyId"`
	CompanyCode      string `json:"companyCode"`
	CompanyName      string `json:"companyName"`
	LocationCode     string `json:"locationCode"`
	LocationName     string `json:"locationName"`
	OrganizationCode string `json:"organizationCode"`
	OrganizationName string `json:"organizationName"`
	PunchIn          bool   `json:"punchIn"`
	Terminated       bool   `json:"terminated"`
}

type TimecardReference struct {
	ReferenceKey       string              `json:"referenceKey"`
	CanEdit            bool                `json:"canEdit"`
	StartDate          string              `json:"startDate"`
	EndDate            string              `json:"endDate"`
	CutOff             string              `json:"cutOff"`
	Year               int64               `json:"year"`
	Month              int64               `json:"month"`
	LastProcessing     string              `json:"lastProcessing"`
	Info               string              `json:"info"`
	ProcessingStatus   string              `json:"processingStatus"`
	Previous           string              `json:"previous"`
	Next               interface{}         `json:"next"`
	Validation         Validation          `json:"validation"`
	DaysWithoutPunch   []interface{}       `json:"daysWithoutPunch"`
	PreviousReferences []PreviousReference `json:"previousReferences"`
	OutdatedTimesheet  bool                `json:"outdatedTimesheet"`
}

type PreviousReference struct {
	Year         int64  `json:"year"`
	Month        int64  `json:"month"`
	ReferenceKey string `json:"referenceKey"`
}

type Validation struct {
	OtherReferencePending bool `json:"otherReferencePending"`
}

type TimetableEntry struct {
	Date               string          `json:"date"`
	TimeOffDescription *string         `json:"timeOffDescription"`
	ExceptionDay       bool            `json:"exceptionDay"`
	Inconsistent       bool            `json:"inconsistent"`
	Changed            bool            `json:"changed"`
	LastChange         string          `json:"lastChange"`
	TimesheetKey       string          `json:"timesheetKey"`
	Timeline           []TimelineEntry `json:"timeline,omitempty"`
	Summary            []SummaryEntry  `json:"summary,omitempty"`
	Status             []Status        `json:"status"`
	Task               *Task           `json:"task,omitempty"`
	Variables          *Variables      `json:"variables,omitempty"`
}

type Status struct {
	Code  string `json:"code"`
	Value string `json:"value"`
	Meta  []Meta `json:"meta"`
}

type Meta struct {
	Type   string `json:"type"`
	Format string `json:"format"`
}

type SummaryEntry struct {
	Description *string `json:"description"`
	List        []List  `json:"list"`
}

type List struct {
	SummaryCode  string                `json:"summaryCode"`
	SummaryName  string                `json:"summaryName"`
	Minutes      int64                 `json:"minutes"`
	SummaryGroup []List                `json:"summaryGroup,omitempty"`
	SummaryItem  []ClassificationEntry `json:"summaryItem,omitempty"`
}

type ClassificationEntry struct {
	Code           string `json:"code"`
	Description    string `json:"description"`
	Minutes        int64  `json:"minutes"`
	Classification string `json:"classification"`
	Inconsistent   bool   `json:"inconsistent"`
}

type Task struct {
	ID            string `json:"_id"`
	DefinitionKey string `json:"definitionKey"`
}

type TimelineEntry struct {
	DateTime    string `json:"dateTime"`
	ItemType    int64  `json:"itemType"`
	TimeColor   int64  `json:"timeColor"`
	LineColor   int64  `json:"lineColor"`
	InsideScale bool   `json:"insideScale"`
}

type Variables struct {
	LastSync                    *string                      `json:"lastSync,omitempty"`
	InternalStatus              string                       `json:"internalStatus"`
	LastSuggestionStatus        interface{}                  `json:"lastSuggestionStatus"`
	TimesheetAdjustment         *TimesheetAdjustment         `json:"timesheetAdjustment"`
	ChangeType                  *string                      `json:"changeType"`
	Receipt                     *Receipt                     `json:"receipt"`
	InconsistencyOwnerAoid      string                       `json:"inconsistencyOwnerAoid"`
	InconsistencyOwnerOoid      string                       `json:"inconsistencyOwnerOoid"`
	Inconsistency               Inconsistency                `json:"inconsistency"`
	WorkDayID                   string                       `json:"workDayId"`
	CompetenceClosed            *bool                        `json:"competenceClosed,omitempty"`
	SuggestionInteractionStatus *SuggestionInteractionStatus `json:"suggestionInteractionStatus,omitempty"`
}

type Inconsistency struct {
	ClassificationList []ClassificationEntry  `json:"classificationList"`
	Actions            Actions                `json:"actions"`
	Justification      []ClassificationEntry  `json:"justification"`
	Result             InconsistencyResult    `json:"result"`
	Employee           InconsistencyEmployee  `json:"employee"`
	Reference          InconsistencyReference `json:"reference"`
	Status             []Status               `json:"status"`
	Hash               []Hash                 `json:"hash"`
	Summary            []SummaryEntry         `json:"summary"`
	Timeline           []TimelineEntry        `json:"timeline"`
	LastCalcEfet       string                 `json:"lastCalcEfet"`
	LastChange         string                 `json:"lastChange"`
	Changed            bool                   `json:"changed"`
	Inconsistent       bool                   `json:"inconsistent"`
	ExceptionDay       bool                   `json:"exceptionDay"`
	Date               string                 `json:"date"`
	TimesheetKey       string                 `json:"timesheetKey"`
}

type Actions struct {
	Result        []ResultElement `json:"result"`
	Label         string          `json:"label"`
	HasSuggestion bool            `json:"hasSuggestion"`
	CanDismiss    bool            `json:"canDismiss"`
}

type ResultElement struct {
	Items []ItemItem `json:"items"`
	Value int64      `json:"value"`
	Type  string     `json:"type"`
	Label string     `json:"label"`
}

type ItemItem struct {
	Items          []ItemItem `json:"items,omitempty"`
	Description    *string    `json:"description,omitempty"`
	Value          int64      `json:"value"`
	Type           string     `json:"type"`
	Observation    *string    `json:"observation,omitempty"`
	Code           *string    `json:"code,omitempty"`
	Suggestion     *string    `json:"suggestion,omitempty"`
	Inconsistent   *bool      `json:"inconsistent,omitempty"`
	Classification *string    `json:"classification,omitempty"`
}

type InconsistencyEmployee struct {
	Avatar       string `json:"avatar"`
	EmployeeCode int64  `json:"employeeCode"`
	Name         string `json:"name"`
	ID           string `json:"id"`
}

type Hash struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type InconsistencyReference struct {
	Employee       TimecardEmployee `json:"employee"`
	LastProcessing string           `json:"lastProcessing"`
	Ano            int64            `json:"ano"`
	Mes            int64            `json:"mes"`
	Tipo           int64            `json:"tipo"`
	End            string           `json:"end"`
	Start          string           `json:"start"`
	ReferenceKey   string           `json:"referenceKey"`
}

type InconsistencyResult struct {
	Items []PurpleItem `json:"items"`
}

type PurpleItem struct {
	Value int64  `json:"value"`
	Type  string `json:"type"`
}

type Receipt struct {
	Delete []Add  `json:"delete,omitempty"`
	Hash   string `json:"hash"`
	Saved  bool   `json:"saved"`
	Add    []Add  `json:"add,omitempty"`
}

type Add struct {
	MovementType  MovementType `json:"movementType"`
	Justification string       `json:"justification"`
	PunchType     string       `json:"punchType"`
	ReferenceDate string       `json:"referenceDate"`
	PunchInTime   string       `json:"punchInTime"`
}

type MovementType struct {
	Default     bool   `json:"default"`
	Description string `json:"description"`
	Code        int64  `json:"code"`
}

type SuggestionInteractionStatus struct {
	InternalMessage InternalMessage `json:"internalMessage"`
}

type InternalMessage struct {
	Timestamp string `json:"timestamp"`
	MessageID string `json:"messageId"`
}

type TimesheetAdjustment struct {
	Timesheet []Timesheet `json:"timesheet"`
}

type Timesheet struct {
	MovementType     int64   `json:"movementType"`
	PunchType        string  `json:"punchType"`
	Justification    string  `json:"justification"`
	OldPunchInTime   *string `json:"oldPunchInTime,omitempty"`
	OldReferenceDate *string `json:"oldReferenceDate,omitempty"`
	ReferenceDate    *string `json:"referenceDate,omitempty"`
	PunchInTime      *string `json:"punchInTime,omitempty"`
}

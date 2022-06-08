package models

type Context struct {
	Contexts []ContextElement `json:"contexts"`
}

type ContextElement struct {
	ID             string      `json:"_id"`
	Orgoid         string      `json:"orgoid"`
	Associateoid   string      `json:"associateoid"`
	Autoloaded     bool        `json:"autoloaded"`
	Attributes     interface{} `json:"attributes"`
	CreationDate   string      `json:"creationDate"`
	ValidationDate string      `json:"validationDate"`
	ContextID      string      `json:"id"`
	Details        Details     `json:"details"`
	FullyFetched   bool        `json:"fullyFetched"`
	Default        bool        `json:"default"`
	UIType         string      `json:"uiType"`
}

type Details struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Value       Value     `json:"value"`
	DependentBy []Details `json:"dependentBy,omitempty"`
}

type Value struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

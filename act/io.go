package act

type Hook struct {
	Name     string         `json:"name"`
	Priority uint           `json:"priority"`
	Params   map[string]any `json:"params"`
	Result   map[string]any `json:"result"`
}

type Input struct {
	Name   string         `json:"name"`
	Policy map[string]any `json:"policy"`
	Signal map[string]any `json:"signal"`
	Hook   Hook           `json:"hook"`
	Sync   bool           `json:"sync"`
}

type Output struct {
	MatchedPolicy string         `json:"matchedPolicy"`
	Metadata      map[string]any `json:"metadata"`
	Verdict       any            `json:"verdict"`
	Terminal      bool           `json:"terminal"`
	Sync          bool           `json:"sync"`
}

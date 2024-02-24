package act

type Input struct {
	Name   string         `json:"name"`
	Policy map[string]any `json:"policy"`
	Signal map[string]any `json:"signal"`
	Sync   bool           `json:"sync"`
}

type Output struct {
	MatchedPolicy string         `json:"matchedPolicy"`
	Metadata      map[string]any `json:"metadata"`
	Verdict       bool           `json:"verdict"`
	Terminal      bool           `json:"terminal"`
	Sync          bool           `json:"sync"`
}

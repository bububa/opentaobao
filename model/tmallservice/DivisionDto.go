package tmallservice

// DivisionDto 
type DivisionDto struct {
    // 1
    DivisionNames   []string `json:"division_names,omitempty" xml:"division_names>string,omitempty"`
}

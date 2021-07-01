package tmallservice

// DivisionDto 结构体
type DivisionDto struct {
	// 1
	DivisionNames []string `json:"division_names,omitempty" xml:"division_names>string,omitempty"`
}

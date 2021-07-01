package tmic

// OptionItemBo 结构体
type OptionItemBo struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// tip
	Tip string `json:"tip,omitempty" xml:"tip,omitempty"`
	// supplement
	Supplement bool `json:"supplement,omitempty" xml:"supplement,omitempty"`
	// exclusion
	Exclusion bool `json:"exclusion,omitempty" xml:"exclusion,omitempty"`
	// randomGroupNumber
	RandomGroupNumber int64 `json:"random_group_number,omitempty" xml:"random_group_number,omitempty"`
	// end
	End bool `json:"end,omitempty" xml:"end,omitempty"`
}

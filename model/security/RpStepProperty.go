package security

// RpStepProperty 结构体
type RpStepProperty struct {
	// property
	Property *RpProperty `json:"property,omitempty" xml:"property,omitempty"`
	// optional
	Optional bool `json:"optional,omitempty" xml:"optional,omitempty"`
	// 是否可选
	IsOptional bool `json:"is_optional,omitempty" xml:"is_optional,omitempty"`
}

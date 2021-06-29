package security

// RpStepProperty 
type RpStepProperty struct {
    // optional
    Optional   bool `json:"optional,omitempty" xml:"optional,omitempty"`
    // property
    Property   *RpProperty `json:"property,omitempty" xml:"property,omitempty"`
    // 是否可选
    IsOptional   bool `json:"is_optional,omitempty" xml:"is_optional,omitempty"`
}

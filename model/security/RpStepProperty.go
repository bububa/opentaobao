package security

// RpStepProperty 
type RpStepProperty struct {

    // optional
    Optional   bool `json:"optional,omitempty"`

    // property
    Property   *RpProperty `json:"property,omitempty"`

    // 是否可选
    IsOptional   bool `json:"is_optional,omitempty"`

}

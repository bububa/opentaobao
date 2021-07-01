package security

// RpStepItem 结构体
type RpStepItem struct {
	// jsonAssist
	JsonAssist string `json:"json_assist,omitempty" xml:"json_assist,omitempty"`
	// properties
	Properties []RpStepProperty `json:"properties,omitempty" xml:"properties>rp_step_property,omitempty"`
	// stepType
	StepType *RpStepType `json:"step_type,omitempty" xml:"step_type,omitempty"`
}

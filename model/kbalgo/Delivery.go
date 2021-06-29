package kbalgo

// Delivery 
type Delivery struct {
    // 分
    Min   string `json:"min,omitempty" xml:"min,omitempty"`
    // step_min
    StepMin   string `json:"step_min,omitempty" xml:"step_min,omitempty"`
    // step_base
    StepBase   string `json:"step_base,omitempty" xml:"step_base,omitempty"`
}

package security

// RpStepItem 
type RpStepItem struct {

    // jsonAssist
    
    JsonAssist   string `json:"json_assist,omitempty" xml:"json_assist,omitempty"`
    

    // properties
    
    Properties   []RpStepProperty `json:"properties,omitempty" xml:"properties,omitempty"`
    

    // stepType
    
    StepType   *RpStepType `json:"step_type,omitempty" xml:"step_type,omitempty"`
    

}

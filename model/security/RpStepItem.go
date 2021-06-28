package security

// RpStepItem 
/* model for simplify = false
type RpStepItem struct {

    // jsonAssist
    
    JsonAssist   string `json:"json_assist,omitempty"`
    

    // properties
    
    Properties  struct {
        RpStepProperty  []RpStepProperty `json:"rp_step_property,omitempty"`
    } `json:"properties,omitempty"`
    

    // stepType
    
    StepType  *struct {
        RpStepType  *RpStepType `json:"rp_step_type,omitempty"`
    } `json:"step_type,omitempty"`
    

}
*/

// RpStepItem 
type RpStepItem struct {

    // jsonAssist
    JsonAssist   string `json:"json_assist,omitempty"`

    // properties
    Properties   []RpStepProperty `json:"properties,omitempty"`

    // stepType
    StepType   *RpStepType `json:"step_type,omitempty"`

}

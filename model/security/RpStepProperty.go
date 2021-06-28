package security

// RpStepProperty 
/* model for simplify = false
type RpStepProperty struct {

    // optional
    
    Optional   bool `json:"optional,omitempty"`
    

    // property
    
    Property  *struct {
        RpProperty  *RpProperty `json:"rp_property,omitempty"`
    } `json:"property,omitempty"`
    

    // 是否可选
    
    IsOptional   bool `json:"is_optional,omitempty"`
    

}
*/

// RpStepProperty 
type RpStepProperty struct {

    // optional
    Optional   bool `json:"optional,omitempty"`

    // property
    Property   *RpProperty `json:"property,omitempty"`

    // 是否可选
    IsOptional   bool `json:"is_optional,omitempty"`

}

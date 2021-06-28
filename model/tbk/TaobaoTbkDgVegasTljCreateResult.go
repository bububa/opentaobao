package tbk

// TaobaoTbkDgVegasTljCreateResult 
/* model for simplify = false
type TaobaoTbkDgVegasTljCreateResult struct {

    // model
    
    Model  *struct {
        RightsInstanceCreateResult  *RightsInstanceCreateResult `json:"rights_instance_create_result,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoTbkDgVegasTljCreateResult 
type TaobaoTbkDgVegasTljCreateResult struct {

    // model
    Model   *RightsInstanceCreateResult `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}

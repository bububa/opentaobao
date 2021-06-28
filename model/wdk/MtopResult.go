package wdk

// MtopResult 
/* model for simplify = false
type MtopResult struct {

    // model
    
    Model  *struct {
        MemberInfoDto  *MemberInfoDto `json:"member_info_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// MtopResult 
type MtopResult struct {

    // model
    Model   *MemberInfoDto `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}

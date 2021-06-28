package util

// TaobaoOpenlinkSessionGetResult 
/* model for simplify = false
type TaobaoOpenlinkSessionGetResult struct {

    // model
    
    Model  *struct {
        SessionInfo  *SessionInfo `json:"session_info,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoOpenlinkSessionGetResult 
type TaobaoOpenlinkSessionGetResult struct {

    // model
    Model   *SessionInfo `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}

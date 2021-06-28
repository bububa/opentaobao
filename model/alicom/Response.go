package alicom

// Response 
/* model for simplify = false
type Response struct {

    // 转呼控制msg
    
    Message   string `json:"message,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // module
    
    ControlRespDto  *struct {
        ControlRespDto  *ControlRespDto `json:"control_resp_dto,omitempty"`
    } `json:"control_resp_dto,omitempty"`
    

    // module
    
    Module   bool `json:"module,omitempty"`
    

}
*/

// Response 
type Response struct {

    // 转呼控制msg
    Message   string `json:"message,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // module
    ControlRespDto   *ControlRespDto `json:"control_resp_dto,omitempty"`

    // module
    Module   bool `json:"module,omitempty"`

}

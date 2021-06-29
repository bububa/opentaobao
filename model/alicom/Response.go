package alicom

// Response 
type Response struct {
    // 转呼控制msg
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // module
    ControlRespDto   *ControlRespDTO `json:"control_resp_dto,omitempty" xml:"control_resp_dto,omitempty"`
    // module
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
}

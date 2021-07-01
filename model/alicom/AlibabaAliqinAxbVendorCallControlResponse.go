package alicom

// AlibabaAliqinAxbVendorCallControlResponse 
type AlibabaAliqinAxbVendorCallControlResponse struct {
    // 转呼控制msg
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // module
    ControlRespDto   *ControlRespDto `json:"control_resp_dto,omitempty" xml:"control_resp_dto,omitempty"`
}

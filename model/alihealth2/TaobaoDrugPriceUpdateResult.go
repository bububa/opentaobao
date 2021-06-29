package alihealth2

// TaobaoDrugPriceUpdateResult 
type TaobaoDrugPriceUpdateResult struct {
    // 异常代码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 异常信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 请求是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}

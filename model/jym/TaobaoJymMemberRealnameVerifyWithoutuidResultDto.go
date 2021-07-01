package jym

// TaobaoJymMemberRealnameVerifyWithoutuidResultDto 
type TaobaoJymMemberRealnameVerifyWithoutuidResultDto struct {
    // 1
    Result   *RealNameVerifyTopDto `json:"result,omitempty" xml:"result,omitempty"`
    // 调用是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 调用接口结果编码
    StateCode   string `json:"state_code,omitempty" xml:"state_code,omitempty"`
    // 调用接口异常信息明细
    ExtraErrMsg   string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
}

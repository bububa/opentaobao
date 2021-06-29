package alihealth2

// TopResultModel 
type TopResultModel struct {
    // model
    Model   *DrugInfoDTO `json:"model,omitempty" xml:"model,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回json字符串，包含商家的省、市、县、和id
    Models   []string `json:"models,omitempty" xml:"models>string,omitempty"`
}

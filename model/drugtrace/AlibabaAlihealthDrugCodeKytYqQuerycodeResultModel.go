package drugtrace

// AlibabaalihealthdrugcodekytyqquerycodeResultModel 结构体
type AlibabaalihealthdrugcodekytyqquerycodeResultModel struct {
	// 内层大对象
	Models []CodeDrugInfoDto `json:"models,omitempty" xml:"models>code_drug_info_dto,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口调用标识：true/false
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

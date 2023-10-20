package alihouse

// AlibabaAlihouseStorePunishInfoQueryResult 结构体
type AlibabaAlihouseStorePunishInfoQueryResult struct {
	// dto
	Data []StorePunishDto `json:"data,omitempty" xml:"data>store_punish_dto,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

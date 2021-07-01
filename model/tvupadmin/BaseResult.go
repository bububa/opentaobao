package tvupadmin

// BaseResult 结构体
type BaseResult struct {
	// 返回码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// retValue
	RetValues []SimpleBotInfo `json:"ret_values,omitempty" xml:"ret_values>simple_bot_info,omitempty"`
	// 返回数据
	RetValue string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

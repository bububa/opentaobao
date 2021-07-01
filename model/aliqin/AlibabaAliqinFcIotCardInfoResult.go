package aliqin

// AlibabaAliqinFcIotCardInfoResult 结构体
type AlibabaAliqinFcIotCardInfoResult struct {
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// "OpenTime":"开户时间","IMSI":"IMSI号","FirstActiveTime":"第一次激活时间","MSISDN":"MSISDN号","ICCID":"ICCID号"
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

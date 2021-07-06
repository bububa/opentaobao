package aliqin

// AlibabaAliqinFcIotCardStatusResult 结构体
type AlibabaAliqinFcIotCardStatusResult struct {
	// "MsisdnStatus":"卡状态","MSISDN":"MSISDN号","ICCID":"ICCID号"
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 1.成功；2.失败
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

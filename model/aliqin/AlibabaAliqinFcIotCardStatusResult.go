package aliqin

// AlibabaAliqinFcIotCardStatusResult 结构体
type AlibabaAliqinFcIotCardStatusResult struct {
	// &#34;MsisdnStatus&#34;:&#34;卡状态&#34;,&#34;MSISDN&#34;:&#34;MSISDN号&#34;,&#34;ICCID&#34;:&#34;ICCID号&#34;
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 1.成功；2.失败
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

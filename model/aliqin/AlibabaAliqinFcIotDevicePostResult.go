package aliqin

// AlibabaAliqinFcIotDevicePostResult 结构体
type AlibabaAliqinFcIotDevicePostResult struct {
	// 响应结果描述
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 响应code 判断以此判断是否提交成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

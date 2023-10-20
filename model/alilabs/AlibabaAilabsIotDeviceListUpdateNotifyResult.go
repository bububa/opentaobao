package alilabs

// AlibabaailabsiotdevicelistupdatenotifyResult 结构体
type AlibabaailabsiotdevicelistupdatenotifyResult struct {
	// 附加信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求响应码，200代表成功
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 设备列表更新通知是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

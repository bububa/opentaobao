package tmallgenie

// TaobaoailabaicloudtopmemonotelistResult 结构体
type TaobaoailabaicloudtopmemonotelistResult struct {
	// 服务的实际返回结果
	Notes []Note `json:"notes,omitempty" xml:"notes>note,omitempty"`
	// 附加信息，典型应用场景是对失败调用进行简述，方便调用方定位问题
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

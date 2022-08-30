package wdk

// UmsResult 结构体
type UmsResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

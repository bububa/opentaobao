package interact

// CommonResult 结构体
type CommonResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 创建的互动实例
	Data *InteractDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否调用成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

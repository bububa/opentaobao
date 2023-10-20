package idle

// AlibabaidlerentitemqueryResult 结构体
type AlibabaidlerentitemqueryResult struct {
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回素材id
	Data *AlibabaidlerentitemqueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 数据是否可用
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

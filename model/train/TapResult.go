package train

// TapResult 结构体
type TapResult struct {
	// 失败msg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 失败code
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回信息
	Module *FreeChildrenTicketDetailRs `json:"module,omitempty" xml:"module,omitempty"`
	// 处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

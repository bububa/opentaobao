package pentraprism

// TaskResult 结构体
type TaskResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误详细信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 做任务时间
	Now string `json:"now,omitempty" xml:"now,omitempty"`
	// 任务完成个数
	FinishCount int64 `json:"finish_count,omitempty" xml:"finish_count,omitempty"`
	// 任务数据返回对象
	Model *OpenTaskInfoVo `json:"model,omitempty" xml:"model,omitempty"`
	// 匹配命中任务数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

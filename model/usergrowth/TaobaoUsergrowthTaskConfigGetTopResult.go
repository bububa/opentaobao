package usergrowth

// TaobaoUsergrowthTaskConfigGetTopResult 结构体
type TaobaoUsergrowthTaskConfigGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务配置
	Data *TaskConfig `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

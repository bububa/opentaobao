package flight

// AlitrippolicyspecialcompressionuploadResultDto 结构体
type AlitrippolicyspecialcompressionuploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

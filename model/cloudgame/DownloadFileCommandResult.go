package cloudgame

// DownloadFileCommandResult 结构体
type DownloadFileCommandResult struct {
	// 文件下载任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 文件下载失败原因
	FailureReason string `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// 文件下载是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

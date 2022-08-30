package cloudgame

// FileDownloadCallbackRequest 结构体
type FileDownloadCallbackRequest struct {
	// 文件下载结果
	Result *DownloadFileCommandResult `json:"result,omitempty" xml:"result,omitempty"`
}

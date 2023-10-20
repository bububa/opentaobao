package cloudgame

import (
	"sync"
)

// FileDownloadCallbackRequest 结构体
type FileDownloadCallbackRequest struct {
	// 文件下载结果
	Result *DownloadFileCommandResult `json:"result,omitempty" xml:"result,omitempty"`
}

var poolFileDownloadCallbackRequest = sync.Pool{
	New: func() any {
		return new(FileDownloadCallbackRequest)
	},
}

// GetFileDownloadCallbackRequest() 从对象池中获取FileDownloadCallbackRequest
func GetFileDownloadCallbackRequest() *FileDownloadCallbackRequest {
	return poolFileDownloadCallbackRequest.Get().(*FileDownloadCallbackRequest)
}

// ReleaseFileDownloadCallbackRequest 释放FileDownloadCallbackRequest
func ReleaseFileDownloadCallbackRequest(v *FileDownloadCallbackRequest) {
	v.Result = nil
	poolFileDownloadCallbackRequest.Put(v)
}

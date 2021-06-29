package util

// TopDownloadRecordDO 
type TopDownloadRecordDO struct {
    // 下载链接
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 文件创建时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 下载链接状态。1:未下载。2:已下载
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}

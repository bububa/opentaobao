package util

import (
	"sync"
)

// TopDownloadRecordDo 结构体
type TopDownloadRecordDo struct {
	// 下载链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 文件创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 下载链接状态。1:未下载。2:已下载
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolTopDownloadRecordDo = sync.Pool{
	New: func() any {
		return new(TopDownloadRecordDo)
	},
}

// GetTopDownloadRecordDo() 从对象池中获取TopDownloadRecordDo
func GetTopDownloadRecordDo() *TopDownloadRecordDo {
	return poolTopDownloadRecordDo.Get().(*TopDownloadRecordDo)
}

// ReleaseTopDownloadRecordDo 释放TopDownloadRecordDo
func ReleaseTopDownloadRecordDo(v *TopDownloadRecordDo) {
	v.Url = ""
	v.Created = ""
	v.Status = 0
	poolTopDownloadRecordDo.Put(v)
}

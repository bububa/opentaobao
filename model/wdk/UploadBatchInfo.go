package wdk

import (
	"sync"
)

// UploadBatchInfo 结构体
type UploadBatchInfo struct {
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 该批次数据的时间范围-起始时间
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 该批次数据的时间范围-结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 批次总量
	BatchCount int64 `json:"batch_count,omitempty" xml:"batch_count,omitempty"`
}

var poolUploadBatchInfo = sync.Pool{
	New: func() any {
		return new(UploadBatchInfo)
	},
}

// GetUploadBatchInfo() 从对象池中获取UploadBatchInfo
func GetUploadBatchInfo() *UploadBatchInfo {
	return poolUploadBatchInfo.Get().(*UploadBatchInfo)
}

// ReleaseUploadBatchInfo 释放UploadBatchInfo
func ReleaseUploadBatchInfo(v *UploadBatchInfo) {
	v.BatchNo = ""
	v.BeginTime = ""
	v.EndTime = ""
	v.BatchCount = 0
	poolUploadBatchInfo.Put(v)
}

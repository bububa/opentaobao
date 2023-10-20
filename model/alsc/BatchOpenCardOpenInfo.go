package alsc

import (
	"sync"
)

// BatchOpenCardOpenInfo 结构体
type BatchOpenCardOpenInfo struct {
	// 结果 &lt; KEY：Id  VALUE：描述（SUCCESS-通过） &gt;
	ResultMap string `json:"result_map,omitempty" xml:"result_map,omitempty"`
	// 是否全部开通成功
	AllSuccess string `json:"all_success,omitempty" xml:"all_success,omitempty"`
}

var poolBatchOpenCardOpenInfo = sync.Pool{
	New: func() any {
		return new(BatchOpenCardOpenInfo)
	},
}

// GetBatchOpenCardOpenInfo() 从对象池中获取BatchOpenCardOpenInfo
func GetBatchOpenCardOpenInfo() *BatchOpenCardOpenInfo {
	return poolBatchOpenCardOpenInfo.Get().(*BatchOpenCardOpenInfo)
}

// ReleaseBatchOpenCardOpenInfo 释放BatchOpenCardOpenInfo
func ReleaseBatchOpenCardOpenInfo(v *BatchOpenCardOpenInfo) {
	v.ResultMap = ""
	v.AllSuccess = ""
	poolBatchOpenCardOpenInfo.Put(v)
}

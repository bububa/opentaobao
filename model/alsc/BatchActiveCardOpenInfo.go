package alsc

import (
	"sync"
)

// BatchActiveCardOpenInfo 结构体
type BatchActiveCardOpenInfo struct {
	// 结果 &lt; KEY：Id  VALUE：描述（SUCCESS-通过） &gt;
	ResultMap string `json:"result_map,omitempty" xml:"result_map,omitempty"`
	// 是否全部激活成功
	AllSuccess bool `json:"all_success,omitempty" xml:"all_success,omitempty"`
}

var poolBatchActiveCardOpenInfo = sync.Pool{
	New: func() any {
		return new(BatchActiveCardOpenInfo)
	},
}

// GetBatchActiveCardOpenInfo() 从对象池中获取BatchActiveCardOpenInfo
func GetBatchActiveCardOpenInfo() *BatchActiveCardOpenInfo {
	return poolBatchActiveCardOpenInfo.Get().(*BatchActiveCardOpenInfo)
}

// ReleaseBatchActiveCardOpenInfo 释放BatchActiveCardOpenInfo
func ReleaseBatchActiveCardOpenInfo(v *BatchActiveCardOpenInfo) {
	v.ResultMap = ""
	v.AllSuccess = false
	poolBatchActiveCardOpenInfo.Put(v)
}

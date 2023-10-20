package tmallservice

import (
	"sync"
)

// AlibabaSscServicecenterServicestoreQueryResult 结构体
type AlibabaSscServicecenterServicestoreQueryResult struct {
	// 明细条目执行结果对象
	ResultDataList []ResultData `json:"result_data_list,omitempty" xml:"result_data_list>result_data,omitempty"`
}

var poolAlibabaSscServicecenterServicestoreQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscServicecenterServicestoreQueryResult)
	},
}

// GetAlibabaSscServicecenterServicestoreQueryResult() 从对象池中获取AlibabaSscServicecenterServicestoreQueryResult
func GetAlibabaSscServicecenterServicestoreQueryResult() *AlibabaSscServicecenterServicestoreQueryResult {
	return poolAlibabaSscServicecenterServicestoreQueryResult.Get().(*AlibabaSscServicecenterServicestoreQueryResult)
}

// ReleaseAlibabaSscServicecenterServicestoreQueryResult 释放AlibabaSscServicecenterServicestoreQueryResult
func ReleaseAlibabaSscServicecenterServicestoreQueryResult(v *AlibabaSscServicecenterServicestoreQueryResult) {
	v.ResultDataList = v.ResultDataList[:0]
	poolAlibabaSscServicecenterServicestoreQueryResult.Put(v)
}

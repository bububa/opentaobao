package traveltrade

import (
	"sync"
)

// TopTripBookInfoResult 结构体
type TopTripBookInfoResult struct {
	// 交易预定详情
	Module *AlitripTravelBookinfoQueryModule `json:"module,omitempty" xml:"module,omitempty"`
}

var poolTopTripBookInfoResult = sync.Pool{
	New: func() any {
		return new(TopTripBookInfoResult)
	},
}

// GetTopTripBookInfoResult() 从对象池中获取TopTripBookInfoResult
func GetTopTripBookInfoResult() *TopTripBookInfoResult {
	return poolTopTripBookInfoResult.Get().(*TopTripBookInfoResult)
}

// ReleaseTopTripBookInfoResult 释放TopTripBookInfoResult
func ReleaseTopTripBookInfoResult(v *TopTripBookInfoResult) {
	v.Module = nil
	poolTopTripBookInfoResult.Put(v)
}

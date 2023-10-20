package hotelalliance

import (
	"sync"
)

// AllianceInfoResult 结构体
type AllianceInfoResult struct {
	// 菲住hid列表
	AllianceHids []int64 `json:"alliance_hids,omitempty" xml:"alliance_hids>int64,omitempty"`
}

var poolAllianceInfoResult = sync.Pool{
	New: func() any {
		return new(AllianceInfoResult)
	},
}

// GetAllianceInfoResult() 从对象池中获取AllianceInfoResult
func GetAllianceInfoResult() *AllianceInfoResult {
	return poolAllianceInfoResult.Get().(*AllianceInfoResult)
}

// ReleaseAllianceInfoResult 释放AllianceInfoResult
func ReleaseAllianceInfoResult(v *AllianceInfoResult) {
	v.AllianceHids = v.AllianceHids[:0]
	poolAllianceInfoResult.Put(v)
}

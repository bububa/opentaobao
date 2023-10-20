package hotelalliance

import (
	"sync"
)

// HmsTopResultSet 结构体
type HmsTopResultSet struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回module
	Result *AllianceInfoResult `json:"result,omitempty" xml:"result,omitempty"`
	// 操作是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolHmsTopResultSet = sync.Pool{
	New: func() any {
		return new(HmsTopResultSet)
	},
}

// GetHmsTopResultSet() 从对象池中获取HmsTopResultSet
func GetHmsTopResultSet() *HmsTopResultSet {
	return poolHmsTopResultSet.Get().(*HmsTopResultSet)
}

// ReleaseHmsTopResultSet 释放HmsTopResultSet
func ReleaseHmsTopResultSet(v *HmsTopResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	v.Success = false
	poolHmsTopResultSet.Put(v)
}

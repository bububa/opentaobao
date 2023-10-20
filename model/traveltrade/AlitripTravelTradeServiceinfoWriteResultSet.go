package traveltrade

import (
	"sync"
)

// AlitripTravelTradeServiceinfoWriteResultSet 结构体
type AlitripTravelTradeServiceinfoWriteResultSet struct {
	// 错误码信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误详细信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 订单标注服务信息是否成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolAlitripTravelTradeServiceinfoWriteResultSet = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradeServiceinfoWriteResultSet)
	},
}

// GetAlitripTravelTradeServiceinfoWriteResultSet() 从对象池中获取AlitripTravelTradeServiceinfoWriteResultSet
func GetAlitripTravelTradeServiceinfoWriteResultSet() *AlitripTravelTradeServiceinfoWriteResultSet {
	return poolAlitripTravelTradeServiceinfoWriteResultSet.Get().(*AlitripTravelTradeServiceinfoWriteResultSet)
}

// ReleaseAlitripTravelTradeServiceinfoWriteResultSet 释放AlitripTravelTradeServiceinfoWriteResultSet
func ReleaseAlitripTravelTradeServiceinfoWriteResultSet(v *AlitripTravelTradeServiceinfoWriteResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.IsSuccess = false
	v.Error = false
	poolAlitripTravelTradeServiceinfoWriteResultSet.Put(v)
}

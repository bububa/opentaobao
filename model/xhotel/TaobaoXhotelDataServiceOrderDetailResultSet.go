package xhotel

import (
	"sync"
)

// TaobaoXhotelDataServiceOrderDetailResultSet 结构体
type TaobaoXhotelDataServiceOrderDetailResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// firstResult
	FirstResult *TopAdsTripSvcQueryResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

var poolTaobaoXhotelDataServiceOrderDetailResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDataServiceOrderDetailResultSet)
	},
}

// GetTaobaoXhotelDataServiceOrderDetailResultSet() 从对象池中获取TaobaoXhotelDataServiceOrderDetailResultSet
func GetTaobaoXhotelDataServiceOrderDetailResultSet() *TaobaoXhotelDataServiceOrderDetailResultSet {
	return poolTaobaoXhotelDataServiceOrderDetailResultSet.Get().(*TaobaoXhotelDataServiceOrderDetailResultSet)
}

// ReleaseTaobaoXhotelDataServiceOrderDetailResultSet 释放TaobaoXhotelDataServiceOrderDetailResultSet
func ReleaseTaobaoXhotelDataServiceOrderDetailResultSet(v *TaobaoXhotelDataServiceOrderDetailResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FirstResult = nil
	poolTaobaoXhotelDataServiceOrderDetailResultSet.Put(v)
}

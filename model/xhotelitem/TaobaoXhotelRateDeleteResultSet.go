package xhotelitem

import (
	"sync"
)

// TaobaoXhotelRateDeleteResultSet 结构体
type TaobaoXhotelRateDeleteResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// rateid-房型id-房价id
	RateidGidRpid string `json:"rateid_gid_rpid,omitempty" xml:"rateid_gid_rpid,omitempty"`
}

var poolTaobaoXhotelRateDeleteResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateDeleteResultSet)
	},
}

// GetTaobaoXhotelRateDeleteResultSet() 从对象池中获取TaobaoXhotelRateDeleteResultSet
func GetTaobaoXhotelRateDeleteResultSet() *TaobaoXhotelRateDeleteResultSet {
	return poolTaobaoXhotelRateDeleteResultSet.Get().(*TaobaoXhotelRateDeleteResultSet)
}

// ReleaseTaobaoXhotelRateDeleteResultSet 释放TaobaoXhotelRateDeleteResultSet
func ReleaseTaobaoXhotelRateDeleteResultSet(v *TaobaoXhotelRateDeleteResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RateidGidRpid = ""
	poolTaobaoXhotelRateDeleteResultSet.Put(v)
}

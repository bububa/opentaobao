package xhotelonlineorder

import (
	"sync"
)

// TaobaoXhotelOrderFutureFacescanPutResultSet 结构体
type TaobaoXhotelOrderFutureFacescanPutResultSet struct {
	// 错误描述
	Errormsg string `json:"errormsg,omitempty" xml:"errormsg,omitempty"`
	// 错误码
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// 是否更新失败。返回true表示更新成功。否则请读取错误码与错误描述
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelOrderFutureFacescanPutResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderFutureFacescanPutResultSet)
	},
}

// GetTaobaoXhotelOrderFutureFacescanPutResultSet() 从对象池中获取TaobaoXhotelOrderFutureFacescanPutResultSet
func GetTaobaoXhotelOrderFutureFacescanPutResultSet() *TaobaoXhotelOrderFutureFacescanPutResultSet {
	return poolTaobaoXhotelOrderFutureFacescanPutResultSet.Get().(*TaobaoXhotelOrderFutureFacescanPutResultSet)
}

// ReleaseTaobaoXhotelOrderFutureFacescanPutResultSet 释放TaobaoXhotelOrderFutureFacescanPutResultSet
func ReleaseTaobaoXhotelOrderFutureFacescanPutResultSet(v *TaobaoXhotelOrderFutureFacescanPutResultSet) {
	v.Errormsg = ""
	v.Errorcode = ""
	v.Success = false
	poolTaobaoXhotelOrderFutureFacescanPutResultSet.Put(v)
}

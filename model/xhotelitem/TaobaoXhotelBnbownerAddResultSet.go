package xhotelitem

import (
	"sync"
)

// TaobaoXhotelBnbownerAddResultSet 结构体
type TaobaoXhotelBnbownerAddResultSet struct {
	// firstResult
	FirstResult *AddOwnerParam `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

var poolTaobaoXhotelBnbownerAddResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbownerAddResultSet)
	},
}

// GetTaobaoXhotelBnbownerAddResultSet() 从对象池中获取TaobaoXhotelBnbownerAddResultSet
func GetTaobaoXhotelBnbownerAddResultSet() *TaobaoXhotelBnbownerAddResultSet {
	return poolTaobaoXhotelBnbownerAddResultSet.Get().(*TaobaoXhotelBnbownerAddResultSet)
}

// ReleaseTaobaoXhotelBnbownerAddResultSet 释放TaobaoXhotelBnbownerAddResultSet
func ReleaseTaobaoXhotelBnbownerAddResultSet(v *TaobaoXhotelBnbownerAddResultSet) {
	v.FirstResult = nil
	poolTaobaoXhotelBnbownerAddResultSet.Put(v)
}

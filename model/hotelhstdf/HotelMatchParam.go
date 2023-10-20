package hotelhstdf

import (
	"sync"
)

// HotelMatchParam 结构体
type HotelMatchParam struct {
	// 卖家酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelMatchParam = sync.Pool{
	New: func() any {
		return new(HotelMatchParam)
	},
}

// GetHotelMatchParam() 从对象池中获取HotelMatchParam
func GetHotelMatchParam() *HotelMatchParam {
	return poolHotelMatchParam.Get().(*HotelMatchParam)
}

// ReleaseHotelMatchParam 释放HotelMatchParam
func ReleaseHotelMatchParam(v *HotelMatchParam) {
	v.Hid = 0
	v.Shid = 0
	poolHotelMatchParam.Put(v)
}

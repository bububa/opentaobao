package alitripmerchant

import (
	"sync"
)

// HotelDetailsGroupParam 结构体
type HotelDetailsGroupParam struct {
	// 组名
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 房型id集合
	RoomIds string `json:"room_ids,omitempty" xml:"room_ids,omitempty"`
	// 取消政策类型
	CancelType string `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 早餐
	Breakfast string `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
}

var poolHotelDetailsGroupParam = sync.Pool{
	New: func() any {
		return new(HotelDetailsGroupParam)
	},
}

// GetHotelDetailsGroupParam() 从对象池中获取HotelDetailsGroupParam
func GetHotelDetailsGroupParam() *HotelDetailsGroupParam {
	return poolHotelDetailsGroupParam.Get().(*HotelDetailsGroupParam)
}

// ReleaseHotelDetailsGroupParam 释放HotelDetailsGroupParam
func ReleaseHotelDetailsGroupParam(v *HotelDetailsGroupParam) {
	v.GroupName = ""
	v.RoomIds = ""
	v.CancelType = ""
	v.Breakfast = ""
	poolHotelDetailsGroupParam.Put(v)
}

package hotelhstdf

import (
	"sync"
)

// GetHotelRoomStaticParam 结构体
type GetHotelRoomStaticParam struct {
	// 第1页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页取100条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 字典类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolGetHotelRoomStaticParam = sync.Pool{
	New: func() any {
		return new(GetHotelRoomStaticParam)
	},
}

// GetGetHotelRoomStaticParam() 从对象池中获取GetHotelRoomStaticParam
func GetGetHotelRoomStaticParam() *GetHotelRoomStaticParam {
	return poolGetHotelRoomStaticParam.Get().(*GetHotelRoomStaticParam)
}

// ReleaseGetHotelRoomStaticParam 释放GetHotelRoomStaticParam
func ReleaseGetHotelRoomStaticParam(v *GetHotelRoomStaticParam) {
	v.Page = 0
	v.PageSize = 0
	v.Type = 0
	poolGetHotelRoomStaticParam.Put(v)
}

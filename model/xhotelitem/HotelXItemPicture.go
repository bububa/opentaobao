package xhotelitem

import (
	"sync"
)

// HotelXItemPicture 结构体
type HotelXItemPicture struct {
	// 图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 是否主图
	IsMain bool `json:"is_main,omitempty" xml:"is_main,omitempty"`
}

var poolHotelXItemPicture = sync.Pool{
	New: func() any {
		return new(HotelXItemPicture)
	},
}

// GetHotelXItemPicture() 从对象池中获取HotelXItemPicture
func GetHotelXItemPicture() *HotelXItemPicture {
	return poolHotelXItemPicture.Get().(*HotelXItemPicture)
}

// ReleaseHotelXItemPicture 释放HotelXItemPicture
func ReleaseHotelXItemPicture(v *HotelXItemPicture) {
	v.Url = ""
	v.IsMain = false
	poolHotelXItemPicture.Put(v)
}

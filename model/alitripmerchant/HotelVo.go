package alitripmerchant

import (
	"sync"
)

// HotelVo 结构体
type HotelVo struct {
	// 品牌名
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 星级
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 酒店详细描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 品牌编码
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 酒店扩展
	HotelFeature *HotelFeatureVo `json:"hotel_feature,omitempty" xml:"hotel_feature,omitempty"`
	// 飞猪旗舰店的ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
}

var poolHotelVo = sync.Pool{
	New: func() any {
		return new(HotelVo)
	},
}

// GetHotelVo() 从对象池中获取HotelVo
func GetHotelVo() *HotelVo {
	return poolHotelVo.Get().(*HotelVo)
}

// ReleaseHotelVo 释放HotelVo
func ReleaseHotelVo(v *HotelVo) {
	v.BrandName = ""
	v.Star = ""
	v.Phone = ""
	v.Description = ""
	v.HotelId = ""
	v.BrandCode = ""
	v.Shid = 0
	v.HotelFeature = nil
	v.Hid = 0
	poolHotelVo.Put(v)
}

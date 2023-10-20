package kbalgo

import (
	"sync"
)

// Poi 结构体
type Poi struct {
	// poi的标签
	Labels []Label `json:"labels,omitempty" xml:"labels>label,omitempty"`
	// POI地址(不包含城市，行政区，商圈)]]
	PoiAddress string `json:"poi_address,omitempty" xml:"poi_address,omitempty"`
	// poi名称
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// app_key
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 营业时间
	ShopBusiness string `json:"shop_business,omitempty" xml:"shop_business,omitempty"`
	// 固定的
	AppSchema string `json:"app_schema,omitempty" xml:"app_schema,omitempty"`
	// latitude
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// county
	County string `json:"county,omitempty" xml:"county,omitempty"`
	// L1DXZ0001M3
	PoiId string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	// longitude
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
}

var poolPoi = sync.Pool{
	New: func() any {
		return new(Poi)
	},
}

// GetPoi() 从对象池中获取Poi
func GetPoi() *Poi {
	return poolPoi.Get().(*Poi)
}

// ReleasePoi 释放Poi
func ReleasePoi(v *Poi) {
	v.Labels = v.Labels[:0]
	v.PoiAddress = ""
	v.PoiName = ""
	v.AppKey = ""
	v.Phone = ""
	v.City = ""
	v.ShopBusiness = ""
	v.AppSchema = ""
	v.Latitude = ""
	v.County = ""
	v.PoiId = ""
	v.Longitude = ""
	v.Province = ""
	poolPoi.Put(v)
}

package travel

import (
	"sync"
)

// PontusTravelItemHotel 结构体
type PontusTravelItemHotel struct {
	// 必填，所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 必填，酒店名称
	CnName string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
	// 必填，酒店等级
	HotelLevel string `json:"hotel_level,omitempty" xml:"hotel_level,omitempty"`
	// 必填，酒店房型
	HouseType string `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 酒店的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222
	Poi string `json:"poi,omitempty" xml:"poi,omitempty"`
	// POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
	PoiResource string `json:"poi_resource,omitempty" xml:"poi_resource,omitempty"`
}

var poolPontusTravelItemHotel = sync.Pool{
	New: func() any {
		return new(PontusTravelItemHotel)
	},
}

// GetPontusTravelItemHotel() 从对象池中获取PontusTravelItemHotel
func GetPontusTravelItemHotel() *PontusTravelItemHotel {
	return poolPontusTravelItemHotel.Get().(*PontusTravelItemHotel)
}

// ReleasePontusTravelItemHotel 释放PontusTravelItemHotel
func ReleasePontusTravelItemHotel(v *PontusTravelItemHotel) {
	v.City = ""
	v.CnName = ""
	v.HotelLevel = ""
	v.HouseType = ""
	v.Poi = ""
	v.PoiResource = ""
	poolPontusTravelItemHotel.Put(v)
}

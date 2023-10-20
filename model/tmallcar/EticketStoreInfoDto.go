package tmallcar

import (
	"sync"
)

// EticketStoreInfoDto 结构体
type EticketStoreInfoDto struct {
	// 街道名称
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 省名称
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区名称
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 经度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 联系方式
	StorePhone string `json:"store_phone,omitempty" xml:"store_phone,omitempty"`
	// 外部门店编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 纬度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 街道编码
	Town int64 `json:"town,omitempty" xml:"town,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 区域编码
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 省编码
	Prov int64 `json:"prov,omitempty" xml:"prov,omitempty"`
}

var poolEticketStoreInfoDto = sync.Pool{
	New: func() any {
		return new(EticketStoreInfoDto)
	},
}

// GetEticketStoreInfoDto() 从对象池中获取EticketStoreInfoDto
func GetEticketStoreInfoDto() *EticketStoreInfoDto {
	return poolEticketStoreInfoDto.Get().(*EticketStoreInfoDto)
}

// ReleaseEticketStoreInfoDto 释放EticketStoreInfoDto
func ReleaseEticketStoreInfoDto(v *EticketStoreInfoDto) {
	v.TownName = ""
	v.ProvName = ""
	v.Address = ""
	v.DistrictName = ""
	v.Latitude = ""
	v.CityName = ""
	v.StorePhone = ""
	v.OuterId = ""
	v.StoreName = ""
	v.Longitude = ""
	v.Town = 0
	v.OrderId = 0
	v.City = 0
	v.StoreId = 0
	v.District = 0
	v.Prov = 0
	poolEticketStoreInfoDto.Put(v)
}

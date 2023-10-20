package maitix

import (
	"sync"
)

// MoaAddressInfo 结构体
type MoaAddressInfo struct {
	// 国家ID，目前只支持中国，传1-如果是快递票必填
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 省ID，国标-如果是快递票必填
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 市ID，国标，快递票该字段必须填
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 区域ID，国标-如果是快递票必填
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolMoaAddressInfo = sync.Pool{
	New: func() any {
		return new(MoaAddressInfo)
	},
}

// GetMoaAddressInfo() 从对象池中获取MoaAddressInfo
func GetMoaAddressInfo() *MoaAddressInfo {
	return poolMoaAddressInfo.Get().(*MoaAddressInfo)
}

// ReleaseMoaAddressInfo 释放MoaAddressInfo
func ReleaseMoaAddressInfo(v *MoaAddressInfo) {
	v.CountryId = 0
	v.ProvinceId = 0
	v.CityId = 0
	v.AreaId = 0
	poolMoaAddressInfo.Put(v)
}

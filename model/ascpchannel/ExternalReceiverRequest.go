package ascpchannel

import (
	"sync"
)

// ExternalReceiverRequest 结构体
type ExternalReceiverRequest struct {
	// 收货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 街道对应名称
	StreetName string `json:"street_name,omitempty" xml:"street_name,omitempty"`
	// 收货人手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 邮编
	Post string `json:"post,omitempty" xml:"post,omitempty"`
	// 市对应的名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 收货人 固定电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 区对应名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 收货人地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 省对应的名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 省对应编码
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 市对应编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 街道对应编码
	StreetCode int64 `json:"street_code,omitempty" xml:"street_code,omitempty"`
	// 区对应编码
	AreaCode int64 `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 区域编码
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}

var poolExternalReceiverRequest = sync.Pool{
	New: func() any {
		return new(ExternalReceiverRequest)
	},
}

// GetExternalReceiverRequest() 从对象池中获取ExternalReceiverRequest
func GetExternalReceiverRequest() *ExternalReceiverRequest {
	return poolExternalReceiverRequest.Get().(*ExternalReceiverRequest)
}

// ReleaseExternalReceiverRequest 释放ExternalReceiverRequest
func ReleaseExternalReceiverRequest(v *ExternalReceiverRequest) {
	v.ContactName = ""
	v.StreetName = ""
	v.MobilePhone = ""
	v.Post = ""
	v.CityName = ""
	v.Phone = ""
	v.AreaName = ""
	v.DetailAddress = ""
	v.ProvinceName = ""
	v.ProvinceCode = 0
	v.CityCode = 0
	v.StreetCode = 0
	v.AreaCode = 0
	v.DivisionId = 0
	poolExternalReceiverRequest.Put(v)
}

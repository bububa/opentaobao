package ascp

import (
	"sync"
)

// CollectResourceDeleteRequest 结构体
type CollectResourceDeleteRequest struct {
	// 行政地址id（菜鸟地址库id）
	AddressIds []string `json:"address_ids,omitempty" xml:"address_ids>string,omitempty"`
	// 中文地址信息
	AddressNames []AddressNames `json:"address_names,omitempty" xml:"address_names>address_names,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 服务类型：1-上门取退
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 能力：1-上门取退
	AbilityType string `json:"ability_type,omitempty" xml:"ability_type,omitempty"`
	// 服务范围地址类型：1-行政地址；2-电子围栏
	ServiceScopeType string `json:"service_scope_type,omitempty" xml:"service_scope_type,omitempty"`
	// 如果传入为行政地址，行政地址传入类型 1- 菜鸟地址库ID 传入 2- 中文地址传入
	AddressType string `json:"address_type,omitempty" xml:"address_type,omitempty"`
	// 电子围栏外部ID（服务商配资源下必须全局唯一）
	RegionCode string `json:"region_code,omitempty" xml:"region_code,omitempty"`
	// 时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolCollectResourceDeleteRequest = sync.Pool{
	New: func() any {
		return new(CollectResourceDeleteRequest)
	},
}

// GetCollectResourceDeleteRequest() 从对象池中获取CollectResourceDeleteRequest
func GetCollectResourceDeleteRequest() *CollectResourceDeleteRequest {
	return poolCollectResourceDeleteRequest.Get().(*CollectResourceDeleteRequest)
}

// ReleaseCollectResourceDeleteRequest 释放CollectResourceDeleteRequest
func ReleaseCollectResourceDeleteRequest(v *CollectResourceDeleteRequest) {
	v.AddressIds = v.AddressIds[:0]
	v.AddressNames = v.AddressNames[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.DeliveryCode = ""
	v.ServiceType = ""
	v.AbilityType = ""
	v.ServiceScopeType = ""
	v.AddressType = ""
	v.RegionCode = ""
	v.RequestTime = 0
	poolCollectResourceDeleteRequest.Put(v)
}

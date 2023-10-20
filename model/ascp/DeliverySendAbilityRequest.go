package ascp

import (
	"sync"
)

// DeliverySendAbilityRequest 结构体
type DeliverySendAbilityRequest struct {
	// 当地址传入方式为菜鸟地址库ID传入时，则必须填写 地址id（数组）；四级/三级/二级
	AddressId []string `json:"address_id,omitempty" xml:"address_id>string,omitempty"`
	// 当地址传入方式为详细地址时填写
	AddressName []AddressName `json:"address_name,omitempty" xml:"address_name>address_name,omitempty"`
	// 地理围栏（经纬度），电子围栏最多5000个点（数组）
	RegionId []RegionId `json:"region_id,omitempty" xml:"region_id>region_id,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商配资源唯一编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 网点ID
	SiteId string `json:"site_id,omitempty" xml:"site_id,omitempty"`
	// 网点地址
	SiteAddress string `json:"site_address,omitempty" xml:"site_address,omitempty"`
	// 能力：1-送货上门
	AbilityType string `json:"ability_type,omitempty" xml:"ability_type,omitempty"`
	// 服务：1-送货上门
	ServiceClass string `json:"service_class,omitempty" xml:"service_class,omitempty"`
	// 品类：0-全部（默认）；1-中小件；2-大件；
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 服务时间
	ServiceTime string `json:"service_time,omitempty" xml:"service_time,omitempty"`
	// 服务容量：包裹件数（不传，默认无上限）
	ServiceCap string `json:"service_cap,omitempty" xml:"service_cap,omitempty"`
	// 服务范围地址类型：1-行政地址；2-地理围栏
	ServiceScopeType string `json:"service_scope_type,omitempty" xml:"service_scope_type,omitempty"`
	// 地址传入方式 1.菜鸟地址库ID传入 2.中文地址传入
	AddressType string `json:"address_type,omitempty" xml:"address_type,omitempty"`
	// 电子围栏外部ID
	RegionCode string `json:"region_code,omitempty" xml:"region_code,omitempty"`
	// 时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolDeliverySendAbilityRequest = sync.Pool{
	New: func() any {
		return new(DeliverySendAbilityRequest)
	},
}

// GetDeliverySendAbilityRequest() 从对象池中获取DeliverySendAbilityRequest
func GetDeliverySendAbilityRequest() *DeliverySendAbilityRequest {
	return poolDeliverySendAbilityRequest.Get().(*DeliverySendAbilityRequest)
}

// ReleaseDeliverySendAbilityRequest 释放DeliverySendAbilityRequest
func ReleaseDeliverySendAbilityRequest(v *DeliverySendAbilityRequest) {
	v.AddressId = v.AddressId[:0]
	v.AddressName = v.AddressName[:0]
	v.RegionId = v.RegionId[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.DeliveryCode = ""
	v.SiteId = ""
	v.SiteAddress = ""
	v.AbilityType = ""
	v.ServiceClass = ""
	v.Category = ""
	v.ServiceTime = ""
	v.ServiceCap = ""
	v.ServiceScopeType = ""
	v.AddressType = ""
	v.RegionCode = ""
	v.RequestTime = 0
	poolDeliverySendAbilityRequest.Put(v)
}

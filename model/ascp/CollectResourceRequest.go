package ascp

import (
	"sync"
)

// CollectResourceRequest 结构体
type CollectResourceRequest struct {
	// 行政地址id（菜鸟地址库id）
	AddressIds []string `json:"address_ids,omitempty" xml:"address_ids>string,omitempty"`
	// 中文地址信息
	AddressNames []AddressNames `json:"address_names,omitempty" xml:"address_names>address_names,omitempty"`
	// 电子围栏（经纬度），电子围栏最多5000个点（数组）
	RegionIds []RegionIds `json:"region_ids,omitempty" xml:"region_ids>region_ids,omitempty"`
	// 指定日期服务能力
	SpecifyDateWorkAbility []SpecifyDateWorkAbility `json:"specify_date_work_ability,omitempty" xml:"specify_date_work_ability>specify_date_work_ability,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 网点编码（仅一个）
	SiteCode string `json:"site_code,omitempty" xml:"site_code,omitempty"`
	// 网点名称（仅一个）
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
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
	// 电子围栏内包含的三级地址id（菜鸟地址库ID），电子围栏内包含多个三级地址时，需传多个，以英文逗号分隔
	RegionAddressId string `json:"region_address_id,omitempty" xml:"region_address_id,omitempty"`
	// 日常上班时间(HH:MM:SS)
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 日常下班时间(HH:MM:SS)
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 日常即时单上门能力，枚举： 1 - 1小时内上门 2 - 2小时内上门 3 - 3小时内上门 4 - 4小时内上门 5 - 不支持即时单 未填写时，默认可2小时内上门
	ImmediateCollectAbility int64 `json:"immediate_collect_ability,omitempty" xml:"immediate_collect_ability,omitempty"`
	// 日常预约单上门能力，枚举： 1 - 1小时预约单可上门 2 - 2小时预约单可上门 3 - 半天预约单（上下午）可上门 4 - 当天预约单可上门
	ReservationAbility int64 `json:"reservation_ability,omitempty" xml:"reservation_ability,omitempty"`
}

var poolCollectResourceRequest = sync.Pool{
	New: func() any {
		return new(CollectResourceRequest)
	},
}

// GetCollectResourceRequest() 从对象池中获取CollectResourceRequest
func GetCollectResourceRequest() *CollectResourceRequest {
	return poolCollectResourceRequest.Get().(*CollectResourceRequest)
}

// ReleaseCollectResourceRequest 释放CollectResourceRequest
func ReleaseCollectResourceRequest(v *CollectResourceRequest) {
	v.AddressIds = v.AddressIds[:0]
	v.AddressNames = v.AddressNames[:0]
	v.RegionIds = v.RegionIds[:0]
	v.SpecifyDateWorkAbility = v.SpecifyDateWorkAbility[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.DeliveryCode = ""
	v.SiteCode = ""
	v.SiteName = ""
	v.ServiceType = ""
	v.AbilityType = ""
	v.ServiceScopeType = ""
	v.AddressType = ""
	v.RegionCode = ""
	v.RegionAddressId = ""
	v.BeginTime = ""
	v.EndTime = ""
	v.RequestTime = 0
	v.ImmediateCollectAbility = 0
	v.ReservationAbility = 0
	poolCollectResourceRequest.Put(v)
}

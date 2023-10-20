package tmallservice

import (
	"sync"
)

// ServiceProviderDto 结构体
type ServiceProviderDto struct {
	// 网点编码
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
	// 网点名称
	ServiceStoreName string `json:"service_store_name,omitempty" xml:"service_store_name,omitempty"`
	// 工人电话
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 工人姓名
	WorkerName string `json:"worker_name,omitempty" xml:"worker_name,omitempty"`
	// 服务商昵称
	TpNick string `json:"tp_nick,omitempty" xml:"tp_nick,omitempty"`
	// 商家网点名称
	SellerStoreName string `json:"seller_store_name,omitempty" xml:"seller_store_name,omitempty"`
	// 商家网点编号
	SellerStoreCode string `json:"seller_store_code,omitempty" xml:"seller_store_code,omitempty"`
	// 注意：provider_code和providerId至少填写一个
	ProviderCode string `json:"provider_code,omitempty" xml:"provider_code,omitempty"`
	// 上门类型
	FulfilType string `json:"fulfil_type,omitempty" xml:"fulfil_type,omitempty"`
	// 服务网络
	ProviderType string `json:"provider_type,omitempty" xml:"provider_type,omitempty"`
	// 网点id
	ServiceStoreId int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
	// 服务商id
	TpId int64 `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
	// 商家网点id
	SellerStoreId int64 `json:"seller_store_id,omitempty" xml:"seller_store_id,omitempty"`
	// 注意：provider_code和providerId至少填写一个
	ProviderId int64 `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// 服务商id值
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolServiceProviderDto = sync.Pool{
	New: func() any {
		return new(ServiceProviderDto)
	},
}

// GetServiceProviderDto() 从对象池中获取ServiceProviderDto
func GetServiceProviderDto() *ServiceProviderDto {
	return poolServiceProviderDto.Get().(*ServiceProviderDto)
}

// ReleaseServiceProviderDto 释放ServiceProviderDto
func ReleaseServiceProviderDto(v *ServiceProviderDto) {
	v.ServiceStoreCode = ""
	v.ServiceStoreName = ""
	v.WorkerMobile = ""
	v.WorkerName = ""
	v.TpNick = ""
	v.SellerStoreName = ""
	v.SellerStoreCode = ""
	v.ProviderCode = ""
	v.FulfilType = ""
	v.ProviderType = ""
	v.ServiceStoreId = 0
	v.TpId = 0
	v.SellerStoreId = 0
	v.ProviderId = 0
	v.SupplierId = 0
	poolServiceProviderDto.Put(v)
}

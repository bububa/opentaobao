package tmallservice

import (
	"sync"
)

// ServiceProvider 结构体
type ServiceProvider struct {
	// 服务商昵称
	TpNick string `json:"tp_nick,omitempty" xml:"tp_nick,omitempty"`
	// 工人手机号
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 门店/网点编码
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
	// 工人姓名
	WorkerName string `json:"worker_name,omitempty" xml:"worker_name,omitempty"`
	// 门店/网点名称
	ServiceStoreName string `json:"service_store_name,omitempty" xml:"service_store_name,omitempty"`
	// isv服务商
	IsvTpName string `json:"isv_tp_name,omitempty" xml:"isv_tp_name,omitempty"`
	// 服务商id
	TpId int64 `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
	// 工人id
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
	// 门店/网点id
	ServiceStoreId int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
	// isv服务商
	IsvTpId int64 `json:"isv_tp_id,omitempty" xml:"isv_tp_id,omitempty"`
}

var poolServiceProvider = sync.Pool{
	New: func() any {
		return new(ServiceProvider)
	},
}

// GetServiceProvider() 从对象池中获取ServiceProvider
func GetServiceProvider() *ServiceProvider {
	return poolServiceProvider.Get().(*ServiceProvider)
}

// ReleaseServiceProvider 释放ServiceProvider
func ReleaseServiceProvider(v *ServiceProvider) {
	v.TpNick = ""
	v.WorkerMobile = ""
	v.ServiceStoreCode = ""
	v.WorkerName = ""
	v.ServiceStoreName = ""
	v.IsvTpName = ""
	v.TpId = 0
	v.WorkerId = 0
	v.ServiceStoreId = 0
	v.IsvTpId = 0
	poolServiceProvider.Put(v)
}

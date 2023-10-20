package wdk

import (
	"sync"
)

// MerchantRoutingInfoRegisterDo 结构体
type MerchantRoutingInfoRegisterDo struct {
	// 仓code，为空时路由为商家维度
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 操作状态1-注册；2-删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolMerchantRoutingInfoRegisterDo = sync.Pool{
	New: func() any {
		return new(MerchantRoutingInfoRegisterDo)
	},
}

// GetMerchantRoutingInfoRegisterDo() 从对象池中获取MerchantRoutingInfoRegisterDo
func GetMerchantRoutingInfoRegisterDo() *MerchantRoutingInfoRegisterDo {
	return poolMerchantRoutingInfoRegisterDo.Get().(*MerchantRoutingInfoRegisterDo)
}

// ReleaseMerchantRoutingInfoRegisterDo 释放MerchantRoutingInfoRegisterDo
func ReleaseMerchantRoutingInfoRegisterDo(v *MerchantRoutingInfoRegisterDo) {
	v.WarehouseCode = ""
	v.Status = 0
	poolMerchantRoutingInfoRegisterDo.Put(v)
}

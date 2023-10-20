package wlbimports

import (
	"sync"
)

// ResourceResult 结构体
type ResourceResult struct {
	// 资源代码
	ResCode string `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 魔方天下美国线
	ResourceName string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// 资源ID
	ResId int64 `json:"res_id,omitempty" xml:"res_id,omitempty"`
	// 时效，单位（小时）
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 服务报价。首重（1磅) CNY￥3元 续重（1磅) CNY￥6元
	DeliveryPrice *BSWeightPrice `json:"delivery_price,omitempty" xml:"delivery_price,omitempty"`
}

var poolResourceResult = sync.Pool{
	New: func() any {
		return new(ResourceResult)
	},
}

// GetResourceResult() 从对象池中获取ResourceResult
func GetResourceResult() *ResourceResult {
	return poolResourceResult.Get().(*ResourceResult)
}

// ReleaseResourceResult 释放ResourceResult
func ReleaseResourceResult(v *ResourceResult) {
	v.ResCode = ""
	v.ResourceName = ""
	v.ResId = 0
	v.DeliveryTime = 0
	v.DeliveryPrice = nil
	poolResourceResult.Put(v)
}

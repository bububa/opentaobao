package qimen

import (
	"sync"
)

// PresalesPackageConsignRequest 结构体
type PresalesPackageConsignRequest struct {
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 外部订单号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 支付时间
	PayTime string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展信息Map
	ExtendProps string `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
	// 金额(单位为分)
	TotalAmount int64 `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

var poolPresalesPackageConsignRequest = sync.Pool{
	New: func() any {
		return new(PresalesPackageConsignRequest)
	},
}

// GetPresalesPackageConsignRequest() 从对象池中获取PresalesPackageConsignRequest
func GetPresalesPackageConsignRequest() *PresalesPackageConsignRequest {
	return poolPresalesPackageConsignRequest.Get().(*PresalesPackageConsignRequest)
}

// ReleasePresalesPackageConsignRequest 释放PresalesPackageConsignRequest
func ReleasePresalesPackageConsignRequest(v *PresalesPackageConsignRequest) {
	v.OwnerCode = ""
	v.OrderCode = ""
	v.PayTime = ""
	v.Remark = ""
	v.ExtendProps = ""
	v.TotalAmount = 0
	poolPresalesPackageConsignRequest.Put(v)
}

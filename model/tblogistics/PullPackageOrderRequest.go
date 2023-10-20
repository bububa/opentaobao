package tblogistics

import (
	"sync"
)

// PullPackageOrderRequest 结构体
type PullPackageOrderRequest struct {
	// 业务类型，取值：JYPKZXCK(集运包裹正向出库)、JYPKNXCK(集运包裹逆向出库)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 物流主体，例如：TaoTian(淘天)
	LogisticsOwner string `json:"logistics_owner,omitempty" xml:"logistics_owner,omitempty"`
	// 服务商仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 包裹货主
	PackageOwnerCode string `json:"package_owner_code,omitempty" xml:"package_owner_code,omitempty"`
	// 服务提供商
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
}

var poolPullPackageOrderRequest = sync.Pool{
	New: func() any {
		return new(PullPackageOrderRequest)
	},
}

// GetPullPackageOrderRequest() 从对象池中获取PullPackageOrderRequest
func GetPullPackageOrderRequest() *PullPackageOrderRequest {
	return poolPullPackageOrderRequest.Get().(*PullPackageOrderRequest)
}

// ReleasePullPackageOrderRequest 释放PullPackageOrderRequest
func ReleasePullPackageOrderRequest(v *PullPackageOrderRequest) {
	v.OrderType = ""
	v.LogisticsOwner = ""
	v.WarehouseCode = ""
	v.PackageOwnerCode = ""
	v.SupplierId = ""
	v.MailNo = ""
	poolPullPackageOrderRequest.Put(v)
}

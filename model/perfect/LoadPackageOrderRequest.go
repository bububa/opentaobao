package perfect

import (
	"sync"
)

// LoadPackageOrderRequest 结构体
type LoadPackageOrderRequest struct {
	// 包裹明细
	PackageDetails []LoadPackageOrderDetailRequest `json:"package_details,omitempty" xml:"package_details>load_package_order_detail_request,omitempty"`
	// 令牌号
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 出库单
	OutboundOrderCode string `json:"outbound_order_code,omitempty" xml:"outbound_order_code,omitempty"`
	// 包裹单号，没有默认使用包裹号，即令牌号
	PackageOrderCode string `json:"package_order_code,omitempty" xml:"package_order_code,omitempty"`
	// 扩展
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
}

var poolLoadPackageOrderRequest = sync.Pool{
	New: func() any {
		return new(LoadPackageOrderRequest)
	},
}

// GetLoadPackageOrderRequest() 从对象池中获取LoadPackageOrderRequest
func GetLoadPackageOrderRequest() *LoadPackageOrderRequest {
	return poolLoadPackageOrderRequest.Get().(*LoadPackageOrderRequest)
}

// ReleaseLoadPackageOrderRequest 释放LoadPackageOrderRequest
func ReleaseLoadPackageOrderRequest(v *LoadPackageOrderRequest) {
	v.PackageDetails = v.PackageDetails[:0]
	v.PackageCode = ""
	v.OutboundOrderCode = ""
	v.PackageOrderCode = ""
	v.Attributes = ""
	poolLoadPackageOrderRequest.Put(v)
}

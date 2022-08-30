package perfect

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

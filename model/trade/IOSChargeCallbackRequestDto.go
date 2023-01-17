package trade

// IOSChargeCallbackRequestDto 结构体
type IOSChargeCallbackRequestDto struct {
	// 商家订单号
	VendorOrderNo string `json:"vendor_order_no,omitempty" xml:"vendor_order_no,omitempty"`
	// 失败的错误码
	FailedCode string `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
	// 交易猫订单号
	JymOrderNo string `json:"jym_order_no,omitempty" xml:"jym_order_no,omitempty"`
	// 商家id
	VendorId string `json:"vendor_id,omitempty" xml:"vendor_id,omitempty"`
	// 商家订单状态
	VendorOrderStatus string `json:"vendor_order_status,omitempty" xml:"vendor_order_status,omitempty"`
	// 商家订单快照
	VendorOrderSnap string `json:"vendor_order_snap,omitempty" xml:"vendor_order_snap,omitempty"`
	// 版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 商家订单充值成功时间,格式yyyyMMddHHmmss
	VendorOrderSuccessTime string `json:"vendor_order_success_time,omitempty" xml:"vendor_order_success_time,omitempty"`
	// 失败原因
	FailedReason string `json:"failed_reason,omitempty" xml:"failed_reason,omitempty"`
}

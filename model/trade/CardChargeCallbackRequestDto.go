package trade

// CardChargeCallbackRequestDto 结构体
type CardChargeCallbackRequestDto struct {
	// 卡密组（注意：卡密是密文）
	CardInfos string `json:"card_infos,omitempty" xml:"card_infos,omitempty"`
	// 失败错误码
	FailedCode string `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
	// 失败原因
	FailedReason string `json:"failed_reason,omitempty" xml:"failed_reason,omitempty"`
	// 交易猫订单号
	JymOrderNo string `json:"jym_order_no,omitempty" xml:"jym_order_no,omitempty"`
	// 订单类型，回填
	JymOrderType string `json:"jym_order_type,omitempty" xml:"jym_order_type,omitempty"`
	// 商家id
	VendorId string `json:"vendor_id,omitempty" xml:"vendor_id,omitempty"`
	// 商家订单号
	VendorOrderNo string `json:"vendor_order_no,omitempty" xml:"vendor_order_no,omitempty"`
	// 商家商品快照
	VendorOrderSnap string `json:"vendor_order_snap,omitempty" xml:"vendor_order_snap,omitempty"`
	// 商家订单状态
	VendorOrderStatus string `json:"vendor_order_status,omitempty" xml:"vendor_order_status,omitempty"`
	// 商家订单充值成功时间
	VendorOrderSuccessTime string `json:"vendor_order_success_time,omitempty" xml:"vendor_order_success_time,omitempty"`
	// 版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
}

package wdk

// SkuStockPublishParamDo 结构体
type SkuStockPublishParamDo struct {
	// 商家门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 这笔单据发生的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 唯一单据号，用于幂等操作
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 商品对应的69码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 移动的数量，正数表式增加，负数表式减少
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 更新类型，1表式覆盖，0表式增量
	UpdateType int64 `json:"update_type,omitempty" xml:"update_type,omitempty"`
	// 当时业务发生的时间戳，单位：ms
	OperationTs int64 `json:"operation_ts,omitempty" xml:"operation_ts,omitempty"`
}

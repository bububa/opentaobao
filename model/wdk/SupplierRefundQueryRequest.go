package wdk

// SupplierRefundQueryRequest 结构体
type SupplierRefundQueryRequest struct {
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 实际售卖商家code
	SourceMerchantCode string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码，从0开始
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 下单终端：APP/POS
	OrderClient string `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 订单渠道
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 1:售中退款  2:售后退款
	DisputeType int64 `json:"dispute_type,omitempty" xml:"dispute_type,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

package wdk

// WarehouseOrderGetRequest 结构体
type WarehouseOrderGetRequest struct {
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

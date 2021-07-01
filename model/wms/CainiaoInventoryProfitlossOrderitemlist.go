package wms

// CainiaoInventoryProfitlossOrderitemlist 结构体
type CainiaoInventoryProfitlossOrderitemlist struct {
	// 损益详情
	OrderItem *CainiaoInventoryProfitlossOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

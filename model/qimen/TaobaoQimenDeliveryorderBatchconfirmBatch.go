package qimen

// TaobaoQimenDeliveryorderBatchconfirmBatch 结构体
type TaobaoQimenDeliveryorderBatchconfirmBatch struct {
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 实发数量(要求batchs节点下所有的实发数量之和等于orderline中的实发数量)
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
}

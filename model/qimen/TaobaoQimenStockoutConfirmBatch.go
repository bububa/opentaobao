package qimen

// TaobaoQimenStockoutConfirmBatch 结构体
type TaobaoQimenStockoutConfirmBatch struct {
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 库存类型
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 实发数量
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
}

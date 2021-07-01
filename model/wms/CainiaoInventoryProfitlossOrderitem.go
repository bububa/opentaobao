package wms

// CainiaoInventoryProfitlossOrderitem 结构体
type CainiaoInventoryProfitlossOrderitem struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家对商品的编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 库存类型 1 可销售库存（正品）  101 残次
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 商品数量
	ItemQty string `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
	// 商品保质期信息，失效日期
	DueDate string `json:"due_date,omitempty" xml:"due_date,omitempty"`
	// 商品保质期信息，生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 生产编码，同一商品可能因商家不同有不同编码
	ProduceCode string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
	// 生产地区
	ProduceArea string `json:"produce_area,omitempty" xml:"produce_area,omitempty"`
	// 批次号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
}

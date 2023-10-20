package ascp

// AlibabadchainaoxiangwmsdeliveryorderconfirmBatch 结构体
type AlibabadchainaoxiangwmsdeliveryorderconfirmBatch struct {
	// 批次编号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"product_date,omitempty" xml:"product_date,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 生产批号
	ProduceCode string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 实发数量(要求batchs节点下所有的实发数量之和等于orderline中的实发数量)
	ActualQty string `json:"actual_qty,omitempty" xml:"actual_qty,omitempty"`
	// 货品sn编码
	SnCode string `json:"sn_code,omitempty" xml:"sn_code,omitempty"`
}

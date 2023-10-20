package tbk

// Ucrowdrankitems 结构体
type Ucrowdrankitems struct {
	// 物料评估-商品ID，material_id=41377时必填
	Itemid string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 物料评估-商品佣金率，如：1234表示12.34%，material_id=41377时选填
	Commirate int64 `json:"commirate,omitempty" xml:"commirate,omitempty"`
	// 物料评估-商品价格，单位：元，material_id=41377时选填
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
}

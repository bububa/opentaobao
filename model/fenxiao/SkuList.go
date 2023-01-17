package fenxiao

// SkuList 结构体
type SkuList struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 市场价
	StandardPrice string `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
	// 代销采购价，单位：元
	CostPrice string `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
	// 经销采购价
	DealerCostPrice string `json:"dealer_cost_price,omitempty" xml:"dealer_cost_price,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// sku的销售属性组合字符串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// SkuID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 关联的后端商品id
	ScitemId int64 `json:"scitem_id,omitempty" xml:"scitem_id,omitempty"`
	// 预扣库存
	ReservedQuantity int64 `json:"reserved_quantity,omitempty" xml:"reserved_quantity,omitempty"`
	// 配额可用库存
	QuotaQuantity int64 `json:"quota_quantity,omitempty" xml:"quota_quantity,omitempty"`
}

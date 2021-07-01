package fenxiao

// DealerOrderDetail 结构体
type DealerOrderDetail struct {
	// 经销采购单编号
	DealerOrderId int64 `json:"dealer_order_id,omitempty" xml:"dealer_order_id,omitempty"`
	// 经销采购单明细id
	DealerDetailId int64 `json:"dealer_detail_id,omitempty" xml:"dealer_detail_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品标题
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 产品规格
	SkuSpec string `json:"sku_spec,omitempty" xml:"sku_spec,omitempty"`
	// 商家编码，是sku则为sku的商家编码，否则是产品的商家编码
	SkuNumber string `json:"sku_number,omitempty" xml:"sku_number,omitempty"`
	// 原始价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	OriginalPrice float64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 最终价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	FinalPrice float64 `json:"final_price,omitempty" xml:"final_price,omitempty"`
	// 采购数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 金额小计(采购数量乘以最终价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	PriceCount float64 `json:"price_count,omitempty" xml:"price_count,omitempty"`
	// 该条明细是否已删除。true：已删除；false：未删除。
	IsDeleted bool `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 产品快照url
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 属性列表，key-value形式。
	Features []Feature `json:"features,omitempty" xml:"features>feature,omitempty"`
}

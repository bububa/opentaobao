package tmallcar

// TopOrderDto 结构体
type TopOrderDto struct {
	// 核销时间
	ConsumeTime string `json:"consume_time,omitempty" xml:"consume_time,omitempty"`
	// 核销门店
	ConsumeStore string `json:"consume_store,omitempty" xml:"consume_store,omitempty"`
	// 预约门店
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// sku属性：颜色分类
	SkuFeature1 string `json:"sku_feature1,omitempty" xml:"sku_feature1,omitempty"`
	// sku属性：套餐名称
	SkuFeature2 string `json:"sku_feature2,omitempty" xml:"sku_feature2,omitempty"`
	// sku_feature_x
	SkuFeatureX string `json:"sku_feature_x,omitempty" xml:"sku_feature_x,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

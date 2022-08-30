package traveltrade

// BuyItemInfo 结构体
type BuyItemInfo struct {
	// 商品类目相关的扩展信息（不再推荐使用，建议使用category_ext_infos_json替代）。KV对形式，多个KV对以英文封号分隔，k1:v1;k2:v2。各类目支持的枚举key详见：https://open.alitrip.com/docs/doc.htm?docType=1&articleId=107548
	CategoryExtInfos string `json:"category_ext_infos,omitempty" xml:"category_ext_infos,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 商家自定义的商品编码
	OutProductId string `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
	// sku_id对应的商家自定义sku编码（即商品上套餐的外部商家编码）
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 用户所购买的Sku信息。包含商品套餐名称（对于门票商品即指票种），出游人群信息等
	SkuProperties string `json:"sku_properties,omitempty" xml:"sku_properties,omitempty"`
	// 出行结束日期，如果没有出行结束日期，则该值为空。（对于wifi/电话卡商品 指激活/租用结束日期；对于流量充值商品 指结束使用日期）
	TripEndDate string `json:"trip_end_date,omitempty" xml:"trip_end_date,omitempty"`
	// 出行开始日期。（对于wifi/电话卡商品 指激活/租用日期；对于门票商品 指入园日期；对于流量充值商品 指开始使用日期）
	TripStartDate string `json:"trip_start_date,omitempty" xml:"trip_start_date,omitempty"`
	// 商品图片
	ItemImage string `json:"item_image,omitempty" xml:"item_image,omitempty"`
	// 商品类目相关的扩展信息，JSON格式。各类目支持的枚举key详见：https://open.alitrip.com/docs/doc.htm?docType=1&articleId=107548
	CategoryExtInfosJson string `json:"category_ext_infos_json,omitempty" xml:"category_ext_infos_json,omitempty"`
	// 场次开始时间
	FsEndTime string `json:"fs_end_time,omitempty" xml:"fs_end_time,omitempty"`
	// 场次结束时间
	FsStartTime string `json:"fs_start_time,omitempty" xml:"fs_start_time,omitempty"`
	// 商品所属叶子类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 购买的数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 商品价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 用户所购买的商品上sku的id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

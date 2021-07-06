package aedropshiper

// AeopFindProductResultDto 结构体
type AeopFindProductResultDto struct {
	// 商品的SKU信息
	AeopAeProductSKUs []AeopAeProductSku `json:"aeop_ae_product_s_k_us,omitempty" xml:"aeop_ae_product_s_k_us>aeop_ae_product_sku,omitempty"`
	// 商品的类目属性
	AeopAeProductPropertys []AeopAeProductProperty `json:"aeop_ae_product_propertys,omitempty" xml:"aeop_ae_product_propertys>aeop_ae_product_property,omitempty"`
	// 商品详描
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 产品的下架日期
	WsOfflineDate string `json:"ws_offline_date,omitempty" xml:"ws_offline_date,omitempty"`
	// 产品的下架原因
	WsDisplay string `json:"ws_display,omitempty" xml:"ws_display,omitempty"`
	// 商品拥有者的login_id
	OwnerMemberId string `json:"owner_member_id,omitempty" xml:"owner_member_id,omitempty"`
	// 产品的状态
	ProductStatusType string `json:"product_status_type,omitempty" xml:"product_status_type,omitempty"`
	// 产品的毛重
	GrossWeight string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 产品的标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 无线详描
	MobileDetail string `json:"mobile_detail,omitempty" xml:"mobile_detail,omitempty"`
	// 产品的货币单位。美元: USD, 卢布: RUB
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 产品创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 产品的主图列表
	ImageURLs string `json:"image_u_r_ls,omitempty" xml:"image_u_r_ls,omitempty"`
	// 单品产品的价格。
	ProductPrice string `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 网站折扣后价格
	ItemOfferSiteSalePrice string `json:"item_offer_site_sale_price,omitempty" xml:"item_offer_site_sale_price,omitempty"`
	// 平均评价星级，1-5星
	AvgEvaluationRating string `json:"avg_evaluation_rating,omitempty" xml:"avg_evaluation_rating,omitempty"`
	// 产品的单位
	ProductUnit int64 `json:"product_unit,omitempty" xml:"product_unit,omitempty"`
	// 产品所在类目的ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 商品多媒体信息，该属性主要包含商品的视频列表
	AeopAEMultimedia *AeopAeMultimedia `json:"aeop_a_e_multimedia,omitempty" xml:"aeop_a_e_multimedia,omitempty"`
	// 商品的备货期
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 产品的有效期
	WsValidNum int64 `json:"ws_valid_num,omitempty" xml:"ws_valid_num,omitempty"`
	// 商品分国家报价的配置
	AeopNationalQuoteConfiguration *AeopNationalQuoteConfiguration `json:"aeop_national_quote_configuration,omitempty" xml:"aeop_national_quote_configuration,omitempty"`
	// 自定义计重的基本产品件数
	BaseUnit int64 `json:"base_unit,omitempty" xml:"base_unit,omitempty"`
	// 产品的长度
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// 产品的高度
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// 产品的宽度
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 错误代码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 库存
	TotalAvailableStock int64 `json:"total_available_stock,omitempty" xml:"total_available_stock,omitempty"`
	// 店铺信息
	StoreInfo *AeopStoreInfo `json:"store_info,omitempty" xml:"store_info,omitempty"`
	// 评价数
	EvaluationCount int64 `json:"evaluation_count,omitempty" xml:"evaluation_count,omitempty"`
	// 订单数
	OrderCount int64 `json:"order_count,omitempty" xml:"order_count,omitempty"`
	// 请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 包装类型
	PackageType bool `json:"package_type,omitempty" xml:"package_type,omitempty"`
}

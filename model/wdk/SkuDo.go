package wdk

// SkuDo 结构体
type SkuDo struct {
	// 供货商信息
	SkuSuppliers []SkuSupplierDo `json:"sku_suppliers,omitempty" xml:"sku_suppliers>sku_supplier_do,omitempty"`
	// 渠道属性
	ChannelProps []ChannelProp `json:"channel_props,omitempty" xml:"channel_props>channel_prop,omitempty"`
	// 子商品信息
	SubSkus []SubSkuDo `json:"sub_skus,omitempty" xml:"sub_skus>sub_sku_do,omitempty"`
	// 条码，支持一品多码,多个条码以半角逗号分隔
	Barcodes string `json:"barcodes,omitempty" xml:"barcodes,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 商家后台类目编码
	CategoryCode string `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 门店或DC编码
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 高度（高）
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 库存单位
	InventoryUnit string `json:"inventory_unit,omitempty" xml:"inventory_unit,omitempty"`
	// 长度(深)
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 主图图片地址
	MainPicUrls string `json:"main_pic_urls,omitempty" xml:"main_pic_urls,omitempty"`
	// 生产商地址
	ManufacturerAddress string `json:"manufacturer_address,omitempty" xml:"manufacturer_address,omitempty"`
	// 生产商名称
	ManufacturerName string `json:"manufacturer_name,omitempty" xml:"manufacturer_name,omitempty"`
	// 详情图片地址
	DetailPicUrls string `json:"detail_pic_urls,omitempty" xml:"detail_pic_urls,omitempty"`
	// 产地，商品生产地点的描述，用于APP和电子价签展示，如果是可售，则必填
	ProducerPlace string `json:"producer_place,omitempty" xml:"producer_place,omitempty"`
	// 商品唯一编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 名称，对商品直观的描述，通常包含了品牌、规格等信息
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 产品标准号
	StandardNo string `json:"standard_no,omitempty" xml:"standard_no,omitempty"`
	// 建议零售价，单位:元
	SuggestedPrice string `json:"suggested_price,omitempty" xml:"suggested_price,omitempty"`
	// 销项税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 单品重量 单位为克/g，必须为整数
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 宽度（宽）
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 平台后台类目编码
	HmCategoryCode string `json:"hm_category_code,omitempty" xml:"hm_category_code,omitempty"`
	// 商品短标题
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 文描
	TxtDesc string `json:"txt_desc,omitempty" xml:"txt_desc,omitempty"`
	// 开票内容
	InvoiceContent string `json:"invoice_content,omitempty" xml:"invoice_content,omitempty"`
	// 输入开票内容
	InputInvoiceContent string `json:"input_invoice_content,omitempty" xml:"input_invoice_content,omitempty"`
	// 财务核算分类
	AccountingCategory string `json:"accounting_category,omitempty" xml:"accounting_category,omitempty"`
	// 净含量
	NetContent string `json:"net_content,omitempty" xml:"net_content,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 商品销售价格，单位:元
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 商品在机构内的生命周期，商品状态；A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡
	LifeStatus string `json:"life_status,omitempty" xml:"life_status,omitempty"`
	// 销售规格
	SaleSpec string `json:"sale_spec,omitempty" xml:"sale_spec,omitempty"`
	// 均重（一个售卖单位平均为多少个库存单位，称重商品必填)，计算库存发布时会以库存数量/均重来进行转换；建议与库存单位保持一致。若库存单位是kg，售卖单位可以是g。此时均重、预扣款重量填0.001。APP最小起购量、APP购买步长填正整数。比如香蕉的库存单位是kg，销售单位是g，APP最小起购量、APP购买步长填500，表示在APP最少购买一斤，每加一档是加一斤。均重、预扣款重量填0.001的作用是，当用户购买500g，乘以0.001系数后=0.5kg，所以扣库存0.5kg
	AvgWeight string `json:"avg_weight,omitempty" xml:"avg_weight,omitempty"`
	// 预扣款重量（购买一个售卖单位按照多少个库存单位来扣款），计算金额时，按照购买的售卖单位数量/预扣款重量*售价 来计算；非称重品填1。若库存单位是kg，销售单位是g，填0.001
	PreMinusWeight string `json:"pre_minus_weight,omitempty" xml:"pre_minus_weight,omitempty"`
	// 标价签类型
	SkuLabelType string `json:"sku_label_type,omitempty" xml:"sku_label_type,omitempty"`
	// 商品卖点；商品副标题，显示在APP商品详情页的标题下方。如可口可乐商品的副标题可以是“夏日必备 解暑神器”。不超过80个字符。双方业务沟通一下，是否生鲜商品要定为必填
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 卖点1内容
	SubTitle1 string `json:"sub_title1,omitempty" xml:"sub_title1,omitempty"`
	// 卖点2内容
	SubTitle2 string `json:"sub_title2,omitempty" xml:"sub_title2,omitempty"`
	// 卖点1名称
	Title1 string `json:"title1,omitempty" xml:"title1,omitempty"`
	// 卖点2名称
	Title2 string `json:"title2,omitempty" xml:"title2,omitempty"`
	// 大仓向门店配货的单位；淘鲜达合作商家默认填与库存单位一致的值
	DeliveryUnit string `json:"delivery_unit,omitempty" xml:"delivery_unit,omitempty"`
	// 一个配货单位等于多少个库存单位；淘鲜达合作商家填默认值1
	DeliverySpec string `json:"delivery_spec,omitempty" xml:"delivery_spec,omitempty"`
	// 商品到本仓的来源；淘鲜达商家填默认值“直配”(1:统配，2:直配，3:越库，4：自产；5：调拨)， 基于采配链路的要求，加工品是不能够做直配的物流模式的，必须维护自产或者统配，目前包了一层逻辑“针对加工成品和加工半成品，如果是直配就转成自产”
	DeliveryWay string `json:"delivery_way,omitempty" xml:"delivery_way,omitempty"`
	// 商品关联的物流中心，DC。配送物流；配送方式是直配，此项不填（输入物流中心编码）
	Logistics string `json:"logistics,omitempty" xml:"logistics,omitempty"`
	// 配出仓,配货仓；配送方式是直配，此项不填(输入仓库编码)
	DeliveryWarehouse string `json:"delivery_warehouse,omitempty" xml:"delivery_warehouse,omitempty"`
	// 会员正常购买该商品的售价，2位小数，单位:元
	MemberPrice string `json:"member_price,omitempty" xml:"member_price,omitempty"`
	// 商品作为原料时对应的单位；淘鲜达合作商家默认填与库存单位一致的值
	CostUnit string `json:"cost_unit,omitempty" xml:"cost_unit,omitempty"`
	// 商品加工所消耗的直接原料成本，去税；加工商品必填，加工品按照原料的消耗估算出的成本。因为淘鲜达暂不涉及成本计算，建议淘鲜达商家填默认值1，单位:元
	CostNoPrice string `json:"cost_no_price,omitempty" xml:"cost_no_price,omitempty"`
	// 商品加工所消耗的直接原料成本，含税，加工商品必填，加工品按照原料的消耗估算出的成本。因为淘鲜达暂不涉及成本计算，建议淘鲜达商家填默认值1，单位:元
	CostTaxPrice string `json:"cost_tax_price,omitempty" xml:"cost_tax_price,omitempty"`
	// 一个库存单位对应多少个成本单位，原料加工的成本单位和存储单位之间的换算。因为淘鲜达暂不涉及成本计算，建议淘鲜达商家填默认值1
	CostExchangeRate string `json:"cost_exchange_rate,omitempty" xml:"cost_exchange_rate,omitempty"`
	// 商品中对于加工可用的比例占整个商品的百分比；从原料到加工可用的原料之间的转化率。为了简化BOM的使用，建议淘鲜达商家填默认值1。
	MassOutputRate string `json:"mass_output_rate,omitempty" xml:"mass_output_rate,omitempty"`
	// 每一组加工生产的单位，加工商品必填，仅当商品为加工品的时候，才能填写。从计量单位表中选择；淘鲜达合作商家默认填与库存单位一致的值
	ProcessingUnit string `json:"processing_unit,omitempty" xml:"processing_unit,omitempty"`
	// 一组加工单位对应多少个库存单位商品；加工商品必填，每一个加工单位对应的存储单位数量。为了简化BOM的使用，建议淘鲜达商家填默认值1。
	ProcsExchangeRate string `json:"procs_exchange_rate,omitempty" xml:"procs_exchange_rate,omitempty"`
	// 存储条件；填常温、冷藏、冷冻、热链、鲜活
	Storage string `json:"storage,omitempty" xml:"storage,omitempty"`
	// 商品拣货时可以超拣的比例，针对非标品。建议淘鲜达商家填默认值0
	PickFloatRate string `json:"pick_float_rate,omitempty" xml:"pick_float_rate,omitempty"`
	// 商品收货时可以超过订货数量多少百分比进行收货；非标品供应商送货入库时，允许超收的比例。淘鲜达合作商家填默认值0
	OverloadRate string `json:"overload_rate,omitempty" xml:"overload_rate,omitempty"`
	// 一个采购单位等于多少个库存单位，淘鲜达合作商家默认填1
	PurchaseSpec string `json:"purchase_spec,omitempty" xml:"purchase_spec,omitempty"`
	// 商品针对供应商订货时，对应的单位；淘鲜达合作商家默认填与库存单位一致的值
	PurchaseUnit string `json:"purchase_unit,omitempty" xml:"purchase_unit,omitempty"`
	// 标价签类型；商品在门店陈列时，采用的售价标签类型。因不采用电子价签，建议淘鲜达商家填默认值“无价签”
	LabelStyleType string `json:"label_style_type,omitempty" xml:"label_style_type,omitempty"`
	// 文描,（同字段txt_desc）；优先使用本字段
	RichText string `json:"rich_text,omitempty" xml:"rich_text,omitempty"`
	// 平台类目编码（同字段hm_category_code），需要商家把自己的类目对应到平台的类目上，此字段暂时只能支持3、4级类目修改；优先使用本字段
	BackCatCode string `json:"back_cat_code,omitempty" xml:"back_cat_code,omitempty"`
	// 门店商品售价（同字段sale_price），单位:元；优先使用本字段
	SkuPrice string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// 商家后台类目编码（同字段category_code），优先使用本字段
	MerchantCatCode string `json:"merchant_cat_code,omitempty" xml:"merchant_cat_code,omitempty"`
	// 在app上显示的商品图片cdn地址（同字段main_pic_urls和detail_pic_urls），https开头，多个图片用英文逗号切割。第一张图片会作为主图；优先使用本字段
	SkuPicUrls string `json:"sku_pic_urls,omitempty" xml:"sku_pic_urls,omitempty"`
	// 税收编码
	TaxCode string `json:"tax_code,omitempty" xml:"tax_code,omitempty"`
	// 配送条件；填常温、冷藏、冷冻
	DeliveryStorage string `json:"delivery_storage,omitempty" xml:"delivery_storage,omitempty"`
	// 淘宝旗舰店同品信息，同城零售使用
	FlagshipStoreItemInfo string `json:"flagship_store_item_info,omitempty" xml:"flagship_store_item_info,omitempty"`
	// 产品聚合码，程序会把聚合码相同的sku聚合到一起；比如衣服有大中小号3中sku，通过聚合码聚合到一起；该字段在同城零售发布场景才能使用
	ProductIdentity string `json:"product_identity,omitempty" xml:"product_identity,omitempty"`
	// 业务类型：txd淘鲜达，elm饿了么，shareStore共享库存，默认txd
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 1一品多码（只支持非称重品）
	SpecType string `json:"spec_type,omitempty" xml:"spec_type,omitempty"`
	// 与母商品的库存转换系数，例如10，则转换到母商品的库存为10*此商品的库存，可空，当新增一品多规格品的时候不能为空
	TransRatio string `json:"trans_ratio,omitempty" xml:"trans_ratio,omitempty"`
	// 母商品的sku_code可空，当新增一品多规格品的时候不能为空
	ParentSkuCode string `json:"parent_sku_code,omitempty" xml:"parent_sku_code,omitempty"`
	// 线上生效开始时间
	AppEffectBeginTime string `json:"app_effect_begin_time,omitempty" xml:"app_effect_begin_time,omitempty"`
	// 线上生效过期时间
	AppEffectEndTime string `json:"app_effect_end_time,omitempty" xml:"app_effect_end_time,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 机构编号
	OrgNo string `json:"org_no,omitempty" xml:"org_no,omitempty"`
	// 商品别名
	AliasName string `json:"alias_name,omitempty" xml:"alias_name,omitempty"`
	// 生产商名称
	ProducerName string `json:"producer_name,omitempty" xml:"producer_name,omitempty"`
	// 厂方货号
	FactoryNo string `json:"factory_no,omitempty" xml:"factory_no,omitempty"`
	// 成份
	Component string `json:"component,omitempty" xml:"component,omitempty"`
	// 产品等级
	Grade string `json:"grade,omitempty" xml:"grade,omitempty"`
	// 食用方式
	EatWay string `json:"eat_way,omitempty" xml:"eat_way,omitempty"`
	// 溯源国标码
	OriginCode string `json:"origin_code,omitempty" xml:"origin_code,omitempty"`
	// 商品进项税率
	InputTaxRate string `json:"input_tax_rate,omitempty" xml:"input_tax_rate,omitempty"`
	// 开票内容
	InputTaxRateCode string `json:"input_tax_rate_code,omitempty" xml:"input_tax_rate_code,omitempty"`
	// 财务核算分类编码
	FinanceTypeCode string `json:"finance_type_code,omitempty" xml:"finance_type_code,omitempty"`
	// 开票内容
	TaxRateCode string `json:"tax_rate_code,omitempty" xml:"tax_rate_code,omitempty"`
	// 采配编码
	PurchaseUnitCode string `json:"purchase_unit_code,omitempty" xml:"purchase_unit_code,omitempty"`
	// 存储单位编码
	InventoryUnitCode string `json:"inventory_unit_code,omitempty" xml:"inventory_unit_code,omitempty"`
	// 货物验收标准
	AcceptanceCriteria string `json:"acceptance_criteria,omitempty" xml:"acceptance_criteria,omitempty"`
	// 配货单位编码
	DeliveryUnitCode string `json:"delivery_unit_code,omitempty" xml:"delivery_unit_code,omitempty"`
	// 加工单位编码
	ProcessingUnitCode string `json:"processing_unit_code,omitempty" xml:"processing_unit_code,omitempty"`
	// 成本单位编码
	CostUnitCode string `json:"cost_unit_code,omitempty" xml:"cost_unit_code,omitempty"`
	// 商品标编码
	SkuLabelTypeCode string `json:"sku_label_type_code,omitempty" xml:"sku_label_type_code,omitempty"`
	// 销售单位（取值选项和采购单位的可选项一致）
	SaleUnitCode string `json:"sale_unit_code,omitempty" xml:"sale_unit_code,omitempty"`
	// 标价签类型编码
	LabelPriceTypeCode string `json:"label_price_type_code,omitempty" xml:"label_price_type_code,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 发布渠道
	ChannelCodes string `json:"channel_codes,omitempty" xml:"channel_codes,omitempty"`
	// 默认供应商
	SupplierNo string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
	// 默认供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 默认物流中心名称
	LogisticsName string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
	// 配货逻辑仓名称
	DeliveryWarehouseName string `json:"delivery_warehouse_name,omitempty" xml:"delivery_warehouse_name,omitempty"`
	// 退货逻辑仓
	ReturnWarehouse string `json:"return_warehouse,omitempty" xml:"return_warehouse,omitempty"`
	// 退货逻辑仓名称
	ReturnWarehouseName string `json:"return_warehouse_name,omitempty" xml:"return_warehouse_name,omitempty"`
	// 配送方式名称
	DeliveryWayName string `json:"delivery_way_name,omitempty" xml:"delivery_way_name,omitempty"`
	// 服务商品,skucode
	ServiceItems string `json:"service_items,omitempty" xml:"service_items,omitempty"`
	// 默认商品价格（同SKU在当前渠道的价格）
	DefaultSkuPrice string `json:"default_sku_price,omitempty" xml:"default_sku_price,omitempty"`
	// 存储条件编码
	StorageCode string `json:"storage_code,omitempty" xml:"storage_code,omitempty"`
	// 渠道编码，默认为4淘鲜达渠道、3对应饿了么渠道、-1对应批发渠道
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 保质期天数，商品的保质期（单位：天）,0表示没有保质期
	ShelfLife int64 `json:"shelf_life,omitempty" xml:"shelf_life,omitempty"`
	// 类型 1:普通商品、2:加工半成品、3:加工成品、4:原材料、5:耗材； 如果是新增商家产品，字段含义是货品类型，1普通商品、2赠品、3包材、4耗材、5组合商品、6分销商品、7附属品、8虚拟商品、9其他、10直投广告、11原材料、13加工产成品
	SkuType int64 `json:"sku_type,omitempty" xml:"sku_type,omitempty"`
	// 存储条件类型 241=常温,242=冷藏,243=冷冻,635=热链,636=室温,637=鲜活
	StorageType int64 `json:"storage_type,omitempty" xml:"storage_type,omitempty"`
	// 是否称重 标识商品是否是称重商品? 1：是0：否（默认为0）
	WeightFlag int64 `json:"weight_flag,omitempty" xml:"weight_flag,omitempty"`
	// 是否进口，是1，否0，默认非进口
	ImportFlag int64 `json:"import_flag,omitempty" xml:"import_flag,omitempty"`
	// app购买时每增加一次购买数量至少要增加多少个售卖单位。非称重品填1。称重品根据实际货品情况填写。步长除了是加购物车的数量之外还承载了最小售卖数量的意义。
	StepQuantity int64 `json:"step_quantity,omitempty" xml:"step_quantity,omitempty"`
	// APP销售时，基于销售单位的起购量
	PurchaseQuantity int64 `json:"purchase_quantity,omitempty" xml:"purchase_quantity,omitempty"`
	// 商品是否适合在app销售，机构商品层级为总控，此字段为总开关，控制所有门店的是否APP可见，商品本身不可售，则app可售必须关闭。App可售关闭，则门店商品维度的app可见不允许打开。1：是? 0：否，（默认否）
	AllowAppSale int64 `json:"allow_app_sale,omitempty" xml:"allow_app_sale,omitempty"`
	// 门店来控制本门店是否在app上让该商品可见，? 1可见? 0不可见
	OnlineSaleFlag int64 `json:"online_sale_flag,omitempty" xml:"online_sale_flag,omitempty"`
	// 禁收时限；收货日期-商品生产日期<限收时限 才允许收货。建议淘鲜达商家填默认值为保质期天数。或按进口商品保质期天数的2/3、非进口商品的1/2填写
	ForbidReceiveDays int64 `json:"forbid_receive_days,omitempty" xml:"forbid_receive_days,omitempty"`
	// 商品生产日期+保质期-当前日期<禁售时限，不能出库销售。建议淘鲜达商家填默认值0。
	ForbidSalesDays int64 `json:"forbid_sales_days,omitempty" xml:"forbid_sales_days,omitempty"`
	// 库存监控报表中，剩余天数少于该天数时需要预警，建议淘鲜达商家填默认值0
	WarnDays int64 `json:"warn_days,omitempty" xml:"warn_days,omitempty"`
	// 商品在门店是否在前场陈列；淘鲜达合作商家用不到， 0：否? 1：是
	FrontDisplayFlag int64 `json:"front_display_flag,omitempty" xml:"front_display_flag,omitempty"`
	// 商品是否是供应商为企业进行定制；淘鲜达合作商家填默认值0， 0 否? 1 是
	FixedFlag int64 `json:"fixed_flag,omitempty" xml:"fixed_flag,omitempty"`
	// 商品的保质期天数（同字段shelf_life），必须为整数，0代表不管理保质期；优先使用本字段
	Period int64 `json:"period,omitempty" xml:"period,omitempty"`
	// 商品类型（同字段sku_type）1:普通商品、2:加工半成品、3:加工成品、4:原材料、5:耗材；优先使用本字段
	ItemTypeNew int64 `json:"item_type_new,omitempty" xml:"item_type_new,omitempty"`
	// 行业属性
	IndustryProps *IndustryPropDo `json:"industry_props,omitempty" xml:"industry_props,omitempty"`
	// forest类目id
	ForestCateId int64 `json:"forest_cate_id,omitempty" xml:"forest_cate_id,omitempty"`
	// 是否易碎品标记，0否1是
	FragileFlag int64 `json:"fragile_flag,omitempty" xml:"fragile_flag,omitempty"`
	// 是否现制现卖，0否1是；比如蛋糕只有下了单才能开始制作
	TemporaryFlag int64 `json:"temporary_flag,omitempty" xml:"temporary_flag,omitempty"`
	// 是否线上的品 1:是  0:否
	IsOnline int64 `json:"is_online,omitempty" xml:"is_online,omitempty"`
	// 商品价格（单位分）和字段sale_price两者只能选其一
	SalePriceUnitCent int64 `json:"sale_price_unit_cent,omitempty" xml:"sale_price_unit_cent,omitempty"`
	// 1、普通商品 2、预售商品
	ItemSaleType int64 `json:"item_sale_type,omitempty" xml:"item_sale_type,omitempty"`
	// 加工时间
	ProcessingTime int64 `json:"processing_time,omitempty" xml:"processing_time,omitempty"`
	// 是否可售，用于查询商品的时候标识商品的可售状态,1可售 0不可售
	SaleFlagForQuery int64 `json:"sale_flag_for_query,omitempty" xml:"sale_flag_for_query,omitempty"`
	// 是否清空会员价, 1清空会员价（操作之后没有会员价），0不清空会员价
	CleanSkuMemberPrice int64 `json:"clean_sku_member_price,omitempty" xml:"clean_sku_member_price,omitempty"`
	// 修改条码策略，默认为策略3。<br/>如商品现有主条码barcode1、非主条码barcode2<br/>1、新增barcodes指定了新的主条码barcode3,会将原主条码barcde1改为非主条码,最终商品有三个条码,即:非主条码barcde1、非主条码barcode2、主条码barcode3<br/> 2、删除barcodes对应条码，除了主条码<br/> 3、替换主条码 则会将现有主条码删除,新增指定主条码,只认barcodes中第一个条码，商品最终结果条码个数为两个条码:非主条码barcode2、主条码barcode3<br/> 4、除主条码外,所有条码替换为指定barcodes,不支持主条码修改。如果barcodes中为barcode3,barcode4,则商品最终条码为:主条码barcode1,非barcode3,barcode4。其中非主条码barcode2被删除
	BarcodeUpdateType int64 `json:"barcode_update_type,omitempty" xml:"barcode_update_type,omitempty"`
}

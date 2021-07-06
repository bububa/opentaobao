package fenxiao

// CnskuFeatureDto 结构体
type CnskuFeatureDto struct {
	// SN管理模式
	SnMode string `json:"sn_mode,omitempty" xml:"sn_mode,omitempty"`
	// 产地
	OriginAddress string `json:"origin_address,omitempty" xml:"origin_address,omitempty"`
	// 服务产品标
	GrayFlag string `json:"gray_flag,omitempty" xml:"gray_flag,omitempty"`
	// 商品类目
	WhcCategory string `json:"whc_category,omitempty" xml:"whc_category,omitempty"`
	// 计量单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 尺寸
	Size string `json:"size,omitempty" xml:"size,omitempty"`
	// 包材分组
	MaterialGroup string `json:"material_group,omitempty" xml:"material_group,omitempty"`
	// 商家商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 对应货值 单位元 小数点后保留2位
	GoodsValue string `json:"goods_value,omitempty" xml:"goods_value,omitempty"`
	// 货号
	GoodsNo string `json:"goods_no,omitempty" xml:"goods_no,omitempty"`
	// 包材分类
	MaterialClass string `json:"material_class,omitempty" xml:"material_class,omitempty"`
	// 规格
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 规格类型，20：大件；10：中小件
	SpecificationType string `json:"specification_type,omitempty" xml:"specification_type,omitempty"`
	// 包装方式：1 原箱发货
	PackagingScheme string `json:"packaging_scheme,omitempty" xml:"packaging_scheme,omitempty"`
	// 产品编码
	CnProductCode string `json:"cn_product_code,omitempty" xml:"cn_product_code,omitempty"`
	// 批准文号
	ApprovalNumber string `json:"approval_number,omitempty" xml:"approval_number,omitempty"`
	// 容积
	CnCubage string `json:"cn_cubage,omitempty" xml:"cn_cubage,omitempty"`
	// 商品名称
	WhcName string `json:"whc_name,omitempty" xml:"whc_name,omitempty"`
	// 零售价
	SkuPrice int64 `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// 有效期天数
	Lifecycle int64 `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	// 净重
	NetWeight int64 `json:"net_weight,omitempty" xml:"net_weight,omitempty"`
	// 保质期禁收天数
	RejectLifecycle int64 `json:"reject_lifecycle,omitempty" xml:"reject_lifecycle,omitempty"`
	// 保质期禁售天数
	LockupLifecycle int64 `json:"lockup_lifecycle,omitempty" xml:"lockup_lifecycle,omitempty"`
	// 保质期临期天数
	AdventLifecycle int64 `json:"advent_lifecycle,omitempty" xml:"advent_lifecycle,omitempty"`
	// 箱规
	Pcs int64 `json:"pcs,omitempty" xml:"pcs,omitempty"`
	// 成本价
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
	// 是否启用序列号（sn）管理
	IsSnMgt bool `json:"is_sn_mgt,omitempty" xml:"is_sn_mgt,omitempty"`
	// 是否贵品
	IsPrecious bool `json:"is_precious,omitempty" xml:"is_precious,omitempty"`
	// 是否规范运输单元
	IsStandardCarton bool `json:"is_standard_carton,omitempty" xml:"is_standard_carton,omitempty"`
	// sn出库管理
	WhcSnOutMode bool `json:"whc_sn_out_mode,omitempty" xml:"whc_sn_out_mode,omitempty"`
	// 是否危险品
	IsDanger bool `json:"is_danger,omitempty" xml:"is_danger,omitempty"`
	// sn入库管理
	WhcSnInMode bool `json:"whc_sn_in_mode,omitempty" xml:"whc_sn_in_mode,omitempty"`
	// 是否统一全仓数据-销售单元
	IsUnifiedAllWh bool `json:"is_unified_all_wh,omitempty" xml:"is_unified_all_wh,omitempty"`
	// 是否统一全仓数据-运输单元
	IsUnifiedAllWhCarton bool `json:"is_unified_all_wh_carton,omitempty" xml:"is_unified_all_wh_carton,omitempty"`
	// 是否需要测量图片-运输单元
	NeedMeasureImageCarton bool `json:"need_measure_image_carton,omitempty" xml:"need_measure_image_carton,omitempty"`
	// 是否Po管理
	IsPoMgt bool `json:"is_po_mgt,omitempty" xml:"is_po_mgt,omitempty"`
	// 是否启用标识
	UseYn bool `json:"use_yn,omitempty" xml:"use_yn,omitempty"`
	// 是否易碎品
	IsHygroscopic bool `json:"is_hygroscopic,omitempty" xml:"is_hygroscopic,omitempty"`
	// 是否规范销售单元
	IsStandard bool `json:"is_standard,omitempty" xml:"is_standard,omitempty"`
	// 是否启用批次管理标识
	IsBatchMgt bool `json:"is_batch_mgt,omitempty" xml:"is_batch_mgt,omitempty"`
	// 是否效期管理
	IsShelflife bool `json:"is_shelflife,omitempty" xml:"is_shelflife,omitempty"`
	// 是否需要测量图片-销售单元
	NeedMeasureImage bool `json:"need_measure_image,omitempty" xml:"need_measure_image,omitempty"`
}

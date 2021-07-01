package lstwarehouse

// Content 结构体
type Content struct {
	// 是否禁止补货
	IsNoRm string `json:"is_no_rm,omitempty" xml:"is_no_rm,omitempty"`
	// 报价状态
	ProductQuoteStatus string `json:"product_quote_status,omitempty" xml:"product_quote_status,omitempty"`
	// 连续缺货天数
	OssDaysStd001 int64 `json:"oss_days_std001,omitempty" xml:"oss_days_std001,omitempty"`
	// 单品编码
	CspuId string `json:"cspu_id,omitempty" xml:"cspu_id,omitempty"`
	// 是否在售
	IsOnline string `json:"is_online,omitempty" xml:"is_online,omitempty"`
	// 销售单位
	SaleUnitDesc string `json:"sale_unit_desc,omitempty" xml:"sale_unit_desc,omitempty"`
	// 当前商品分层
	ItmLevel string `json:"itm_level,omitempty" xml:"itm_level,omitempty"`
	// 补货配送方式
	RmShipType string `json:"rm_ship_type,omitempty" xml:"rm_ship_type,omitempty"`
	// 出清状态
	LiquidationStatus string `json:"liquidation_status,omitempty" xml:"liquidation_status,omitempty"`
	// 仓库名称
	WhName string `json:"wh_name,omitempty" xml:"wh_name,omitempty"`
	// 单品名称
	CspuName string `json:"cspu_name,omitempty" xml:"cspu_name,omitempty"`
	// 仓库类型名称
	SubWhTypeName string `json:"sub_wh_type_name,omitempty" xml:"sub_wh_type_name,omitempty"`
	// 前端商品列表
	ItemsId string `json:"items_id,omitempty" xml:"items_id,omitempty"`
	// 公司名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 数据日期
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 在途数量
	AllocOpQty string `json:"alloc_op_qty,omitempty" xml:"alloc_op_qty,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 条码
	BarCodeList string `json:"bar_code_list,omitempty" xml:"bar_code_list,omitempty"`
	// 是否OB
	IsOb string `json:"is_ob,omitempty" xml:"is_ob,omitempty"`
	// 货品id
	BackOfferId string `json:"back_offer_id,omitempty" xml:"back_offer_id,omitempty"`
	// 箱规数
	CartonPcs string `json:"carton_pcs,omitempty" xml:"carton_pcs,omitempty"`
	// 可售库存
	DeportItmQtyStd008 int64 `json:"deport_itm_qty_std008,omitempty" xml:"deport_itm_qty_std008,omitempty"`
	// 残次品在仓商品件数
	DeportItmQtyStd005 int64 `json:"deport_itm_qty_std005,omitempty" xml:"deport_itm_qty_std005,omitempty"`
	// 是否缺货
	IsOos string `json:"is_oos,omitempty" xml:"is_oos,omitempty"`
	// 首次入库时间
	FstInboundTime string `json:"fst_inbound_time,omitempty" xml:"fst_inbound_time,omitempty"`
	// 最小起批量
	MinBuyQty string `json:"min_buy_qty,omitempty" xml:"min_buy_qty,omitempty"`
	// 最小购买倍数
	MinBuyMultiple string `json:"min_buy_multiple,omitempty" xml:"min_buy_multiple,omitempty"`
	// 是否滞销
	IsDull string `json:"is_dull,omitempty" xml:"is_dull,omitempty"`
	// 二级类目名称
	CateLv2Name string `json:"cate_lv2_name,omitempty" xml:"cate_lv2_name,omitempty"`
	// 上级仓名称
	ParentWhName string `json:"parent_wh_name,omitempty" xml:"parent_wh_name,omitempty"`
	// 大区名称
	BigAreaName string `json:"big_area_name,omitempty" xml:"big_area_name,omitempty"`
	// 首次入库天数
	FstInboundDays int64 `json:"fst_inbound_days,omitempty" xml:"fst_inbound_days,omitempty"`
	// 一级类目名称
	CateLv1Name string `json:"cate_lv1_name,omitempty" xml:"cate_lv1_name,omitempty"`
	// 三级类目名称
	CateLv3Name string `json:"cate_lv3_name,omitempty" xml:"cate_lv3_name,omitempty"`
	// 滞销库存临界天数
	UnsaleInstockDays string `json:"unsale_instock_days,omitempty" xml:"unsale_instock_days,omitempty"`
	// 品牌名称列表
	BrandNameList []string `json:"brand_name_list,omitempty" xml:"brand_name_list>string,omitempty"`
	// 供应商memberId
	SupplierMemberId string `json:"supplier_member_id,omitempty" xml:"supplier_member_id,omitempty"`
	// 成功/失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 商品Id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 错误code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 实仓：cainiao，商家仓：customer
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 仓库code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

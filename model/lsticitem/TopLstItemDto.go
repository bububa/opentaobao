package lsticitem

// TopLstItemDto 结构体
type TopLstItemDto struct {
	// 条码列表
	BarCodeList []string `json:"bar_code_list,omitempty" xml:"bar_code_list>string,omitempty"`
	// 图片列表
	ImgList []string `json:"img_list,omitempty" xml:"img_list>string,omitempty"`
	// 库存集合
	AvailableStockList []Stock `json:"available_stock_list,omitempty" xml:"available_stock_list>stock,omitempty"`
	// 商品类型（售卖属性） normal：通常品 gift：赠品（比如买a送a，买a送b，送的商品就是赠品
	ItemType string `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 品牌名
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 单品类型（单品自身属性） normal：普通品（可以售卖） combine：组合品（作为组套商品进行售卖） zengpin：赠品（非标品，只能作为赠品，不能上架售卖） mixed：混箱品
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 单品编码
	CspuId string `json:"cspu_id,omitempty" xml:"cspu_id,omitempty"`
	// 商家货号/商家商品编码
	CargoNumber string `json:"cargo_number,omitempty" xml:"cargo_number,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 最新上架时间
	OnSaleTime string `json:"on_sale_time,omitempty" xml:"on_sale_time,omitempty"`
	// 商品创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// cancel：撤销 no_pass_audit：审核不通过 on_sale：上架销售中 un_sale_waiting_audit：下架待审核 waiting_audit：待审核 waiting_choose：审核通过，待上架
	ItemStatus string `json:"item_status,omitempty" xml:"item_status,omitempty"`
	// 二级类目名字
	SecondCategoryName string `json:"second_category_name,omitempty" xml:"second_category_name,omitempty"`
	// 长标题
	FullItemTitle string `json:"full_item_title,omitempty" xml:"full_item_title,omitempty"`
	// 仓库类型supplier：虚仓；cainiao：实仓
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 二级类目Id
	SecondCategoryId string `json:"second_category_id,omitempty" xml:"second_category_id,omitempty"`
	// 短标题
	ShortItemTitle string `json:"short_item_title,omitempty" xml:"short_item_title,omitempty"`
	// 商品Id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

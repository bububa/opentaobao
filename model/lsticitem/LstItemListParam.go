package lsticitem

// LstItemListParam 结构体
type LstItemListParam struct {
	// 商品创建时间-开始
	GmtCreateStart string `json:"gmt_create_start,omitempty" xml:"gmt_create_start,omitempty"`
	// 最近上架时间-开始
	GmtOnSaleStart string `json:"gmt_on_sale_start,omitempty" xml:"gmt_on_sale_start,omitempty"`
	// 商品修改时间-结束
	GmtModifiedEnd string `json:"gmt_modified_end,omitempty" xml:"gmt_modified_end,omitempty"`
	// 商品修改时间-开始
	GmtModifiedStart string `json:"gmt_modified_start,omitempty" xml:"gmt_modified_start,omitempty"`
	// 商品类型列表 normal：通常商品 gift：赠品
	ItemTypeList []string `json:"item_type_list,omitempty" xml:"item_type_list>string,omitempty"`
	// 最近上架时间-结束
	GmtOnSaleEnd string `json:"gmt_on_sale_end,omitempty" xml:"gmt_on_sale_end,omitempty"`
	// 商品创建时间-结束
	GmtCreateEnd string `json:"gmt_create_end,omitempty" xml:"gmt_create_end,omitempty"`
	// 货号列表
	CargoNumberList []string `json:"cargo_number_list,omitempty" xml:"cargo_number_list>string,omitempty"`
	// 商品状态列表： cancel：撤销 no_pass_audit：审核不通过 on_sale：上架销售中 un_sale_waiting_audit：下架待审核 waiting_audit：待审核 waiting_choose：审核通过，待上架
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 返回列表结果集每页条数。取值范围:大于零的整数;最大值：50;
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 商品ID列表
	ItemIdList []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
	// 页码。取值范围:大于零的整数;默认值为1，即返回第一页数据。
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

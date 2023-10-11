package simba

// TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult 结构体
type TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	StdCategoryVOTopBulkData *TopBulkData `json:"std_category_v_o_top_bulk_data,omitempty" xml:"std_category_v_o_top_bulk_data,omitempty"`
}

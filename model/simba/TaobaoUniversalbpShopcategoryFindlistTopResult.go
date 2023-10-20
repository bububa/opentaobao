package simba

// TaobaouniversalbpshopcategoryfindlistTopResult 结构体
type TaobaouniversalbpshopcategoryfindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	ShopCategoryVOTopBulkData *TopBulkData `json:"shop_category_v_o_top_bulk_data,omitempty" xml:"shop_category_v_o_top_bulk_data,omitempty"`
}

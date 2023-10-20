package simba

// TaobaouniversalbpmemberfindbrandinfolistTopResult 结构体
type TaobaouniversalbpmemberfindbrandinfolistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	BrandInfoVOTopBulkData *TopBulkData `json:"brand_info_v_o_top_bulk_data,omitempty" xml:"brand_info_v_o_top_bulk_data,omitempty"`
}

package tbk

// TaobaotbkdgnewuserordergetData 结构体
type TaobaotbkdgnewuserordergetData struct {
	// resultList
	Results []TaobaotbkdgnewuserordergetMapData `json:"results,omitempty" xml:"results>taobaotbkdgnewuserorderget_map_data,omitempty"`
	// 页码
	Pageno int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小
	Pagesize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否有下一页
	Hasnext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

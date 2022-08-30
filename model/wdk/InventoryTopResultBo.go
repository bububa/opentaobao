package wdk

// InventoryTopResultBo 结构体
type InventoryTopResultBo struct {
	// inventoryInfoList
	InventoryInfoList []Inventoryinfolist `json:"inventory_info_list,omitempty" xml:"inventory_info_list>inventoryinfolist,omitempty"`
	// queryPageNum
	QueryPageNum int64 `json:"query_page_num,omitempty" xml:"query_page_num,omitempty"`
	// queryPageSize
	QueryPageSize int64 `json:"query_page_size,omitempty" xml:"query_page_size,omitempty"`
	// totalPageNum
	TotalPageNum int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
}

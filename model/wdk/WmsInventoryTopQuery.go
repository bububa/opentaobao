package wdk

// WmsInventoryTopQuery 结构体
type WmsInventoryTopQuery struct {
	// 仓库编号
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 查询页数
	QueryPageNum int64 `json:"query_page_num,omitempty" xml:"query_page_num,omitempty"`
}

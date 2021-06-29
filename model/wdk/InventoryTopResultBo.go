package wdk

// InventoryTopResultBo 
type InventoryTopResultBo struct {
    // 库存信息列表
    InventoryInfoList   []Inventoryinfolist `json:"inventory_info_list,omitempty" xml:"inventory_info_list>inventoryinfolist,omitempty"`
    // 总页数
    TotalPageNum   int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
    // 分页大小
    QueryPageSize   int64 `json:"query_page_size,omitempty" xml:"query_page_size,omitempty"`
    // 查询页码
    QueryPageNum   int64 `json:"query_page_num,omitempty" xml:"query_page_num,omitempty"`
    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 当前请求的最大ID
    CurrentMaxId   int64 `json:"current_max_id,omitempty" xml:"current_max_id,omitempty"`
}

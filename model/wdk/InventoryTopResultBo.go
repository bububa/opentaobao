package wdk

// InventoryTopResultBo 
/* model for simplify = false
type InventoryTopResultBo struct {

    // 库存信息列表
    
    InventoryInfoList  struct {
        Inventoryinfolist  []Inventoryinfolist `json:"inventoryinfolist,omitempty"`
    } `json:"inventory_info_list,omitempty"`
    

    // 总页数
    
    TotalPageNum   int64 `json:"total_page_num,omitempty"`
    

    // 分页大小
    
    QueryPageSize   int64 `json:"query_page_size,omitempty"`
    

    // 查询页码
    
    QueryPageNum   int64 `json:"query_page_num,omitempty"`
    

    // 仓库编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // 当前请求的最大ID
    
    CurrentMaxId   int64 `json:"current_max_id,omitempty"`
    

}
*/

// InventoryTopResultBo 
type InventoryTopResultBo struct {

    // 库存信息列表
    InventoryInfoList   []Inventoryinfolist `json:"inventory_info_list,omitempty"`

    // 总页数
    TotalPageNum   int64 `json:"total_page_num,omitempty"`

    // 分页大小
    QueryPageSize   int64 `json:"query_page_size,omitempty"`

    // 查询页码
    QueryPageNum   int64 `json:"query_page_num,omitempty"`

    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 当前请求的最大ID
    CurrentMaxId   int64 `json:"current_max_id,omitempty"`

}

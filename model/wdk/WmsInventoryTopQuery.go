package wdk

// WmsInventoryTopQuery 
/* model for simplify = false
type WmsInventoryTopQuery struct {

    // 仓库编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // 查询页数
    
    QueryPageNum   int64 `json:"query_page_num,omitempty"`
    

    // 上一次调用的返回current_max_id，第一次调用为0
    
    LastMaxId   int64 `json:"last_max_id,omitempty"`
    

}
*/

// WmsInventoryTopQuery 
type WmsInventoryTopQuery struct {

    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 查询页数
    QueryPageNum   int64 `json:"query_page_num,omitempty"`

    // 上一次调用的返回current_max_id，第一次调用为0
    LastMaxId   int64 `json:"last_max_id,omitempty"`

}

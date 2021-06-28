package wdk

// OrderListSyncPagedResult 
/* model for simplify = false
type OrderListSyncPagedResult struct {

    // returnMsg
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // returnCode
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 返回订单总数量
    
    TotalNumber   int64 `json:"total_number,omitempty"`
    

    // orders
    
    Orders  struct {
        OrderSyncDto  []OrderSyncDto `json:"order_sync_dto,omitempty"`
    } `json:"orders,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // 返回下一查询页的序号。如果返回值是-1，则无下一页。数据拉取完成。
    
    NextIndex   int64 `json:"next_index,omitempty"`
    

}
*/

// OrderListSyncPagedResult 
type OrderListSyncPagedResult struct {

    // returnMsg
    ReturnMsg   string `json:"return_msg,omitempty"`

    // returnCode
    ReturnCode   string `json:"return_code,omitempty"`

    // 返回订单总数量
    TotalNumber   int64 `json:"total_number,omitempty"`

    // orders
    Orders   []OrderSyncDto `json:"orders,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 返回下一查询页的序号。如果返回值是-1，则无下一页。数据拉取完成。
    NextIndex   int64 `json:"next_index,omitempty"`

}

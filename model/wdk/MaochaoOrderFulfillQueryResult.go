package wdk

// MaochaoOrderFulfillQueryResult 
/* model for simplify = false
type MaochaoOrderFulfillQueryResult struct {

    // 是否调用成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返回码
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 返回码说明
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 履约单列表
    
    FulfillOrderList  struct {
        MaochaoWdkOrderFulfillDto  []MaochaoWdkOrderFulfillDto `json:"maochao_wdk_order_fulfill_dto,omitempty"`
    } `json:"fulfill_order_list,omitempty"`
    

}
*/

// MaochaoOrderFulfillQueryResult 
type MaochaoOrderFulfillQueryResult struct {

    // 是否调用成功
    Success   bool `json:"success,omitempty"`

    // 返回码
    ReturnCode   string `json:"return_code,omitempty"`

    // 返回码说明
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 履约单列表
    FulfillOrderList   []MaochaoWdkOrderFulfillDto `json:"fulfill_order_list,omitempty"`

}

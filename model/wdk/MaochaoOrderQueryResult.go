package wdk

// MaochaoOrderQueryResult 
/* model for simplify = false
type MaochaoOrderQueryResult struct {

    // 是否调用成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返回码
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 返回码说明
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 子订单列表
    
    SubOrderList  struct {
        MaochaoWdkOrderDto  []MaochaoWdkOrderDto `json:"maochao_wdk_order_dto,omitempty"`
    } `json:"sub_order_list,omitempty"`
    

}
*/

// MaochaoOrderQueryResult 
type MaochaoOrderQueryResult struct {

    // 是否调用成功
    Success   bool `json:"success,omitempty"`

    // 返回码
    ReturnCode   string `json:"return_code,omitempty"`

    // 返回码说明
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 子订单列表
    SubOrderList   []MaochaoWdkOrderDto `json:"sub_order_list,omitempty"`

}

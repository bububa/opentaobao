package wms

// TaobaoWlbWmsItemCombinationGetResult 
/* model for simplify = false
type TaobaoWlbWmsItemCombinationGetResult struct {

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 子货品列表
    
    SubItemList  struct {
        SubItemList  []SubItemList `json:"sub_item_list,omitempty"`
    } `json:"sub_item_list,omitempty"`
    

}
*/

// TaobaoWlbWmsItemCombinationGetResult 
type TaobaoWlbWmsItemCombinationGetResult struct {

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 子货品列表
    SubItemList   []SubItemList `json:"sub_item_list,omitempty"`

}

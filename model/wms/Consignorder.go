package wms

// Consignorder 
/* model for simplify = false
type Consignorder struct {

    // 仓库物流订单信息列表
    
    ConsignOrderItemList  struct {
        Consignorderitemlist  []Consignorderitemlist `json:"consignorderitemlist,omitempty"`
    } `json:"consign_order_item_list,omitempty"`
    

    // 仓库订单编码
    
    StoreOrderCode   string `json:"store_order_code,omitempty"`
    

    // 错误编码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 配送编码
    
    TmsCode   string `json:"tms_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 仓库编码
    
    StoreCode   string `json:"store_code,omitempty"`
    

}
*/

// Consignorder 
type Consignorder struct {

    // 仓库物流订单信息列表
    ConsignOrderItemList   []Consignorderitemlist `json:"consign_order_item_list,omitempty"`

    // 仓库订单编码
    StoreOrderCode   string `json:"store_order_code,omitempty"`

    // 错误编码
    ErrorCode   string `json:"error_code,omitempty"`

    // 配送编码
    TmsCode   string `json:"tms_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 仓库编码
    StoreCode   string `json:"store_code,omitempty"`

}

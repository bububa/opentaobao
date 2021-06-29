package wms

// Consignorder 
type Consignorder struct {
    // 仓库物流订单信息列表
    ConsignOrderItemList   []Consignorderitemlist `json:"consign_order_item_list,omitempty" xml:"consign_order_item_list>consignorderitemlist,omitempty"`
    // 仓库订单编码
    StoreOrderCode   string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`
    // 错误编码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 配送编码
    TmsCode   string `json:"tms_code,omitempty" xml:"tms_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 仓库编码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}

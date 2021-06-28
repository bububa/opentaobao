package waybill

// WaybillApplyNewRequest 
type WaybillApplyNewRequest struct {

    // 物流服务商编码
    
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    

    // 收\发货地址
    
    ShippingAddress   *WaybillAddress `json:"shipping_address,omitempty" xml:"shipping_address,omitempty"`
    

    // 订单数据
    
    TradeOrderInfoCols   []TradeOrderInfo `json:"trade_order_info_cols,omitempty" xml:"trade_order_info_cols,omitempty"`
    

}

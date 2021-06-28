package waybill

// WaybillApplyNewRequest 
/* model for simplify = false
type WaybillApplyNewRequest struct {

    // 物流服务商编码
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 收\发货地址
    
    ShippingAddress  *struct {
        WaybillAddress  *WaybillAddress `json:"waybill_address,omitempty"`
    } `json:"shipping_address,omitempty"`
    

    // 订单数据
    
    TradeOrderInfoCols  struct {
        TradeOrderInfo  []TradeOrderInfo `json:"trade_order_info,omitempty"`
    } `json:"trade_order_info_cols,omitempty"`
    

}
*/

// WaybillApplyNewRequest 
type WaybillApplyNewRequest struct {

    // 物流服务商编码
    CpCode   string `json:"cp_code,omitempty"`

    // 收\发货地址
    ShippingAddress   *WaybillAddress `json:"shipping_address,omitempty"`

    // 订单数据
    TradeOrderInfoCols   []TradeOrderInfo `json:"trade_order_info_cols,omitempty"`

}

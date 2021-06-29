package lstlogistics2

// LstOfflineOrderCancalParam 
type LstOfflineOrderCancalParam struct {
    // 买家手机号
    BuyerMobile   string `json:"buyer_mobile,omitempty" xml:"buyer_mobile,omitempty"`
    // 外部主订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

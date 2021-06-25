package wdk

// WdkOpenPushSoPoRequest 
type WdkOpenPushSoPoRequest struct {

    // 淘系子订单列表，一次最多20个
    SubTradeOrderNoList   []String `json:"sub_trade_order_no_list,omitempty"`

    // 销售=10，销退=20
    SalesType   int64 `json:"sales_type,omitempty"`

}

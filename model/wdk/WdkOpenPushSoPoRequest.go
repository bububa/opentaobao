package wdk

// WdkOpenPushSoPoRequest 
type WdkOpenPushSoPoRequest struct {

    // 淘系子订单列表，一次最多20个
    
    SubTradeOrderNoList   []string `json:"sub_trade_order_no_list,omitempty" xml:"sub_trade_order_no_list>string,omitempty"`
    

    // 销售=10，销退=20
    
    SalesType   int64 `json:"sales_type,omitempty" xml:"sales_type,omitempty"`
    

}

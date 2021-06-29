package omniorder

// QueryOmniOrderRequest 
type QueryOmniOrderRequest struct {

    // 订单创建结束时间，秒时间戳
    
    EndCreated   int64 `json:"end_created,omitempty" xml:"end_created,omitempty"`
    

    // 页码
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 品牌官旗的sellerId
    
    BrandSellerId   int64 `json:"brand_seller_id,omitempty" xml:"brand_seller_id,omitempty"`
    

    // 页大小，最大不超过100
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 订单创建起始时间，秒时间戳
    
    StartCreated   int64 `json:"start_created,omitempty" xml:"start_created,omitempty"`
    

    // 订单状态，可选值：WAIT_BUYER_PAY(等待买家付款), WAIT_SELLER_SEND_GOODS(等待卖家发货), SELLER_CONSIGNED_PART(卖家部分发货), WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货]), TRADE_BUYER_SIGNED(买家已签收（货到付款专用）), TRADE_FINISHED(交易成功), TRADE_CLOSED(交易关闭), TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭), TRADE_NO_CREATE_PAY(没有创建外部交易（支付宝交易）)
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

}

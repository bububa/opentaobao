package logistic

// DeliveryEventDto 
type DeliveryEventDto struct {

    // 淘宝订单ID
    
    TaobaoOrderId   int64 `json:"taobao_order_id,omitempty" xml:"taobao_order_id,omitempty"`
    

    // 配送商名称
    
    DistributorName   string `json:"distributor_name,omitempty" xml:"distributor_name,omitempty"`
    

    // 揽件员姓名
    
    ShopperName   string `json:"shopper_name,omitempty" xml:"shopper_name,omitempty"`
    

    // 快递单号
    
    DeliveryNo   string `json:"delivery_no,omitempty" xml:"delivery_no,omitempty"`
    

    // 事件发生时间
    
    EventTime   string `json:"event_time,omitempty" xml:"event_time,omitempty"`
    

    // 事件类型：呼叫物流=4; 揽件=7; 送达=8; 配送取消=9
    
    EventType   int64 `json:"event_type,omitempty" xml:"event_type,omitempty"`
    

    // 取消原因（请传以下类型）：CHAIN_STORE_NOT_EXIST("门店不存在"),     CHAIN_STORE_NOT_VERIFIED("门店未审核"),     MERCHANT_ARREARAGE("商户欠费"),     MERCHANT_NOT_EXISTS("商户不存在"),     REQUIRE_RECEIVE_TIME_INVALID("预计送达时间不正确"),     MERCHANT_PUSH_TOO_LATE_ERROR("请求配送过晚"),      NO_CONTRACTED_ERROR("商家未签约"),     MERCHANT_CALL_LATE_ERROR("呼叫配送晚"),     REJECT_ORDER("系统拒单"),     OUT_OF_DELIVERY_RANGE("超出配送范围"),     OVER_DELIVERY_AREA_ERROR("超出配送范围"),      CONSUMER_LOCATION_TOO_FAR_ERROR("超出服务范围"),     GOODS_WEIGHT_EXCEPTION_ERROR("货物重量异常"),     PRODUCT_OVERWEIGHT("无责取消-商品过重或体积过大"),     DELIVERY_OUT_OF_SERVICE("超出服务范围"),     MERCHANT_CANCEL("无责取消-商户申请取消"),      MERCHANT_HAS_NOT_OPENED("无责取消-商户未开店"),     MERCHANT_INTERRUPT_DELIVERY_ERROR("无责取消-商户未开店"),     MERCHANT_CANCLE_BY_CUSTOMER_SERVICE("点我达运力独有-商家要求客服取消"),     MERCHANT_DONT_GIVE_FOOD("无责取消-商户不给餐"),     MERCHANT_FOOD_ERROR("无责取消-商户出货问题"),      MERCHANT_OUT_OF_STOCK("无责取消-商户缺货"),     USER_CANCEL("用户发起取消"),     USER_CANCEL_BY_CODE("用户取消码取消"),     USER_CANCEL_COMPENSATORY("用户有偿取消"),     USER_ADDRESS_ERROR("异常报备-顾客定位有误"),      USER_NOT_ANSWER_ERROR("异常报备-联系不上顾客"),     USER_RETURN_ORDER_ERROR("无责取消-顾客拒收订单"),     CARRIER_CANCEL("配送商取消-挂单时长内无人接单取消"),     CARRIER_DONT_GIVE_FOOD_CANCEL("骑手抢单后规定时间内未到店/未取餐取消"),     CARRIER_GRAB_ERROR("运力主动取消-抢错了"),      CARRIER_LONG_TIME_CANCEL("挂单时长内无人接单取消"),     CARRIER_NOT_ENOUGH_TIME("运力主动取消-时间来不及"),     CARRIER_OTHER_ERROR("运力主动取消-运力其他原因"),     CARRIER_REMARK_EXCEPTION_ERROR("骑手标记异常"),     DELIVERY_TIMEOUT("配送超时，系统标记异常（超时关单）"),      FOR_UPDATE_TIP("加小费取消"),     SYSTEM_CANCEL("系统取消"),     SYSTEM_MARKED_ERROR("系统自动标记异常--订单超过3小时未送达"),     CARRIER_NOT_ON_TIME_ERROR("配送商未按时配送");
    
    CancelReason   string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
    

    // 揽件员电话
    
    ShopperPhone   string `json:"shopper_phone,omitempty" xml:"shopper_phone,omitempty"`
    

}

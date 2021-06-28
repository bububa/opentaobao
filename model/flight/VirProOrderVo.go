package flight

// VirProOrderVo 
/* model for simplify = false
type VirProOrderVo struct {

    // 预订时间，辅营订单创建时间
    
    BookTime   string `json:"book_time,omitempty"`
    

    // 机票订单号
    
    FlightOrderId   int64 `json:"flight_order_id,omitempty"`
    

    // 乘机人购买辅营产品详情
    
    PassengerAuxVos  struct {
        PassengerAuxVo  []PassengerAuxVo `json:"passenger_aux_vo,omitempty"`
    } `json:"passenger_aux_vos,omitempty"`
    

    // 支付宝流水号，存在多个辅营订单对应一个支付宝流水号的情况
    
    PayNo   string `json:"pay_no,omitempty"`
    

    // 辅营订单金额
    
    PayPrice   int64 `json:"pay_price,omitempty"`
    

    // 支付时间，订单为支付成功或出货成功时返回
    
    PayTime   string `json:"pay_time,omitempty"`
    

    // 辅营订单号
    
    OrderId   int64 `json:"order_id,omitempty"`
    

    // 辅营订单状态，1-	待支付 2-	支付成功 3-	辅营出货成功 4-	订单取消
    
    OrderStatus   int64 `json:"order_status,omitempty"`
    

}
*/

// VirProOrderVo 
type VirProOrderVo struct {

    // 预订时间，辅营订单创建时间
    BookTime   string `json:"book_time,omitempty"`

    // 机票订单号
    FlightOrderId   int64 `json:"flight_order_id,omitempty"`

    // 乘机人购买辅营产品详情
    PassengerAuxVos   []PassengerAuxVo `json:"passenger_aux_vos,omitempty"`

    // 支付宝流水号，存在多个辅营订单对应一个支付宝流水号的情况
    PayNo   string `json:"pay_no,omitempty"`

    // 辅营订单金额
    PayPrice   int64 `json:"pay_price,omitempty"`

    // 支付时间，订单为支付成功或出货成功时返回
    PayTime   string `json:"pay_time,omitempty"`

    // 辅营订单号
    OrderId   int64 `json:"order_id,omitempty"`

    // 辅营订单状态，1-	待支付 2-	支付成功 3-	辅营出货成功 4-	订单取消
    OrderStatus   int64 `json:"order_status,omitempty"`

}

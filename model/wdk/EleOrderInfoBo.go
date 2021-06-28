package wdk

// EleOrderInfoBo 
/* model for simplify = false
type EleOrderInfoBo struct {

    // 订单费用明细
    
    OrderDetailFee  *struct {
        OrderDetailFee  *OrderDetailFee `json:"order_detail_fee,omitempty"`
    } `json:"order_detail_fee,omitempty"`
    

    // 损失承担方
    
    ResponsibleParty   string `json:"responsible_party,omitempty"`
    

    // 金额
    
    Amount   string `json:"amount,omitempty"`
    

    // 账务时间
    
    TradeCreateTime   string `json:"trade_create_time,omitempty"`
    

    // 下单时间
    
    OrderCreateTime   string `json:"order_create_time,omitempty"`
    

    // 实际付款主体
    
    PayEntity   string `json:"pay_entity,omitempty"`
    

    // 饿了么订单id
    
    EleOrderId   string `json:"ele_order_id,omitempty"`
    

    // 订单id
    
    OrderId   string `json:"order_id,omitempty"`
    

    // 订单来源
    
    OrderFrom   string `json:"order_from,omitempty"`
    

    // 订单序号
    
    OrderIndex   string `json:"order_index,omitempty"`
    

}
*/

// EleOrderInfoBo 
type EleOrderInfoBo struct {

    // 订单费用明细
    OrderDetailFee   *OrderDetailFee `json:"order_detail_fee,omitempty"`

    // 损失承担方
    ResponsibleParty   string `json:"responsible_party,omitempty"`

    // 金额
    Amount   string `json:"amount,omitempty"`

    // 账务时间
    TradeCreateTime   string `json:"trade_create_time,omitempty"`

    // 下单时间
    OrderCreateTime   string `json:"order_create_time,omitempty"`

    // 实际付款主体
    PayEntity   string `json:"pay_entity,omitempty"`

    // 饿了么订单id
    EleOrderId   string `json:"ele_order_id,omitempty"`

    // 订单id
    OrderId   string `json:"order_id,omitempty"`

    // 订单来源
    OrderFrom   string `json:"order_from,omitempty"`

    // 订单序号
    OrderIndex   string `json:"order_index,omitempty"`

}

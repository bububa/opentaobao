package wdk

// EleBillBo 
type EleBillBo struct {

    // 账单日期，时间戳
    Date   string `json:"date,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

    // 订单列表
    OrderList   []EleOrderInfoBo `json:"order_list,omitempty"`

    // 订单费用明细
    OrderDetailFee   *OrderDetailFee `json:"order_detail_fee,omitempty"`

    // 应付金额
    ExpendFee   string `json:"expend_fee,omitempty"`

    // 单量
    OrderCount   string `json:"order_count,omitempty"`

    // 未结算金额，单位:分
    PayFee   string `json:"pay_fee,omitempty"`

    // 实际付款主体
    PayEntity   string `json:"pay_entity,omitempty"`

}

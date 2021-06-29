package omniorder

// CommissionResultDto 
type CommissionResultDto struct {
    // 分佣id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 订单类型
    BizOrderType   int64 `json:"biz_order_type,omitempty" xml:"biz_order_type,omitempty"`
    // 订单金额，单位分
    BizOrderMoney   int64 `json:"biz_order_money,omitempty" xml:"biz_order_money,omitempty"`
    // 实付金额，单位分
    OrderPayMoney   int64 `json:"order_pay_money,omitempty" xml:"order_pay_money,omitempty"`
    // 创建时间
    OrderCreateTime   string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
    // 支付时间
    OrderPayTime   string `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
    // 结束时间
    OrderEndTime   string `json:"order_end_time,omitempty" xml:"order_end_time,omitempty"`
    // 佣金，单位分
    CommissionMoney   int64 `json:"commission_money,omitempty" xml:"commission_money,omitempty"`
    // 分佣导购id
    CommissionEmployeeId   int64 `json:"commission_employee_id,omitempty" xml:"commission_employee_id,omitempty"`
    // 分佣导购类型：1-招募导购,2-当前导购,3-首单导购,4-外部订单导购，5-好友，6-接单导购
    CommissionEmployeeType   int64 `json:"commission_employee_type,omitempty" xml:"commission_employee_type,omitempty"`
    // 分佣导购名称
    CommissionEmployeeName   string `json:"commission_employee_name,omitempty" xml:"commission_employee_name,omitempty"`
    // 分佣导购所属门店id
    EmployeeStoreId   int64 `json:"employee_store_id,omitempty" xml:"employee_store_id,omitempty"`
    // 分佣导购所属门店名称
    EmployeeStoreName   string `json:"employee_store_name,omitempty" xml:"employee_store_name,omitempty"`
    // 旗舰店名称
    SellerName   string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
    // 消费者入会时间
    BuyerJoinTime   string `json:"buyer_join_time,omitempty" xml:"buyer_join_time,omitempty"`
    // 分佣时间
    CommissionTime   string `json:"commission_time,omitempty" xml:"commission_time,omitempty"`
    // 消费者nick
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    // 导购工号
    WorkId   string `json:"work_id,omitempty" xml:"work_id,omitempty"`
    // 支付订单
    PayOrderId   string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
    // 商品id
    AuctionId   int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
    // 订单id
    OrderIdString   string `json:"order_id_string,omitempty" xml:"order_id_string,omitempty"`
    // 子订单id
    BizOrderIdString   string `json:"biz_order_id_string,omitempty" xml:"biz_order_id_string,omitempty"`
}

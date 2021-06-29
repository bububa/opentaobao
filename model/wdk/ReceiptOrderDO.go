package wdk

// ReceiptOrderDO 
type ReceiptOrderDO struct {
    // 实收金额，单位分
    ActualAmt   int64 `json:"actual_amt,omitempty" xml:"actual_amt,omitempty"`
    // 应收金额
    AskAmt   int64 `json:"ask_amt,omitempty" xml:"ask_amt,omitempty"`
    // 找零金额
    ChangeAmt   int64 `json:"change_amt,omitempty" xml:"change_amt,omitempty"`
    // 折扣优惠金额
    DiscountAmt   int64 `json:"discount_amt,omitempty" xml:"discount_amt,omitempty"`
    // 会员卡号
    MemberCardNum   string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
    // 会员优惠
    MemberDiscount   string `json:"member_discount,omitempty" xml:"member_discount,omitempty"`
    // 收银员姓名
    OpName   string `json:"op_name,omitempty" xml:"op_name,omitempty"`
    // 收银员工号
    OpNum   string `json:"op_num,omitempty" xml:"op_num,omitempty"`
    // 原款机号
    OriginalPosNo   string `json:"original_pos_no,omitempty" xml:"original_pos_no,omitempty"`
    // 原流水号
    OriginalSerialNum   string `json:"original_serial_num,omitempty" xml:"original_serial_num,omitempty"`
    // 原交易时间
    OriginalTrdTime   string `json:"original_trd_time,omitempty" xml:"original_trd_time,omitempty"`
    // 损溢金额
    OverflowAmt   int64 `json:"overflow_amt,omitempty" xml:"overflow_amt,omitempty"`
    // 支付渠道
    PayChannels   []ReceiptPayChannelDo `json:"pay_channels,omitempty" xml:"pay_channels>receipt_pay_channel_do,omitempty"`
    // 款机号
    PosNo   string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
    // 流水号
    SerialNum   string `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
    // 门店号
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 子订单信息
    SubOrders   []ReceiptSubOrderDo `json:"sub_orders,omitempty" xml:"sub_orders>receipt_sub_order_do,omitempty"`
    // 成交时间
    TrdTime   string `json:"trd_time,omitempty" xml:"trd_time,omitempty"`
    // 交易类型
    TrdType   int64 `json:"trd_type,omitempty" xml:"trd_type,omitempty"`
    // 来源，商家级别，当前取值：RENRENLE / SANJIANG
    OrderFrom   string `json:"order_from,omitempty" xml:"order_from,omitempty"`
    // 商户码
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    // 阿里用户ID
    AliUserid   string `json:"ali_userid,omitempty" xml:"ali_userid,omitempty"`
}

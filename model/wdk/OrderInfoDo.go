package wdk

// OrderInfoDo 
type OrderInfoDo struct {

    // 小票付款渠道
    PayChannels   []ReceiptPayChannelDo `json:"pay_channels,omitempty"`

    // 商品明细
    SubOrders   []SubOrderInfoDo `json:"sub_orders,omitempty"`

    // 来源，商家级别，当前取值：RENRENLE / SANJIANG
    OrderFrom   string `json:"order_from,omitempty"`

    // 原交易时间
    OriginalTrdTime   string `json:"original_trd_time,omitempty"`

    // 原流水号
    OriginalSerialNum   string `json:"original_serial_num,omitempty"`

    // 原款机号
    OriginalPosNo   string `json:"original_pos_no,omitempty"`

    // 折扣优惠金额
    DiscountAmt   int64 `json:"discount_amt,omitempty"`

    // 损溢金额
    OverflowAmt   int64 `json:"overflow_amt,omitempty"`

    // 会员优惠
    MemberDiscount   int64 `json:"member_discount,omitempty"`

    // 会员卡号
    MemberCardNum   string `json:"member_card_num,omitempty"`

    // 找零金额
    ChangeAmt   int64 `json:"change_amt,omitempty"`

    // 实收金额
    ActualAmt   int64 `json:"actual_amt,omitempty"`

    // 应收金额
    AskAmt   int64 `json:"ask_amt,omitempty"`

    // 收银员姓名
    OpName   string `json:"op_name,omitempty"`

    // 收银员工号
    OpNum   string `json:"op_num,omitempty"`

    // 交易类型
    TrdType   int64 `json:"trd_type,omitempty"`

    // 成交时间
    TrdTime   string `json:"trd_time,omitempty"`

    // 流水号
    SerialNum   string `json:"serial_num,omitempty"`

    // 款机号
    PosNo   string `json:"pos_no,omitempty"`

    // 渠道店id
    StoreId   string `json:"store_id,omitempty"`

    // 阿里userID
    AliUserId   string `json:"ali_user_id,omitempty"`

}

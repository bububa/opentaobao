package wdk

// MerchantOrderInfo 
type MerchantOrderInfo struct {

    // 渠道订单id
    ChannelOrderId   string `json:"channel_order_id,omitempty"`

    // 真实手机号
    RealPhone   string `json:"real_phone,omitempty"`

    // 投放跟踪id
    TraceId   string `json:"trace_id,omitempty"`

    // 折扣优惠金额
    DiscountAmt   int64 `json:"discount_amt,omitempty"`

    // 会员卡号
    MemberCardNum   string `json:"member_card_num,omitempty"`

    // 订单ID，商家订单流水号
    OrderId   string `json:"order_id,omitempty"`

    // 作用在父订单的优惠券信息
    CouponInfo   string `json:"coupon_info,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 履约邮费
    PostFee   int64 `json:"post_fee,omitempty"`

    // pos机号
    PosNo   string `json:"pos_no,omitempty"`

    // 子订单信息
    SubOrderList   []MerchantSubOrderInfo `json:"sub_order_list,omitempty"`

    // 经营店
    StoreId   string `json:"store_id,omitempty"`

    // 渠道（分类）
    OrderChannel   string `json:"order_channel,omitempty"`

    // 作用在父订单的活动信息
    ActivityInfo   string `json:"activity_info,omitempty"`

    // 总金额
    TotalAmt   int64 `json:"total_amt,omitempty"`

    // 渠道用户ID
    ChannelUserId   string `json:"channel_user_id,omitempty"`

    // 支付方式
    PayChannelList   []MerchantOrderPaymentInfo `json:"pay_channel_list,omitempty"`

    // 扩项属性，键值对的json
    ExtendProperty   string `json:"extend_property,omitempty"`

    // 外部门店编码
    OutStoreId   string `json:"out_store_id,omitempty"`

    // 商家会员ID
    MemberId   string `json:"member_id,omitempty"`

    // 实付金额
    ActualAmt   int64 `json:"actual_amt,omitempty"`

    // isv系统中完整的订单信息
    OriginWholeData   string `json:"origin_whole_data,omitempty"`

    // 商家侧统一用户标识ID
    UnionUserId   string `json:"union_user_id,omitempty"`

    // 上传数据统计信息
    UploadBatchInfo   *UploadBatchInfo `json:"upload_batch_info,omitempty"`

    // 订单来源渠道细粒度分类
    CustomizeOrderChannel   string `json:"customize_order_channel,omitempty"`

}

package openmall

// TopRefundVo 
type TopRefundVo struct {
    // 退款单ID
    RefundId   int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    // 对应订单ID
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    // 退款单状态，此接口存在延迟，详情参考taobao.openmall.refund.get
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 退款单创建时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 请求退款金额
    RefundFee   string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 货物状态。可选值BUYER_NOT_RECEIVED (买家未收到货) BUYER_RECEIVED (买家已收到货) BUYER_RETURNED_GOODS (买家已退货)
    GoodStatus   string `json:"good_status,omitempty" xml:"good_status,omitempty"`
    // 买家是否需要退货。可选值:true(是),false(否)
    HasGoodReturn   bool `json:"has_good_return,omitempty" xml:"has_good_return,omitempty"`
    // 更新时间
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    // 商品数量
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    // 商品ID
    NumIid   int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
    // 实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    // 退款阶段，可选值：onsale/aftersale
    RefundPhase   string `json:"refund_phase,omitempty" xml:"refund_phase,omitempty"`
    // 退款超时结构
    RefundRemindTimeout   *RefundRemindTimeout `json:"refund_remind_timeout,omitempty" xml:"refund_remind_timeout,omitempty"`
    // 创建交易时的物流方式。 可选值：ems, express, post, free
    ShippingType   string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
    // 商品SKU信息
    Sku   string `json:"sku,omitempty" xml:"sku,omitempty"`
    // 交易总金额
    TotalFee   string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
    // 卖家退货地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 退货运单号
    Sid   string `json:"sid,omitempty" xml:"sid,omitempty"`
    // 物流公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 退货邮编
    PostCode   string `json:"post_code,omitempty" xml:"post_code,omitempty"`
    // 退货时收货人电话
    FixPhone   string `json:"fix_phone,omitempty" xml:"fix_phone,omitempty"`
    // 退货时收货人手机号
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 退货时收货人姓名
    ConsigneeFullName   string `json:"consignee_full_name,omitempty" xml:"consignee_full_name,omitempty"`
    // 当该退款单为 未发货仅退款 申请时，因卖家坚持发货而导致关单的情况下，该字段值为true；其余条件为空或false
    ClosedBySellerShip   bool `json:"closed_by_seller_ship,omitempty" xml:"closed_by_seller_ship,omitempty"`
}

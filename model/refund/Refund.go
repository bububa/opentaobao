package refund

// Refund 
type Refund struct {
    // 物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS).
    ShippingType   string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
    // 不需客服介入1;需要客服介入2;客服已经介入3;客服初审完成 4;客服主管复审失败5;客服处理完成6;
    CsStatus   int64 `json:"cs_status,omitempty" xml:"cs_status,omitempty"`
    // 退款先行垫付默认的未申请状态 0;退款先行垫付申请中  1;退款先行垫付，垫付完成 2;退款先行垫付，卖家拒绝收货 3;退款先行垫付，垫付关闭 4;退款先行垫付，垫付分账成功 5;
    AdvanceStatus   int64 `json:"advance_status,omitempty" xml:"advance_status,omitempty"`
    // 分账给淘宝的钱
    SplitTaobaoFee   string `json:"split_taobao_fee,omitempty" xml:"split_taobao_fee,omitempty"`
    // 分账给卖家的钱
    SplitSellerFee   string `json:"split_seller_fee,omitempty" xml:"split_seller_fee,omitempty"`
    // 退款单号
    RefundId   string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    // 淘宝交易单号
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    // 子订单号。如果是单笔交易oid会等于tid
    Oid   int64 `json:"oid,omitempty" xml:"oid,omitempty"`
    // 支付宝交易号
    AlipayNo   string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
    // 交易总金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
    TotalFee   string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 退款申请时间。格式:yyyy-MM-dd HH:mm:ss
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 更新时间。格式:yyyy-MM-dd HH:mm:ss
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    // 退款对应的订单交易状态。可选值TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) 取自&quot;http://open.taobao.com/dev/index.php/%E4%BA%A4%E6%98%93%E7%8A%B6%E6%80%81&quot;
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 货物状态。可选值BUYER_NOT_RECEIVED (买家未收到货) BUYER_RECEIVED (买家已收到货) BUYER_RETURNED_GOODS (买家已退货)
    GoodStatus   string `json:"good_status,omitempty" xml:"good_status,omitempty"`
    // 买家是否需要退货。可选值:true(是),false(否)
    HasGoodReturn   bool `json:"has_good_return,omitempty" xml:"has_good_return,omitempty"`
    // 退还金额(退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分
    RefundFee   string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 支付给卖家的金额(交易总金额-退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    // 退款原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 退款说明
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 商品标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 商品购买数量
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    // 退货时间。格式:yyyy-MM-dd HH:mm:ss
    GoodReturnTime   string `json:"good_return_time,omitempty" xml:"good_return_time,omitempty"`
    // 物流公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 退货运单号
    Sid   string `json:"sid,omitempty" xml:"sid,omitempty"`
    // 卖家收货地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 退款超时结构RefundRemindTimeout
    RefundRemindTimeout   *RefundRemindTimeout `json:"refund_remind_timeout,omitempty" xml:"refund_remind_timeout,omitempty"`
    // 申请退款的商品数字编号
    NumIid   int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
    // 退款阶段，可选值：onsale/aftersale
    RefundPhase   string `json:"refund_phase,omitempty" xml:"refund_phase,omitempty"`
    // 退款版本号（时间戳）
    RefundVersion   int64 `json:"refund_version,omitempty" xml:"refund_version,omitempty"`
    // 商品SKU信息
    Sku   string `json:"sku,omitempty" xml:"sku,omitempty"`
    // 退款扩展属性
    Attribute   string `json:"attribute,omitempty" xml:"attribute,omitempty"`
    // 商品外部商家编码
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 退款约束，可选值：cannot_refuse（不允许操作），refund_onweb（需要到网页版操作）
    OperationContraint   string `json:"operation_contraint,omitempty" xml:"operation_contraint,omitempty"`
    // 买家账号的OpenUID
    BuyerOpenUid   string `json:"buyer_open_uid,omitempty" xml:"buyer_open_uid,omitempty"`
    // 逆向特殊部分退类型：价保服务、退差返现
    SpecialRefundType   string `json:"special_refund_type,omitempty" xml:"special_refund_type,omitempty"`
}

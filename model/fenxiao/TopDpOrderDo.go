package fenxiao

// TopDpOrderDo 
type TopDpOrderDo struct {

    // 供应商来源网站, values: taobao, alibaba
    
    SupplierFrom   string `json:"supplier_from,omitempty" xml:"supplier_from,omitempty"`
    

    // 供应商在来源网站的帐号名。
    
    SupplierUsername   string `json:"supplier_username,omitempty" xml:"supplier_username,omitempty"`
    

    // 分销商在来源网站的帐号名。
    
    DistributorUsername   string `json:"distributor_username,omitempty" xml:"distributor_username,omitempty"`
    

    // 采购单创建时间。格式:yyyy-MM-dd HH:mm:ss
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 支付宝交易号。
    
    AlipayNo   string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
    

    // 采购单总额（不含邮费,精确到2位小数;单位:元。如:200.07，表示:200元7分 )
    
    TotalFee   string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
    

    // 采购单邮费。(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
    
    PostFee   string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    

    // 分销商实付金额。(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
    
    DistributorPayment   string `json:"distributor_payment,omitempty" xml:"distributor_payment,omitempty"`
    

    // 订单快照URL
    
    SnapshotUrl   string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
    

    // 付款时间。格式:yyyy-MM-dd HH:mm:ss
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 物流发货时间。格式:yyyy-MM-dd HH:mm:ss
    
    ConsignTime   string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
    

    // 交易修改时间。格式:yyyy-MM-dd HH:mm:ss
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

    // 买家详细信息
    
    Receiver   *TopReceiverDo `json:"receiver,omitempty" xml:"receiver,omitempty"`
    

    // 供应商备注
    
    SupplierMemo   string `json:"supplier_memo,omitempty" xml:"supplier_memo,omitempty"`
    

    // 分销流水号，分销平台产生的主键
    
    FenxiaoId   int64 `json:"fenxiao_id,omitempty" xml:"fenxiao_id,omitempty"`
    

    // 支付方式：ALIPAY_SURETY（支付宝担保交易）、ALIPAY_CHAIN（分账交易）、TRANSFER（线下转账）、PREPAY（预存款）、IMMEDIATELY（即时到账）、CASHGOODS（先款后货）、ACCOUNT_PERIOD（账期支付）
    
    PayType   string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    

    // 分销方式：AGENT（代销）、DEALER（经销）
    
    TradeType   string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    

    // 分销商来源网站（taobao）。
    
    DistributorFrom   string `json:"distributor_from,omitempty" xml:"distributor_from,omitempty"`
    

    // 供应商交易ID 非采购单ID，如果改发货状态 是需要该ID，ID在用户未付款前为0，付款后有具体值（发货时使用该ID）
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 采购单交易状态。可选值： <br>WAIT_BUYER_PAY(等待付款)<br> WAIT_SELLER_SEND_GOODS(已付款，待发货）<br> WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)<br> PAID_FORBID_CONSIGN(已付款，禁止发货 ps:只有大快消行业的才有)<br> TRADE_FINISHED(交易成功)<br> TRADE_CLOSED(交易关闭)<br> WAIT_BUYER_CONFIRM_GOODS_ACOUNTED(已付款（已分账），已发货。只对代销分账支持)<br> PAY_ACOUNTED_GOODS_CONFIRM （已分账发货成功）<br> PAY_WAIT_ACOUNT_GOODS_CONFIRM（已付款，确认收货）
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 买家nick，供应商查询不会返回买家昵称，分销商查询才会返回。
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 采购单留言。（代销模式下信息包括买家留言和分销商留言）
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

    // 主订单ID （经销不显示）
    
    TcOrderId   int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
    

    // 配送方式，FAST(快速)、EMS、ORDINARY(平邮)、SELLER(卖家包邮)
    
    Shipping   string `json:"shipping,omitempty" xml:"shipping,omitempty"`
    

    // 物流公司
    
    LogisticsCompanyName   string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
    

    // 运单号
    
    LogisticsId   string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
    

    // 自定义key
    
    IsvCustomKey   []string `json:"isv_custom_key,omitempty" xml:"isv_custom_key>string,omitempty"`
    

    // 自定义值
    
    IsvCustomValue   []string `json:"isv_custom_value,omitempty" xml:"isv_custom_value>string,omitempty"`
    

    // 交易结束时间
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 返回供应商备注旗帜vlaue在1-5之间。非1-5之间，都采用1作为默认。 1:红色 2:黄色 3:绿色 4:蓝色 5:粉红色
    
    SupplierFlag   int64 `json:"supplier_flag,omitempty" xml:"supplier_flag,omitempty"`
    

    // 买家支付给分销商的总金额。注意买家购买的商品可能不是全部来自同一供货商，请同时参考子单上的相关金额。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
    
    BuyerPayment   string `json:"buyer_payment,omitempty" xml:"buyer_payment,omitempty"`
    

    // 采购单留言列表
    
    OrderMessages   []OrderMessages `json:"order_messages,omitempty" xml:"order_messages,omitempty"`
    

    // 子订单的详细信息列表。
    
    SubPurchaseOrders   []SubOrderDetail `json:"sub_purchase_orders,omitempty" xml:"sub_purchase_orders,omitempty"`
    

    // 加密后的买家淘宝ID，长度为32位
    
    BuyerTaobaoId   string `json:"buyer_taobao_id,omitempty" xml:"buyer_taobao_id,omitempty"`
    

    // 主订单属性信息，key-value形式：<br/>orderNovice ：订单发票抬头；<br/>orderNoviceContent ：代表发票明细
    
    Features   []FeatureDo `json:"features,omitempty" xml:"features,omitempty"`
    

    // [架海金梁独有字段，非架海金梁用户请勿关心]子单物流发货信息
    
    LogisticsInfos   []ErpLogisticsInfo `json:"logistics_infos,omitempty" xml:"logistics_infos,omitempty"`
    

    // 支付宝交易号，国际担保交易用
    
    AlipayOrderNo   string `json:"alipay_order_no,omitempty" xml:"alipay_order_no,omitempty"`
    

    // 已执行确认收货的金额
    
    ConfirmPaidFeeYuan   string `json:"confirm_paid_fee_yuan,omitempty" xml:"confirm_paid_fee_yuan,omitempty"`
    

    // 渠道编码，市场
    
    ChannelCode   int64 `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
    

}

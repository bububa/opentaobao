package openmall

// TopTradeDetailVo 
type TopTradeDetailVo struct {
    // 订单ID
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    // 外部outid
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    // 实付金额
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    // 邮费
    PostFee   string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    // 商品ID
    NumIid   int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
    // 商品数量
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    // 订单状态，该字段存在延迟，请参考taoboa.openmall.trade.get
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 支付宝交易号
    AlipayNo   string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
    // 下单用户渠道
    BuyerChannel   string `json:"buyer_channel,omitempty" xml:"buyer_channel,omitempty"`
    // 下单用户渠道openid
    BuyerChannelId   string `json:"buyer_channel_id,omitempty" xml:"buyer_channel_id,omitempty"`
    // 买家备注
    BuyerMemo   string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
    // 卖家发货时间。格式:yyyy-MM-dd HH:mm:ss
    ConsignTime   string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
    // 交易创建时间。格式:yyyy-MM-dd HH:mm:ss
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 下单媒体渠道
    Distributor   string `json:"distributor,omitempty" xml:"distributor,omitempty"`
    // 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 商家的预计发货时间
    EstConTime   string `json:"est_con_time,omitempty" xml:"est_con_time,omitempty"`
    // 是否包含邮费
    HasPostFee   bool `json:"has_post_fee,omitempty" xml:"has_post_fee,omitempty"`
    // 订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空
    MarkDesc   string `json:"mark_desc,omitempty" xml:"mark_desc,omitempty"`
    // 交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    // 订单列表
    Orders   []TopOrderVo `json:"orders,omitempty" xml:"orders>top_order_vo,omitempty"`
    // 付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 收货人的详细地址
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    // 收货人的所在城市 注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面 建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
    ReceiverCity   string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
    // 收货人国籍
    ReceiverCountry   string `json:"receiver_country,omitempty" xml:"receiver_country,omitempty"`
    // 收货人的所在地区 注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面 建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
    ReceiverDistrict   string `json:"receiver_district,omitempty" xml:"receiver_district,omitempty"`
    // 收货人手机号
    ReceiverMobile   string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
    // 收货人姓名
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    // 收货人电话
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
    // 收货人的所在省份
    ReceiverState   string `json:"receiver_state,omitempty" xml:"receiver_state,omitempty"`
    // 收货人街道地址
    ReceiverTown   string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
    // 收货人的邮编
    ReceiverZip   string `json:"receiver_zip,omitempty" xml:"receiver_zip,omitempty"`
    // 商家
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递)
    ShippingType   string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
    // 运单号，订单实际物流单号取 orders中的invoice_no和logistics_company
    Sid   string `json:"sid,omitempty" xml:"sid,omitempty"`
    // 交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
    TotalFee   string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
    // 交易备注。
    TradeMemo   string `json:"trade_memo,omitempty" xml:"trade_memo,omitempty"`
}

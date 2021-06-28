package trade

// Trade 
type Trade struct {

    // 交易编号 (父订单的交易编号)
    
    Tid   string `json:"tid,omitempty" xml:"tid,omitempty"`
    

    // 商品购买数量。取值范围：大于零的整数,对于一个trade对应多个order的时候（一笔主订单，对应多笔子订单），num=0，num是一个跟商品关联的属性，一笔订单对应多比子订单的时候，主订单上的num无意义。
    
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    

    // 商品数字编号
    
    NumIid   int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
    

    // 交易状态。可选值:    * TRADE_NO_CREATE_PAY(没有创建支付宝交易)    * WAIT_BUYER_PAY(等待买家付款)    * SELLER_CONSIGNED_PART(卖家部分发货)    * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)    * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)    * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)    * TRADE_FINISHED(交易成功)    * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)    * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)    * PAY_PENDING(国际信用卡支付付款确认中)    * WAIT_PRE_AUTH_CONFIRM(0元购合约中)	* PAID_FORBID_CONSIGN(拼团中订单或者发货强管控的订单，已付款但禁止发货)
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 交易类型列表，同时查询多种交易类型可用逗号分隔。默认同时查询guarantee_trade, auto_delivery, ec, cod的4种交易类型的数据 可选值 fixed(一口价) auction(拍卖) guarantee_trade(一口价、拍卖) auto_delivery(自动发货) independent_simple_trade(旺店入门版交易) independent_shop_trade(旺店标准版交易) ec(直冲) cod(货到付款) fenxiao(分销) game_equipment(游戏装备) shopex_trade(ShopEX交易) netcn_trade(万网交易) external_trade(统一外部交易)o2o_offlinetrade（O2O交易）step (万人团)nopaid(无付款订单)pre_auth_type(预授权0元购机交易)
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
    
    TotalFee   string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
    

    // 交易创建时间。格式:yyyy-MM-dd HH:mm:ss
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 订单列表
    
    Orders   []Order `json:"orders,omitempty" xml:"orders,omitempty"`
    

    // 买家的openuid,入参fields中传入buyer_nick ，才能返回
    
    BuyerOpenUid   string `json:"buyer_open_uid,omitempty" xml:"buyer_open_uid,omitempty"`
    

    // 付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 卖家备注（与淘宝网上订单的卖家备注对应，只有卖家才能查看该字段）
    
    SellerMemo   string `json:"seller_memo,omitempty" xml:"seller_memo,omitempty"`
    

    // 买家备注（与淘宝网上订单的买家备注对应，只有买家才能查看该字段）
    
    BuyerMemo   string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
    

    // 邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分
    
    PostFee   string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    

    // 是否包含邮费。与available_confirm_fee同时使用。可选值:true(包含),false(不包含)
    
    HasPostFee   bool `json:"has_post_fee,omitempty" xml:"has_post_fee,omitempty"`
    

    // 优惠详情
    
    PromotionDetails   []PromotionDetail `json:"promotion_details,omitempty" xml:"promotion_details,omitempty"`
    

    // 分阶段付款的订单状态（例如万人团订单等），目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)
    
    StepTradeStatus   string `json:"step_trade_status,omitempty" xml:"step_trade_status,omitempty"`
    

    // 分阶段付款的已付金额（万人团订单已付金额）
    
    StepPaidFee   string `json:"step_paid_fee,omitempty" xml:"step_paid_fee,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
    
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    

    // 交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

    // 商品图片绝对途径
    
    PicPath   string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
    

    // 卖家是否已评价。可选值:true(已评价),false(未评价)
    
    SellerRate   bool `json:"seller_rate,omitempty" xml:"seller_rate,omitempty"`
    

    // 收货人的姓名
    
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    

    // 收货人的所在省份
    
    ReceiverState   string `json:"receiver_state,omitempty" xml:"receiver_state,omitempty"`
    

    // 收货人的详细地址
    
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    

    // 收货人的邮编
    
    ReceiverZip   string `json:"receiver_zip,omitempty" xml:"receiver_zip,omitempty"`
    

    // 收货人的手机号码
    
    ReceiverMobile   string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
    

    // 收货人的电话号码
    
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
    

    // 卖家发货时间。格式:yyyy-MM-dd HH:mm:ss
    
    ConsignTime   string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
    

    // 卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
    
    ReceivedPayment   string `json:"received_payment,omitempty" xml:"received_payment,omitempty"`
    

    // 商家的预计发货时间
    
    EstConTime   string `json:"est_con_time,omitempty" xml:"est_con_time,omitempty"`
    

    // 收货人国籍
    
    ReceiverCountry   string `json:"receiver_country,omitempty" xml:"receiver_country,omitempty"`
    

    // 收货人街道地址
    
    ReceiverTown   string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
    

    // 天猫国际官网直供主订单关税税费
    
    OrderTaxFee   string `json:"order_tax_fee,omitempty" xml:"order_tax_fee,omitempty"`
    

    // 满返红包的金额；如果没有满返红包，则值为 0.00
    
    PaidCouponFee   string `json:"paid_coupon_fee,omitempty" xml:"paid_coupon_fee,omitempty"`
    

    // 门店自提，总店发货，分店取货的门店自提订单标识
    
    ShopPick   string `json:"shop_pick,omitempty" xml:"shop_pick,omitempty"`
    

    // 同tid
    
    TidStr   string `json:"tid_str,omitempty" xml:"tid_str,omitempty"`
    

    // 为tmall.daogoubao.cloudstore时表示云店链路
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    

    // 值为1，且bizCode不为tmall.daogoubao.cloudstore时，为旗舰店订单
    
    CloudStore   string `json:"cloud_store,omitempty" xml:"cloud_store,omitempty"`
    

    // 预售单为true，否则false (云店订单专用)
    
    NewPresell   bool `json:"new_presell,omitempty" xml:"new_presell,omitempty"`
    

    // 优享购为true，否则false(云店订单专用)
    
    YouXiang   bool `json:"you_xiang,omitempty" xml:"you_xiang,omitempty"`
    

    // 默认为0,0 表示用户主动支付1 表示系统代扣2 表示保险赔付
    
    PayChannel   string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
    

    // 交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 可以使用trade.promotion_details查询系统优惠系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：元。如：200.07，表示：200元7分
    
    DiscountFee   string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
    

    // 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 买家留言
    
    BuyerMessage   string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
    

    // 买家备注旗帜（与淘宝网上订单的买家备注旗帜对应，只有买家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5
    
    BuyerFlag   int64 `json:"buyer_flag,omitempty" xml:"buyer_flag,omitempty"`
    

    // 卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5
    
    SellerFlag   int64 `json:"seller_flag,omitempty" xml:"seller_flag,omitempty"`
    

    // 买家昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // top动态字段
    
    TradeAttr   string `json:"trade_attr,omitempty" xml:"trade_attr,omitempty"`
    

    // 订单中是否包含运费险订单，如果包含运费险订单返回true，不包含运费险订单，返回false
    
    HasYfx   bool `json:"has_yfx,omitempty" xml:"has_yfx,omitempty"`
    

    // 订单的运费险，单位为元
    
    YfxFee   string `json:"yfx_fee,omitempty" xml:"yfx_fee,omitempty"`
    

    // 使用信用卡支付金额数
    
    CreditCardFee   string `json:"credit_card_fee,omitempty" xml:"credit_card_fee,omitempty"`
    

    // 订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空
    
    MarkDesc   string `json:"mark_desc,omitempty" xml:"mark_desc,omitempty"`
    

    // 创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递),ems(EMS),virtual(虚拟发货)，25(次日必达)，26(预约配送)。
    
    ShippingType   string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
    

    // 买家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分
    
    BuyerCodFee   string `json:"buyer_cod_fee,omitempty" xml:"buyer_cod_fee,omitempty"`
    

    // 卖家手工调整金额，精确到2位小数，单位：元。如：200.07，表示：200元7分。来源于订单价格修改，如果有多笔子订单的时候，这个为0，单笔的话则跟[order].adjust_fee一样
    
    AdjustFee   string `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
    

    // 交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔
    
    TradeFrom   string `json:"trade_from,omitempty" xml:"trade_from,omitempty"`
    

    // 服务子订单列表
    
    ServiceOrders   []ServiceOrder `json:"service_orders,omitempty" xml:"service_orders,omitempty"`
    

    // 买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false
    
    BuyerRate   bool `json:"buyer_rate,omitempty" xml:"buyer_rate,omitempty"`
    

    // 收货人的所在城市<br/>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br/>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
    
    ReceiverCity   string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
    

    // 收货人的所在地区<br/>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br/>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
    
    ReceiverDistrict   string `json:"receiver_district,omitempty" xml:"receiver_district,omitempty"`
    

    // 物流标签
    
    ServiceTags   []LogisticsTag `json:"service_tags,omitempty" xml:"service_tags,omitempty"`
    

    // 导购宝=crm
    
    O2o   string `json:"o2o,omitempty" xml:"o2o,omitempty"`
    

    // 导购员id
    
    O2oGuideId   string `json:"o2o_guide_id,omitempty" xml:"o2o_guide_id,omitempty"`
    

    // 导购员门店id
    
    O2oShopId   string `json:"o2o_shop_id,omitempty" xml:"o2o_shop_id,omitempty"`
    

    // 导购员名称
    
    O2oGuideName   string `json:"o2o_guide_name,omitempty" xml:"o2o_guide_name,omitempty"`
    

    // 导购门店名称
    
    O2oShopName   string `json:"o2o_shop_name,omitempty" xml:"o2o_shop_name,omitempty"`
    

    // 导购宝提货方式，inshop：店内提货，online：线上发货
    
    O2oDelivery   string `json:"o2o_delivery,omitempty" xml:"o2o_delivery,omitempty"`
    

    // 交易扩展表信息
    
    TradeExt   *TradeExt `json:"trade_ext,omitempty" xml:"trade_ext,omitempty"`
    

    // 天猫电子凭证家装
    
    EticketServiceAddr   string `json:"eticket_service_addr,omitempty" xml:"eticket_service_addr,omitempty"`
    

    // 处方药未审核状态
    
    RxAuditStatus   string `json:"rx_audit_status,omitempty" xml:"rx_audit_status,omitempty"`
    

    // 时间段
    
    EsRange   string `json:"es_range,omitempty" xml:"es_range,omitempty"`
    

    // 时间
    
    EsDate   string `json:"es_date,omitempty" xml:"es_date,omitempty"`
    

    // 时间
    
    OsDate   string `json:"os_date,omitempty" xml:"os_date,omitempty"`
    

    // 时间段
    
    OsRange   string `json:"os_range,omitempty" xml:"os_range,omitempty"`
    

    // 订单中使用红包付款的金额
    
    CouponFee   int64 `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`
    

    // 分阶段交易的特权定金订单ID
    
    O2oEtOrderId   string `json:"o2o_et_order_id,omitempty" xml:"o2o_et_order_id,omitempty"`
    

    // 邮关订单
    
    PostGateDeclare   bool `json:"post_gate_declare,omitempty" xml:"post_gate_declare,omitempty"`
    

    // 跨境订单
    
    CrossBondedDeclare   bool `json:"cross_bonded_declare,omitempty" xml:"cross_bonded_declare,omitempty"`
    

    // 全渠道商品通相关字段
    
    OmnichannelParam   string `json:"omnichannel_param,omitempty" xml:"omnichannel_param,omitempty"`
    

    // 组合商品
    
    Assembly   string `json:"assembly,omitempty" xml:"assembly,omitempty"`
    

    // TOP拦截标识，0不拦截，1拦截
    
    TopHold   int64 `json:"top_hold,omitempty" xml:"top_hold,omitempty"`
    

    // 星盘标识字段
    
    OmniAttr   string `json:"omni_attr,omitempty" xml:"omni_attr,omitempty"`
    

    // 星盘业务字段
    
    OmniParam   string `json:"omni_param,omitempty" xml:"omni_param,omitempty"`
    

    // 聚划算一起买字段
    
    ForbidConsign   int64 `json:"forbid_consign,omitempty" xml:"forbid_consign,omitempty"`
    

    // 采购订单标识
    
    Identity   string `json:"identity,omitempty" xml:"identity,omitempty"`
    

    // 天猫拼团拦截标示
    
    TeamBuyHold   int64 `json:"team_buy_hold,omitempty" xml:"team_buy_hold,omitempty"`
    

    // shareGroupHold
    
    ShareGroupHold   int64 `json:"share_group_hold,omitempty" xml:"share_group_hold,omitempty"`
    

    // 天猫国际拦截
    
    OfpHold   int64 `json:"ofp_hold,omitempty" xml:"ofp_hold,omitempty"`
    

    // 组装O2O多阶段尾款订单的明细数据 总阶段数，当前阶数，阶段金额（单位：分），支付状态，例如 3_1_100_paid ; 3_2_2000_nopaid
    
    O2oStepTradeDetail   string `json:"o2o_step_trade_detail,omitempty" xml:"o2o_step_trade_detail,omitempty"`
    

    // 特权定金订单的尾款订单ID
    
    O2oStepOrderId   string `json:"o2o_step_order_id,omitempty" xml:"o2o_step_order_id,omitempty"`
    

    // 分阶段订单的特权定金抵扣金额，单位：分
    
    O2oVoucherPrice   string `json:"o2o_voucher_price,omitempty" xml:"o2o_voucher_price,omitempty"`
    

    // 天猫国际计税优惠金额
    
    OrderTaxPromotionFee   string `json:"order_tax_promotion_fee,omitempty" xml:"order_tax_promotion_fee,omitempty"`
    

    // 聚划算火拼标记
    
    DelayCreateDelivery   int64 `json:"delay_create_delivery,omitempty" xml:"delay_create_delivery,omitempty"`
    

    // top定义订单类型
    
    Toptype   int64 `json:"toptype,omitempty" xml:"toptype,omitempty"`
    

    // serviceType
    
    ServiceType   string `json:"service_type,omitempty" xml:"service_type,omitempty"`
    

    // o2oServiceMobile
    
    O2oServiceMobile   string `json:"o2o_service_mobile,omitempty" xml:"o2o_service_mobile,omitempty"`
    

    // o2oServiceName
    
    O2oServiceName   string `json:"o2o_service_name,omitempty" xml:"o2o_service_name,omitempty"`
    

    // o2oServiceState
    
    O2oServiceState   string `json:"o2o_service_state,omitempty" xml:"o2o_service_state,omitempty"`
    

    // o2oServiceCity
    
    O2oServiceCity   string `json:"o2o_service_city,omitempty" xml:"o2o_service_city,omitempty"`
    

    // o2oServiceDistrict
    
    O2oServiceDistrict   string `json:"o2o_service_district,omitempty" xml:"o2o_service_district,omitempty"`
    

    // o2oServiceTown
    
    O2oServiceTown   string `json:"o2o_service_town,omitempty" xml:"o2o_service_town,omitempty"`
    

    // o2oServiceAddress
    
    O2oServiceAddress   string `json:"o2o_service_address,omitempty" xml:"o2o_service_address,omitempty"`
    

    // o2oStepTradeDetailNew
    
    O2oStepTradeDetailNew   string `json:"o2o_step_trade_detail_new,omitempty" xml:"o2o_step_trade_detail_new,omitempty"`
    

    // o2oXiaopiao
    
    O2oXiaopiao   string `json:"o2o_xiaopiao,omitempty" xml:"o2o_xiaopiao,omitempty"`
    

    // o2oContract
    
    O2oContract   string `json:"o2o_contract,omitempty" xml:"o2o_contract,omitempty"`
    

    // 新零售门店编码
    
    RetailStoreCode   string `json:"retail_store_code,omitempty" xml:"retail_store_code,omitempty"`
    

    // 新零售线下订单id
    
    RetailOutOrderId   string `json:"retail_out_order_id,omitempty" xml:"retail_out_order_id,omitempty"`
    

    // rechargeFee
    
    RechargeFee   string `json:"recharge_fee,omitempty" xml:"recharge_fee,omitempty"`
    

    // platformSubsidyFee
    
    PlatformSubsidyFee   string `json:"platform_subsidy_fee,omitempty" xml:"platform_subsidy_fee,omitempty"`
    

    // nrOffline
    
    NrOffline   string `json:"nr_offline,omitempty" xml:"nr_offline,omitempty"`
    

    // 网厅订单垂直表信息
    
    WttParam   string `json:"wtt_param,omitempty" xml:"wtt_param,omitempty"`
    

    // logisticsInfos
    
    LogisticsInfos   []LogisticsInfo `json:"logistics_infos,omitempty" xml:"logistics_infos,omitempty"`
    

    // nrStoreOrderId
    
    NrStoreOrderId   string `json:"nr_store_order_id,omitempty" xml:"nr_store_order_id,omitempty"`
    

    // 门店 ID
    
    NrShopId   string `json:"nr_shop_id,omitempty" xml:"nr_shop_id,omitempty"`
    

    // 门店名称
    
    NrShopName   string `json:"nr_shop_name,omitempty" xml:"nr_shop_name,omitempty"`
    

    // 导购员ID
    
    NrShopGuideId   string `json:"nr_shop_guide_id,omitempty" xml:"nr_shop_guide_id,omitempty"`
    

    // 导购员名称
    
    NrShopGuideName   string `json:"nr_shop_guide_name,omitempty" xml:"nr_shop_guide_name,omitempty"`
    

    // sortInfo
    
    SortInfo   string `json:"sort_info,omitempty" xml:"sort_info,omitempty"`
    

    // 1已排序 2不排序
    
    Sorted   int64 `json:"sorted,omitempty" xml:"sorted,omitempty"`
    

    // 一小时达不处理订单
    
    NrNoHandle   string `json:"nr_no_handle,omitempty" xml:"nr_no_handle,omitempty"`
    

    // isGift
    
    IsGift   bool `json:"is_gift,omitempty" xml:"is_gift,omitempty"`
    

    // doneeNick
    
    DoneeNick   string `json:"donee_nick,omitempty" xml:"donee_nick,omitempty"`
    

    // doneeUid
    
    DoneeOpenUid   string `json:"donee_open_uid,omitempty" xml:"donee_open_uid,omitempty"`
    

    // suningShopCode
    
    SuningShopCode   string `json:"suning_shop_code,omitempty" xml:"suning_shop_code,omitempty"`
    

    // suningShopValid
    
    SuningShopValid   int64 `json:"suning_shop_valid,omitempty" xml:"suning_shop_valid,omitempty"`
    

    // retailStoreId
    
    RetailStoreId   string `json:"retail_store_id,omitempty" xml:"retail_store_id,omitempty"`
    

    // isIstore
    
    IsIstore   bool `json:"is_istore,omitempty" xml:"is_istore,omitempty"`
    

    // ua
    
    Ua   string `json:"ua,omitempty" xml:"ua,omitempty"`
    

    // 截单时间
    
    CutoffMinutes   string `json:"cutoff_minutes,omitempty" xml:"cutoff_minutes,omitempty"`
    

    // 时效：天
    
    EsTime   string `json:"es_time,omitempty" xml:"es_time,omitempty"`
    

    // 发货时间
    
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    

    // 揽收时间
    
    CollectTime   string `json:"collect_time,omitempty" xml:"collect_time,omitempty"`
    

    // 派送时间
    
    DispatchTime   string `json:"dispatch_time,omitempty" xml:"dispatch_time,omitempty"`
    

    // 签收时间
    
    SignTime   string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
    

    // 派送CP
    
    DeliveryCps   string `json:"delivery_cps,omitempty" xml:"delivery_cps,omitempty"`
    

    // linkedmall透传参数
    
    LinkedmallExtInfo   string `json:"linkedmall_ext_info,omitempty" xml:"linkedmall_ext_info,omitempty"`
    

    // 新零售全渠道订单：订单类型，自提订单：pickUp，电商发货：tmall，门店发货（配送、骑手）：storeSend
    
    RtOmniSendType   string `json:"rt_omni_send_type,omitempty" xml:"rt_omni_send_type,omitempty"`
    

    // 新零售全渠道订单：发货门店ID
    
    RtOmniStoreId   string `json:"rt_omni_store_id,omitempty" xml:"rt_omni_store_id,omitempty"`
    

    // 新零售全渠道订单：商家自有发货门店编码
    
    RtOmniOuterStoreId   string `json:"rt_omni_outer_store_id,omitempty" xml:"rt_omni_outer_store_id,omitempty"`
    

    // 同城预约配送开始时间
    
    TcpsStart   string `json:"tcps_start,omitempty" xml:"tcps_start,omitempty"`
    

    // 同城业务类型,com.tmall.dsd:定时送,storeDsd-fn-3-1:淘速达3公里蜂鸟配送
    
    TcpsCode   string `json:"tcps_code,omitempty" xml:"tcps_code,omitempty"`
    

    // 同城预约配送结束时间
    
    TcpsEnd   string `json:"tcps_end,omitempty" xml:"tcps_end,omitempty"`
    

    // 
    
    MTariffFee   string `json:"m_tariff_fee,omitempty" xml:"m_tariff_fee,omitempty"`
    

    // 时效服务身份，如tmallPromise代表天猫时效承诺
    
    TimingPromise   string `json:"timing_promise,omitempty" xml:"timing_promise,omitempty"`
    

    // 时效服务字段，服务字段，会有多个服务值，以英文半角逗号&quot;,&quot;切割
    
    PromiseService   string `json:"promise_service,omitempty" xml:"promise_service,omitempty"`
    

    // 苏宁预约安装，用户安装时间段
    
    OiRange   string `json:"oi_range,omitempty" xml:"oi_range,omitempty"`
    

    // 苏宁预约安装，用户安装时间
    
    OiDate   string `json:"oi_date,omitempty" xml:"oi_date,omitempty"`
    

    // 苏宁预约安装，暂不安装
    
    HoldInstall   string `json:"hold_install,omitempty" xml:"hold_install,omitempty"`
    

    // 外部会员id
    
    OuterPartnerMemberId   string `json:"outer_partner_member_id,omitempty" xml:"outer_partner_member_id,omitempty"`
    

    // 叶子分类
    
    RootCat   string `json:"root_cat,omitempty" xml:"root_cat,omitempty"`
    

    // 1-gifting订单
    
    Gifting   string `json:"gifting,omitempty" xml:"gifting,omitempty"`
    

    // 1-coffee gifting订单
    
    GiftingTakeout   string `json:"gifting_takeout,omitempty" xml:"gifting_takeout,omitempty"`
    

    // 订单来源
    
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    

    // 居然之家同城站订单类型 deposit：预约到店，direct：直接购买，tail：尾款核销
    
    EasyHomeCityType   string `json:"easy_home_city_type,omitempty" xml:"easy_home_city_type,omitempty"`
    

    // 同城站关联订单号
    
    NrDepositOrderId   string `json:"nr_deposit_order_id,omitempty" xml:"nr_deposit_order_id,omitempty"`
    

    // 摊位id
    
    NrStoreCode   string `json:"nr_store_code,omitempty" xml:"nr_store_code,omitempty"`
    

    // 使用淘金币的数量，以分为单位，和订单标propoint中间那一段一样，没有返回null
    
    Propoint   string `json:"propoint,omitempty" xml:"propoint,omitempty"`
    

    // 1-周期送订单
    
    ZqsOrderTag   string `json:"zqs_order_tag,omitempty" xml:"zqs_order_tag,omitempty"`
    

    // 天鲜配冰柜id
    
    TxpFreezerId   string `json:"txp_freezer_id,omitempty" xml:"txp_freezer_id,omitempty"`
    

    // 天鲜配自提方式
    
    TxpReceiveMethod   string `json:"txp_receive_method,omitempty" xml:"txp_receive_method,omitempty"`
    

    // 同城购门店ID
    
    BrandLightShopStoreId   string `json:"brand_light_shop_store_id,omitempty" xml:"brand_light_shop_store_id,omitempty"`
    

    // 同城购订单source
    
    BrandLightShopSource   string `json:"brand_light_shop_source,omitempty" xml:"brand_light_shop_source,omitempty"`
    

    // 透出的额外信息
    
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
    

    // 收货地址有变更，返回&quot;1&quot;
    
    Lm   string `json:"lm,omitempty" xml:"lm,omitempty"`
    

    // 新康众定制数据
    
    NczExtAttr   string `json:"ncz_ext_attr,omitempty" xml:"ncz_ext_attr,omitempty"`
    

    // 标识完美履约订单
    
    IsWmly   string `json:"is_wmly,omitempty" xml:"is_wmly,omitempty"`
    

    // 全渠道包裹信息
    
    OmniPackage   string `json:"omni_package,omitempty" xml:"omni_package,omitempty"`
    

    // 购物金信息输出
    
    ExpandcardInfo   *ExpandCardInfo `json:"expandcard_info,omitempty" xml:"expandcard_info,omitempty"`
    

    // 苹果发票详情
    
    InvoiceDetailAfterRefund   string `json:"invoice_detail_after_refund,omitempty" xml:"invoice_detail_after_refund,omitempty"`
    

    // 苹果发票详情
    
    InvoiceDetailPay   string `json:"invoice_detail_pay,omitempty" xml:"invoice_detail_pay,omitempty"`
    

    // 苹果发票详情
    
    InvoiceDetailMidRefund   string `json:"invoice_detail_mid_refund,omitempty" xml:"invoice_detail_mid_refund,omitempty"`
    

    // 买卡订单本金
    
    ExpandCardBasicPrice   string `json:"expand_card_basic_price,omitempty" xml:"expand_card_basic_price,omitempty"`
    

    // 买卡订单权益金
    
    ExpandCardExpandPrice   string `json:"expand_card_expand_price,omitempty" xml:"expand_card_expand_price,omitempty"`
    

    // 用卡订单所用的本金
    
    ExpandCardBasicPriceUsed   string `json:"expand_card_basic_price_used,omitempty" xml:"expand_card_basic_price_used,omitempty"`
    

    // 用卡订单所用的权益金
    
    ExpandCardExpandPriceUsed   string `json:"expand_card_expand_price_used,omitempty" xml:"expand_card_expand_price_used,omitempty"`
    

    // 是否是Openmall订单
    
    IsOpenmall   bool `json:"is_openmall,omitempty" xml:"is_openmall,omitempty"`
    

    // asdp业务身份
    
    AsdpBizType   string `json:"asdp_biz_type,omitempty" xml:"asdp_biz_type,omitempty"`
    

    // (收货人+手机号+收货地址+create）4字段返回值都都不能为空，否则生成不了oaid
    
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    

    // 是否是码上收订单
    
    VLogisticsCreate   bool `json:"v_logistics_create,omitempty" xml:"v_logistics_create,omitempty"`
    

    // 是否是非物流订单
    
    QRPay   bool `json:"q_r_pay,omitempty" xml:"q_r_pay,omitempty"`
    

    // 关联下单订单
    
    OrderFollowId   string `json:"order_follow_id,omitempty" xml:"order_follow_id,omitempty"`
    

    // 无物流信息返回true，平台属性，业务不要依赖
    
    NoShipping   bool `json:"no_shipping,omitempty" xml:"no_shipping,omitempty"`
    

    // 送货上门标
    
    AsdpAds   string `json:"asdp_ads,omitempty" xml:"asdp_ads,omitempty"`
    

    // 是否屏蔽发货
    
    IsShShip   bool `json:"is_sh_ship,omitempty" xml:"is_sh_ship,omitempty"`
    

    // 抢单状态0,未处理待分发;1,抢单中;2,已抢单;3,已发货;-1,超时;-2,处理异常;-3,匹配失败;-4,取消抢单;-5,退款取消;-9,逻辑删除
    
    O2oSnatchStatus   string `json:"o2o_snatch_status,omitempty" xml:"o2o_snatch_status,omitempty"`
    

    // 垂直市场
    
    Market   string `json:"market,omitempty" xml:"market,omitempty"`
    

    // 电子凭证扫码购-扫码支付订单type
    
    EtType   string `json:"et_type,omitempty" xml:"et_type,omitempty"`
    

    // 扫码购关联门店
    
    EtShopId   int64 `json:"et_shop_id,omitempty" xml:"et_shop_id,omitempty"`
    

    // 门店预约自提订单标
    
    Obs   string `json:"obs,omitempty" xml:"obs,omitempty"`
    

    // 透出前置营销工具
    
    Pmtp   string `json:"pmtp,omitempty" xml:"pmtp,omitempty"`
    

    // 判断订单是否有买家留言，有买家留言返回true，否则返回false
    
    HasBuyerMessage   bool `json:"has_buyer_message,omitempty" xml:"has_buyer_message,omitempty"`
    

    // threeplTiming
    
    ThreeplTiming   bool `json:"threepl_timing,omitempty" xml:"threepl_timing,omitempty"`
    

    // 是否是智慧门店订单，只有true，或者 null 两种情况
    
    IsO2oPassport   bool `json:"is_o2o_passport,omitempty" xml:"is_o2o_passport,omitempty"`
    

    // tmallDelivery
    
    TmallDelivery   bool `json:"tmall_delivery,omitempty" xml:"tmall_delivery,omitempty"`
    

    // 天猫直送服务
    
    CnService   string `json:"cn_service,omitempty" xml:"cn_service,omitempty"`
    

}

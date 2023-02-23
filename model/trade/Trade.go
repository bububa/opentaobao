package trade

// Trade 结构体
type Trade struct {
	// 订单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 优惠详情
	PromotionDetails []PromotionDetail `json:"promotion_details,omitempty" xml:"promotion_details>promotion_detail,omitempty"`
	// 物流标签
	ServiceTags []LogisticsTag `json:"service_tags,omitempty" xml:"service_tags>logistics_tag,omitempty"`
	// 服务子订单列表
	ServiceOrders []ServiceOrders `json:"service_orders,omitempty" xml:"service_orders>service_orders,omitempty"`
	// 交易类型列表，同时查询多种交易类型可用逗号分隔。默认同时查询guarantee_trade, auto_delivery, ec, cod的4种交易类型的数据 可选值 fixed(一口价) auction(拍卖) guarantee_trade(一口价、拍卖) auto_delivery(自动发货) independent_simple_trade(旺店入门版交易) independent_shop_trade(旺店标准版交易) ec(直冲) cod(货到付款) fenxiao(分销) game_equipment(游戏装备) shopex_trade(ShopEX交易) netcn_trade(万网交易) external_trade(统一外部交易)o2o_offlinetrade（O2O交易）step (万人团)nopaid(无付款订单)pre_auth_type(预授权0元购机交易)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 交易创建时间。格式:yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 交易编号 (父订单的交易编号)
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 交易状态。可选值:    * TRADE_NO_CREATE_PAY(没有创建支付宝交易)    * WAIT_BUYER_PAY(等待买家付款)    * SELLER_CONSIGNED_PART(卖家部分发货)    * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)    * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)    * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)    * TRADE_FINISHED(交易成功)    * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)    * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)    * PAY_PENDING(国际信用卡支付付款确认中)    * WAIT_PRE_AUTH_CONFIRM(0元购合约中)	* PAID_FORBID_CONSIGN(拼团中订单或者发货强管控的订单，已付款但禁止发货)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 卖家备注（与淘宝网上订单的卖家备注对应，只有卖家才能查看该字段）
	SellerMemo string `json:"seller_memo,omitempty" xml:"seller_memo,omitempty"`
	// 买家备注（与淘宝网上订单的买家备注对应，只有买家才能查看该字段）
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分
	PostFee string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 分阶段付款的订单状态（例如万人团订单等），目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)
	StepTradeStatus string `json:"step_trade_status,omitempty" xml:"step_trade_status,omitempty"`
	// 分阶段付款的已付金额（万人团订单已付金额）
	StepPaidFee string `json:"step_paid_fee,omitempty" xml:"step_paid_fee,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 买家的openuid,入参fields中传入buyer_nick ，才能返回
	BuyerOpenUid string `json:"buyer_open_uid,omitempty" xml:"buyer_open_uid,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 物流运单号
	Sid string `json:"sid,omitempty" xml:"sid,omitempty"`
	// 建议使用trade.promotion_details查询系统优惠系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：元。如：200.07，表示：200元7分
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 卖家手工调整金额，精确到2位小数，单位：元。如：200.07，表示：200元7分。来源于订单价格修改，如果有多笔子订单的时候，这个为0，单笔的话则跟[order].adjust_fee一样
	AdjustFee string `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 卖家发货时间。格式:yyyy-MM-dd HH:mm:ss
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
	ReceivedPayment string `json:"received_payment,omitempty" xml:"received_payment,omitempty"`
	// 交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分
	CommissionFee string `json:"commission_fee,omitempty" xml:"commission_fee,omitempty"`
	// 支付宝交易号，如：2009112081173831
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 商品图片绝对途径
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。
	CodFee string `json:"cod_fee,omitempty" xml:"cod_fee,omitempty"`
	// 货到付款物流状态。初始状态 NEW_CREATED,接单成功 ACCEPTED_BY_COMPANY,接单失败 REJECTED_BY_COMPANY,接单超时 RECIEVE_TIMEOUT,揽收成功 TAKEN_IN_SUCCESS,揽收失败 TAKEN_IN_FAILED,揽收超时 TAKEN_TIMEOUT,签收成功 SIGN_IN,签收失败 REJECTED_BY_OTHER_SIDE,订单等待发送给物流公司 WAITING_TO_BE_SENT,用户取消物流订单 CANCELED
	CodStatus string `json:"cod_status,omitempty" xml:"cod_status,omitempty"`
	// 创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递),ems(EMS),virtual(虚拟发货)，25(次日必达)，26(预约配送)。
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// 物流到货时效截单时间，格式 HH:mm
	ArriveCutTime string `json:"arrive_cut_time,omitempty" xml:"arrive_cut_time,omitempty"`
	// 导购宝=crm
	O2o string `json:"o2o,omitempty" xml:"o2o,omitempty"`
	// 导购员id
	O2oGuideId string `json:"o2o_guide_id,omitempty" xml:"o2o_guide_id,omitempty"`
	// 导购员名称
	O2oGuideName string `json:"o2o_guide_name,omitempty" xml:"o2o_guide_name,omitempty"`
	// 导购员门店id
	O2oShopId string `json:"o2o_shop_id,omitempty" xml:"o2o_shop_id,omitempty"`
	// 导购门店名称
	O2oShopName string `json:"o2o_shop_name,omitempty" xml:"o2o_shop_name,omitempty"`
	// 导购宝提货方式，inshop：店内提货，online：线上发货
	O2oDelivery string `json:"o2o_delivery,omitempty" xml:"o2o_delivery,omitempty"`
	// 外部订单号
	O2oOutTradeId string `json:"o2o_out_trade_id,omitempty" xml:"o2o_out_trade_id,omitempty"`
	// 天猫汽车服务预约时间
	EtSerTime string `json:"et_ser_time,omitempty" xml:"et_ser_time,omitempty"`
	// 电子凭证预约门店地址
	EtShopName string `json:"et_shop_name,omitempty" xml:"et_shop_name,omitempty"`
	// 电子凭证核销门店地址
	EtVerifiedShopName string `json:"et_verified_shop_name,omitempty" xml:"et_verified_shop_name,omitempty"`
	// 车牌号码
	EtPlateNumber string `json:"et_plate_number,omitempty" xml:"et_plate_number,omitempty"`
	// 天猫国际官网直供主订单关税税费
	OrderTaxFee string `json:"order_tax_fee,omitempty" xml:"order_tax_fee,omitempty"`
	// 电子凭证服务上门服务的安装地址
	EticketServiceAddr string `json:"eticket_service_addr,omitempty" xml:"eticket_service_addr,omitempty"`
	// 分阶段交易的特权定金订单ID
	O2oEtOrderId string `json:"o2o_et_order_id,omitempty" xml:"o2o_et_order_id,omitempty"`
	// 天猫国际计税优惠金额
	OrderTaxPromotionFee string `json:"order_tax_promotion_fee,omitempty" xml:"order_tax_promotion_fee,omitempty"`
	// 透出前置营销工具
	Pmtp string `json:"pmtp,omitempty" xml:"pmtp,omitempty"`
	// 收货人的所在省份
	ReceiverState string `json:"receiver_state,omitempty" xml:"receiver_state,omitempty"`
	// 收货人的所在城市&lt;br/&gt;注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面&lt;br/&gt;建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 交易中剩余的确认收货金额（这个金额会随着子订单确认收货而不断减少，交易成功后会变为零）。精确到2位小数;单位:元。如:200.07，表示:200 元7分
	AvailableConfirmFee string `json:"available_confirm_fee,omitempty" xml:"available_confirm_fee,omitempty"`
	// 超时到期时间。格式:yyyy-MM-dd HH:mm:ss。业务规则：前提条件：只有在买家已付款，卖家已发货的情况下才有效如果申请了退款，那么超时会落在子订单上；比如说3笔ABC，A申请了，那么返回的是BC的列表, 主定单不存在如果没有申请过退款，那么超时会挂在主定单上；比如ABC，返回主定单，ABC的超时和主定单相同
	TimeoutActionTime string `json:"timeout_action_time,omitempty" xml:"timeout_action_time,omitempty"`
	// 交易快照地址
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔
	TradeFrom string `json:"trade_from,omitempty" xml:"trade_from,omitempty"`
	// 买卡订单本金
	ExpandCardBasicPrice string `json:"expand_card_basic_price,omitempty" xml:"expand_card_basic_price,omitempty"`
	// 买卡订单权益金
	ExpandCardExpandPrice string `json:"expand_card_expand_price,omitempty" xml:"expand_card_expand_price,omitempty"`
	// 用卡订单所用的本金
	ExpandCardBasicPriceUsed string `json:"expand_card_basic_price_used,omitempty" xml:"expand_card_basic_price_used,omitempty"`
	// 用卡订单所用的权益金
	ExpandCardExpandPriceUsed string `json:"expand_card_expand_price_used,omitempty" xml:"expand_card_expand_price_used,omitempty"`
	// 处方药审核状态
	RxAuditStatus string `json:"rx_audit_status,omitempty" xml:"rx_audit_status,omitempty"`
	// CRM系统中特有的ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 同步到卖家库的时间，taobao.trades.sold.incrementv.get接口返回此字段
	AsyncModified string `json:"async_modified,omitempty" xml:"async_modified,omitempty"`
	// 买家下单的地区
	BuyerArea string `json:"buyer_area,omitempty" xml:"buyer_area,omitempty"`
	// 交易外部来源：ownshop(商家官网)
	TradeSource string `json:"trade_source,omitempty" xml:"trade_source,omitempty"`
	// 订单将在此时间前发出，主要用于预售订单
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空
	MarkDesc string `json:"mark_desc,omitempty" xml:"mark_desc,omitempty"`
	// 商品数字编号
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品购买数量。取值范围：大于零的整数,对于一个trade对应多个order的时候（一笔主订单，对应多笔子订单），num=0，num是一个跟商品关联的属性，一笔订单对应多比子订单的时候，主订单上的num无意义。
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 买家使用积分,下单时生成，且一直不变。格式:100;单位:个.
	PointFee int64 `json:"point_fee,omitempty" xml:"point_fee,omitempty"`
	// 买家实际使用积分（扣除部分退款使用的积分），交易完成后生成（交易成功或关闭），交易未完成时该字段值为0。格式:100;单位:个
	RealPointFee int64 `json:"real_point_fee,omitempty" xml:"real_point_fee,omitempty"`
	// 买家获得积分,返点的积分。格式:100;单位:个。返点的积分要交易成功之后才能获得。
	BuyerObtainPointFee int64 `json:"buyer_obtain_point_fee,omitempty" xml:"buyer_obtain_point_fee,omitempty"`
	// 物流到货时效，单位天
	ArriveInterval int64 `json:"arrive_interval,omitempty" xml:"arrive_interval,omitempty"`
	// 物流发货时效，单位小时
	ConsignInterval int64 `json:"consign_interval,omitempty" xml:"consign_interval,omitempty"`
	// 购物金信息输出
	ExpandcardInfo *ExpandCardInfo `json:"expandcard_info,omitempty" xml:"expandcard_info,omitempty"`
	// 使用红包付款金额
	CouponFee int64 `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`
	// 卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5
	SellerFlag int64 `json:"seller_flag,omitempty" xml:"seller_flag,omitempty"`
	// 是否包含邮费。与available_confirm_fee同时使用。可选值:true(包含),false(不包含)
	HasPostFee bool `json:"has_post_fee,omitempty" xml:"has_post_fee,omitempty"`
	// 卖家是否已评价。可选值:true(已评价),false(未评价)
	SellerRate bool `json:"seller_rate,omitempty" xml:"seller_rate,omitempty"`
	// 买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false
	BuyerRate bool `json:"buyer_rate,omitempty" xml:"buyer_rate,omitempty"`
	// 表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。
	IsDaixiao bool `json:"is_daixiao,omitempty" xml:"is_daixiao,omitempty"`
	// 表示订单交易是否网厅订单。 如果该订单是网厅订单，那么该值为true；反之，该值为false。
	IsWt bool `json:"is_wt,omitempty" xml:"is_wt,omitempty"`
	// 买家可以通过此字段查询是否当前交易可以评论，can_rate=true可以评价，false则不能评价。
	CanRate bool `json:"can_rate,omitempty" xml:"can_rate,omitempty"`
	// 卖家是否可以对订单进行评价
	SellerCanRate bool `json:"seller_can_rate,omitempty" xml:"seller_can_rate,omitempty"`
	// 是否是多次发货的订单如果卖家对订单进行多次发货，则为true否则为false
	IsPartConsign bool `json:"is_part_consign,omitempty" xml:"is_part_consign,omitempty"`
	// 判断订单是否有买家留言，有买家留言返回true，否则返回false
	HasBuyerMessage bool `json:"has_buyer_message,omitempty" xml:"has_buyer_message,omitempty"`
}

package fenxiao

// SubOrderDetail 结构体
type SubOrderDetail struct {
	// 分销平台上商品商家编码。
	ItemOuterId string `json:"item_outer_id,omitempty" xml:"item_outer_id,omitempty"`
	// SKU商家编码。
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// SKU属性值。如: 颜色:红色:别名;尺码:L:别名
	SkuProperties string `json:"sku_properties,omitempty" xml:"sku_properties,omitempty"`
	// 快照地址。
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 创建时间。格式 yyyy-MM-dd HH:mm:ss 。
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 交易状态。可选值： <br>WAIT_BUYER_PAY(等待付款)<br> WAIT_SELLER_SEND_GOODS(已付款，待发货）<br> WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)<br> PAID_FORBID_CONSIGN(已付款，禁止发货 ps:只有大快消行业的才有)<br> TRADE_FINISHED(交易成功)<br> TRADE_CLOSED(交易关闭)<br> WAIT_BUYER_CONFIRM_GOODS_ACOUNTED(已付款（已分账），已发货。只对代销分账支持)<br> PAY_ACOUNTED_GOODS_CONFIRM （已分账发货成功）<br> PAY_WAIT_ACOUNT_GOODS_CONFIRM（已付款，确认收货）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 退款金额
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 子采购单id，淘宝交易主键，采购单未付款时为0.（只有支付宝 付款才有这个id，其余付款形式该字段为0）
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 分销平台的子采购单主键
	FenxiaoId int64 `json:"fenxiao_id,omitempty" xml:"fenxiao_id,omitempty"`
	// 分销产品的SKU id。当存在时才会有值，建议使用sku_outer_id，sku_properties这两个值
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// TC子订单ID（经销不显示）
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 分销平台上的产品id，同FenxiaoProduct 的pid
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 代销采购单对应下游200订单状态。可选值：WAIT_SELLER_SEND_GOODS(已付款，待发货)WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)TRADE_CLOSED(已退款成功)TRADE_REFUNDING(退款中)TRADE_FINISHED(交易成功)TRADE_CLOSED_BY_TAOBAO(交易关闭)
	Order200Status string `json:"order_200_status,omitempty" xml:"order_200_status,omitempty"`
	// 分销商店铺中宝贝一口价
	AuctionPrice string `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 产品的采购数量。取值范围:大于零的整数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 采购的产品标题。
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 产品的采购价格。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 分销商应付金额=num(采购数量)*price(采购价)。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 分销商实付金额=total_fee（分销商应付金额）+改价-优惠。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	DistributorPayment string `json:"distributor_payment,omitempty" xml:"distributor_payment,omitempty"`
	// 买家订单上对应的子单零售金额，除以num（数量）后等于最终宝贝的零售价格（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	BuyerPayment string `json:"buyer_payment,omitempty" xml:"buyer_payment,omitempty"`
	// 发票应开金额。根据买家实际付款去除邮费后，按各个子单(商品)金额比例进行分摊后的金额，仅供开发票时做票面金额参考。
	BillFee string `json:"bill_fee,omitempty" xml:"bill_fee,omitempty"`
	// 后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 商品优惠类型：聚划算、秒杀或其他
	TcPreferentialType string `json:"tc_preferential_type,omitempty" xml:"tc_preferential_type,omitempty"`
	// 优惠金额，始终为正数，单位是分，不带小数
	TcDiscountFee int64 `json:"tc_discount_fee,omitempty" xml:"tc_discount_fee,omitempty"`
	// 商品的卖出金额调整，金额增加时为正数，金额减少时为负数，单位是分，不带小数
	TcAdjustFee int64 `json:"tc_adjust_fee,omitempty" xml:"tc_adjust_fee,omitempty"`
	// 优惠活动的折扣金额
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 优惠活动类型0=无优惠1=限时折
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// Feature对象列表目前已有的属性：attr_key为 www，attr_value为1 表示是www子订单；attr_key为 wwwStoreCode，attr_value是www子订单发货的仓库编码；attr_key为 isWt，attr_value为1 表示是网厅子订单；attr_key为wtInfo,attr_value为网厅相关合约信息；attr_key为shipper,attr_value为cn表示菜鸟发货；attr_key为 storeCode，attr_value为仓库信息； attr_key为 erpHold，attr_value为1表示强管控中， attr_value为2表示分单完成；
	Features []FeatureDo `json:"features,omitempty" xml:"features>feature_do,omitempty"`
	// 前台商品SKU ID，不存在时为0。2015年3月15日之前创建的采购单该字段都是0。
	AuctionSkuId int64 `json:"auction_sku_id,omitempty" xml:"auction_sku_id,omitempty"`
	// 前台分销商品的宝贝ID，不存在时为0。2015年4月15日之前创建的采购单该字段都是0。
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// topMemoDTO
	TopMemo *TopMemoDto `json:"top_memo,omitempty" xml:"top_memo,omitempty"`
	// 老的SKU属性值。如: 颜色:红色:别名;尺码:L:别名
	OldSkuProperties string `json:"old_sku_properties,omitempty" xml:"old_sku_properties,omitempty"`
}

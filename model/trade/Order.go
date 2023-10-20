package trade

import (
	"sync"
)

// Order 结构体
type Order struct {
	// 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品的最小库存单位Sku的id.可以通过taobao.item.sku.get获取详细的Sku信息
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 订单状态（请关注此状态，如果为TRADE_CLOSED_BY_TAOBAO状态，则不要对此订单进行发货，切记啊！）。可选值: &lt;ul&gt;&lt;li&gt;TRADE_NO_CREATE_PAY(没有创建支付宝交易) &lt;li&gt;WAIT_BUYER_PAY(等待买家付款) &lt;li&gt;WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) &lt;li&gt;WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) &lt;li&gt;TRADE_BUYER_SIGNED(买家已签收,货到付款专用) &lt;li&gt;TRADE_FINISHED(交易成功) &lt;li&gt;TRADE_CLOSED(付款以后用户退款成功，交易自动关闭) &lt;li&gt;TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)&lt;li&gt;PAY_PENDING(国际信用卡支付付款确认中)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 子订单编号
	Oid string `json:"oid,omitempty" xml:"oid,omitempty"`
	// 应付金额（商品价格 * 商品数量 + 手工调整金额 - 子订单级订单优惠金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 外部网店自己定义的Sku编号
	OuterSkuId string `json:"outer_sku_id,omitempty" xml:"outer_sku_id,omitempty"`
	// 商家外部编码(可与商家外部系统对接)。外部商家自己定义的商品Item的id，可以通过taobao.items.custom.get获取商品的Item的信息
	OuterIid string `json:"outer_iid,omitempty" xml:"outer_iid,omitempty"`
	// 商品图片的绝对路径
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 定制信息
	Customization string `json:"customization,omitempty" xml:"customization,omitempty"`
	// 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 子订单级订单优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 分摊之后的实付金额
	DivideOrderFee string `json:"divide_order_fee,omitempty" xml:"divide_order_fee,omitempty"`
	// 优惠分摊
	PartMjzDiscount string `json:"part_mjz_discount,omitempty" xml:"part_mjz_discount,omitempty"`
	// 退款状态。退款状态。可选值 WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 手工调整金额.格式为:1.01;单位:元;精确到小数点后两位.
	AdjustFee string `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// SKU的值。如：机身颜色:黑色;手机套餐:官方标配
	SkuPropertiesName string `json:"sku_properties_name,omitempty" xml:"sku_properties_name,omitempty"`
	// 套餐的值。如：M8原装电池:便携支架:M8专用座充:莫凡保护袋
	ItemMealName string `json:"item_meal_name,omitempty" xml:"item_meal_name,omitempty"`
	// 卖家类型，可选值为：B（商城商家），C（普通卖家）
	SellerType string `json:"seller_type,omitempty" xml:"seller_type,omitempty"`
	// 对应门票有效期的外部id
	TicketOuterId string `json:"ticket_outer_id,omitempty" xml:"ticket_outer_id,omitempty"`
	// 门票有效期的key
	TicketExpdateKey string `json:"ticket_expdate_key,omitempty" xml:"ticket_expdate_key,omitempty"`
	// 发货的仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 支持家装类物流的类型
	TmserSpuCode string `json:"tmser_spu_code,omitempty" xml:"tmser_spu_code,omitempty"`
	// 天猫汽车服务预约时间
	EtSerTime string `json:"et_ser_time,omitempty" xml:"et_ser_time,omitempty"`
	// 电子凭证预约门店地址
	EtShopName string `json:"et_shop_name,omitempty" xml:"et_shop_name,omitempty"`
	// 电子凭证核销门店地址
	EtVerifiedShopName string `json:"et_verified_shop_name,omitempty" xml:"et_verified_shop_name,omitempty"`
	// 车牌号码
	EtPlateNumber string `json:"et_plate_number,omitempty" xml:"et_plate_number,omitempty"`
	// 天猫国际子订单关税税费
	SubOrderTaxFee string `json:"sub_order_tax_fee,omitempty" xml:"sub_order_tax_fee,omitempty"`
	// 天猫国际子订单关税税率
	SubOrderTaxRate string `json:"sub_order_tax_rate,omitempty" xml:"sub_order_tax_rate,omitempty"`
	// 征集预售订单征集状态：1（征集中），2（征集成功），3（征集失败）
	ZhengjiStatus string `json:"zhengji_status,omitempty" xml:"zhengji_status,omitempty"`
	// 天猫国际子订单计税优惠金额
	SubOrderTaxPromotionFee string `json:"sub_order_tax_promotion_fee,omitempty" xml:"sub_order_tax_promotion_fee,omitempty"`
	// 天猫国际子订单包税金额
	TaxCouponDiscount string `json:"tax_coupon_discount,omitempty" xml:"tax_coupon_discount,omitempty"`
	// 特殊退款状态
	SpecialRefundType string `json:"special_refund_type,omitempty" xml:"special_refund_type,omitempty"`
	// 订单快照URL
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 订单超时到期时间。格式:yyyy-MM-dd HH:mm:ss
	TimeoutActionTime string `json:"timeout_action_time,omitempty" xml:"timeout_action_time,omitempty"`
	// 子订单的交易结束时间说明：子订单有单独的结束时间，与主订单的结束时间可能有所不同，在有退款发起的时候或者是主订单分阶段付款的时候，子订单的结束时间会早于主订单的结束时间，所以开放这个字段便于订单结束状态的判断
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 子订单来源,如jhs(聚划算)、taobao(淘宝)、wap(无线)
	OrderFrom string `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 子订单发货时间，当卖家对订单进行了多次发货，子订单的发货时间和主订单的发货时间可能不一样了，那么就需要以子订单的时间为准。（没有进行多次发货的订单，主订单的发货时间和子订单的发货时间都一样）
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 子订单的运送方式（卖家对订单进行多次发货之后，一个主订单下的子订单的运送方式可能不同，用order.shipping_type来区分子订单的运送方式）
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// 子订单发货的快递公司名称
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 子订单所在包裹的运单号
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// bind_oid字段的升级，支持返回绑定的多个子订单，多个子订单以半角逗号分隔
	BindOids string `json:"bind_oids,omitempty" xml:"bind_oids,omitempty"`
	// 套餐ID
	ItemMealId string `json:"item_meal_id,omitempty" xml:"item_meal_id,omitempty"`
	// 购物金核销子订单权益金分摊金额（单位为元）
	ExpandCardExpandPriceUsedSuborder string `json:"expand_card_expand_price_used_suborder,omitempty" xml:"expand_card_expand_price_used_suborder,omitempty"`
	// 购物金核销子订单本金分摊金额（单位为元）
	ExpandCardBasicPriceUsedSuborder string `json:"expand_card_basic_price_used_suborder,omitempty" xml:"expand_card_basic_price_used_suborder,omitempty"`
	// 是否为闲鱼订单 1是0否
	IsIdle string `json:"is_idle,omitempty" xml:"is_idle,omitempty"`
	// 商品数字ID
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 购买数量。取值范围:大于零的整数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 最近退款的id
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 交易商品对应的类目ID
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 买家是否已评价。可选值：true(已评价)，false(未评价)
	BuyerRate bool `json:"buyer_rate,omitempty" xml:"buyer_rate,omitempty"`
	// 卖家是否已评价。可选值：true(已评价)，false(未评价)
	SellerRate bool `json:"seller_rate,omitempty" xml:"seller_rate,omitempty"`
	// 表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。
	IsDaixiao bool `json:"is_daixiao,omitempty" xml:"is_daixiao,omitempty"`
	// 子订单是否是www订单
	IsWww bool `json:"is_www,omitempty" xml:"is_www,omitempty"`
	// 天猫国际子订单是否包税
	TaxFree bool `json:"tax_free,omitempty" xml:"tax_free,omitempty"`
	// 子订单是否优惠贬值
	IsDevalueFee bool `json:"is_devalue_fee,omitempty" xml:"is_devalue_fee,omitempty"`
	// 是否超卖
	IsOversold bool `json:"is_oversold,omitempty" xml:"is_oversold,omitempty"`
}

var poolOrder = sync.Pool{
	New: func() any {
		return new(Order)
	},
}

// GetOrder() 从对象池中获取Order
func GetOrder() *Order {
	return poolOrder.Get().(*Order)
}

// ReleaseOrder 释放Order
func ReleaseOrder(v *Order) {
	v.Price = ""
	v.SkuId = ""
	v.Status = ""
	v.Oid = ""
	v.TotalFee = ""
	v.OuterSkuId = ""
	v.OuterIid = ""
	v.PicPath = ""
	v.Title = ""
	v.Customization = ""
	v.Payment = ""
	v.DiscountFee = ""
	v.DivideOrderFee = ""
	v.PartMjzDiscount = ""
	v.RefundStatus = ""
	v.AdjustFee = ""
	v.SkuPropertiesName = ""
	v.ItemMealName = ""
	v.SellerType = ""
	v.TicketOuterId = ""
	v.TicketExpdateKey = ""
	v.StoreCode = ""
	v.TmserSpuCode = ""
	v.EtSerTime = ""
	v.EtShopName = ""
	v.EtVerifiedShopName = ""
	v.EtPlateNumber = ""
	v.SubOrderTaxFee = ""
	v.SubOrderTaxRate = ""
	v.ZhengjiStatus = ""
	v.SubOrderTaxPromotionFee = ""
	v.TaxCouponDiscount = ""
	v.SpecialRefundType = ""
	v.SnapshotUrl = ""
	v.TimeoutActionTime = ""
	v.EndTime = ""
	v.OrderFrom = ""
	v.ConsignTime = ""
	v.ShippingType = ""
	v.LogisticsCompany = ""
	v.InvoiceNo = ""
	v.BindOids = ""
	v.ItemMealId = ""
	v.ExpandCardExpandPriceUsedSuborder = ""
	v.ExpandCardBasicPriceUsedSuborder = ""
	v.IsIdle = ""
	v.NumIid = 0
	v.Num = 0
	v.RefundId = 0
	v.Cid = 0
	v.BuyerRate = false
	v.SellerRate = false
	v.IsDaixiao = false
	v.IsWww = false
	v.TaxFree = false
	v.IsDevalueFee = false
	v.IsOversold = false
	poolOrder.Put(v)
}

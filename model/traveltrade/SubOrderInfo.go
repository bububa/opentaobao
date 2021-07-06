package traveltrade

// SubOrderInfo 结构体
type SubOrderInfo struct {
	// 出行人信息
	Travellers []TravellerInfo `json:"travellers,omitempty" xml:"travellers>traveller_info,omitempty"`
	// 子订单交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 子订单超时自动关单（卖家未能及时发货）时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 子订单扩展属性。KV对结构。每个KV对间使用分号隔开。如：k1:v1;k2:v2
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 退款状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意);WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货);WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货);SELLER_REFUSE_BUYER(卖家拒绝退款);CLOSED(退款关闭);SUCCESS(退款成功)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 子订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// string类型子订单id
	SubOrderIdString string `json:"sub_order_id_string,omitempty" xml:"sub_order_id_string,omitempty"`
	// 订单处理时间 精确到秒
	ProcessTime string `json:"process_time,omitempty" xml:"process_time,omitempty"`
	// 卖家手工调整金额，单位：分
	AdjustFee int64 `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// 订单所属业务类型。0-未知，1-度假（自由行，跟团游），2-普通签证，3-门票，4-wifi，7-当地玩乐，9-邮轮，10-用车，12-电话卡，17-流量充值，18-港澳签注，20-旅行租赁
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 购买的商品信息（商品、sku信息，单价、数量、出行日期等信息）
	BuyItemInfo *BuyItemInfo `json:"buy_item_info,omitempty" xml:"buy_item_info,omitempty"`
	// 联系人/取件人信息。通用联系人信息，没有配置度假出行人模板的商品订单（如wifi电话卡等类目），从这里获取联系人或取件人信息
	Contactor *TravellerInfo `json:"contactor,omitempty" xml:"contactor,omitempty"`
	// 子订单优惠金额（如打折，VIP，满就送等），不包含带资优惠（如红包、旅行券）。单位：分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 订单类型。1、普通订单，2、二次确认订单，3、二次预约订单
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 买家实付金额。单位：分
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 子订单退款编号。如果子订单没有发生退款，则该值为空
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 应付金额（商品价格 * 商品数量 + 手工调整金额 - 子订单优惠金额），大于等于payment，因为可能包含了带资优惠（如红包、旅行券）。单位：分
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 二次确认状态，1、待商家确认 2、确认不通过 3、确认通过 4、过期未确认 5、买家申请退款成功 6、票号验证
	ConfirmStatus int64 `json:"confirm_status,omitempty" xml:"confirm_status,omitempty"`
}

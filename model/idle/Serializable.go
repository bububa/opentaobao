package idle

import (
	"sync"
)

// Serializable 结构体
type Serializable struct {
	// 明细
	PayDetails []SubPayBillDto `json:"pay_details,omitempty" xml:"pay_details>sub_pay_bill_dto,omitempty"`
	// 代扣计划列表
	PlanIds []string `json:"plan_ids,omitempty" xml:"plan_ids>string,omitempty"`
	// 商户订单号（唯一建）
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 代扣计划
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// 支付宝流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// INIT:初始状态;PARTIAL_SUCCESSED:支付部分成功;SUCCESSED:支付成功;FAILED:支付失败
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 业务编码：R1:回收
	PayBizCode string `json:"pay_biz_code,omitempty" xml:"pay_biz_code,omitempty"`
	// 状态描述
	PayStatusDesc string `json:"pay_status_desc,omitempty" xml:"pay_status_desc,omitempty"`
	// 回收商appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 买家nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 卖家nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家支付宝id,用于回收商支付宝打款
	SellerAlipayUserId string `json:"seller_alipay_user_id,omitempty" xml:"seller_alipay_user_id,omitempty"`
	// 卖家支付宝账号,已脱敏
	SellerAlipayAccount string `json:"seller_alipay_account,omitempty" xml:"seller_alipay_account,omitempty"`
	// 卖家收货地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 卖家手机号码
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 取件类型 1：顺风 2：上门取件
	ShipType string `json:"ship_type,omitempty" xml:"ship_type,omitempty"`
	// 取件时间
	ShipTime string `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
	// 卖家姓名
	SellerRealName string `json:"seller_real_name,omitempty" xml:"seller_real_name,omitempty"`
	// 支付宝签约code
	ZfbDkCode string `json:"zfb_dk_code,omitempty" xml:"zfb_dk_code,omitempty"`
	// 估价id
	ApprizeId string `json:"apprize_id,omitempty" xml:"apprize_id,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 村
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 芝麻分
	ZmScore string `json:"zm_score,omitempty" xml:"zm_score,omitempty"`
	// 未来的卡券id
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 加价券金额
	CouponFee string `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`
	// 一、按百分比 {     &#34;actType&#34;: 1,     &#34;beg&#34;: &#34;2017-12-01 12:00:00&#34;,     &#34;couponFee&#34;: 4400,     &#34;couponId&#34;: &#34;-2&#34;,     &#34;desc&#34;: &#34;满50加价10%&#34;,     &#34;end&#34;: &#34;2017-12-13 12:00:00&#34;,     &#34;idleBizCode&#34;: &#34;9&#34;,     &#34;low&#34;: 5000,     &#34;percent&#34;: 0.1 }  二、按固定加价 {     &#34;actType&#34;: 2,     &#34;beg&#34;: &#34;2017-12-01 12:00:00&#34;,     &#34;couponFee&#34;: 3000,     &#34;couponId&#34;: &#34;-3&#34;,     &#34;desc&#34;: &#34;满50固定加价30&#34;,     &#34;end&#34;: &#34;2017-12-13 12:00:00&#34;,     &#34;idleBizCode&#34;: &#34;9&#34;,     &#34;low&#34;: 5000,     &#34;addFee&#34;: 3000 }  三、阶梯价 {     &#34;actType&#34;: 3,     &#34;beg&#34;: &#34;2017-12-21 16:00:00&#34;,     &#34;couponFee&#34;: 0,     &#34;couponId&#34;: &#34;-5&#34;,     &#34;desc&#34;: &#34;阶梯价&#34;,     &#34;end&#34;: &#34;2019-02-31 00:00:00&#34;,     &#34;idleBizCode&#34;: &#34;8&#34;,     &#34;low&#34;: 0,     &#34;range&#34;: [{         &#34;actType&#34;: 1,         &#34;couponFee&#34;: 0,         &#34;desc&#34;: &#34;100~2000元,加价10%&#34;,         &#34;low&#34;: 10000,         &#34;percent&#34;: 0.1,         &#34;up&#34;: 200000     },     {         &#34;actType&#34;: 1,         &#34;couponFee&#34;: 0,         &#34;desc&#34;: &#34;2000~3000元,加价20%&#34;,         &#34;low&#34;: 200000,         &#34;percent&#34;: 0.2,         &#34;up&#34;: 300000     },     {         &#34;actType&#34;: 2,         &#34;addFee&#34;: 70000,         &#34;couponFee&#34;: 0,         &#34;desc&#34;: &#34;3000元以上,加价700&#34;,         &#34;low&#34;: 300000,         &#34;up&#34;: 0     }] }
	CouponRule string `json:"coupon_rule,omitempty" xml:"coupon_rule,omitempty"`
	// 回收渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 1 好评,0 默认中评
	RateGrade string `json:"rate_grade,omitempty" xml:"rate_grade,omitempty"`
	// 评价内容
	RateContent string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
	// 卖家关闭订单原因
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 买家关闭订单原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 卖家申请退回原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// onlien:线上环境  pre：测试环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 600以下为Z5，600-649为Z4，650-699为Z3，700-749为Z2，750及以上为Z1
	ZmLevel string `json:"zm_level,omitempty" xml:"zm_level,omitempty"`
	// 1:支付宝现金 2：天猫红包
	IdlePayType string `json:"idle_pay_type,omitempty" xml:"idle_pay_type,omitempty"`
	// 渠道内的业务数据json格式 比如 ship=1 服装类的不需要发货,weight=5-15 代表服装5-15kg, userLevel=vip 代表免议价用户,sellerRealPhone 淘宝账号绑定的手机号
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// spuId
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// xydk:闲鱼代扣，信用订单使用
	DkType string `json:"dk_type,omitempty" xml:"dk_type,omitempty"`
	// 回收订单对应的快递单号
	ShipMailNo string `json:"ship_mail_no,omitempty" xml:"ship_mail_no,omitempty"`
	// 多笔估价场景，订单纬度对应的多笔估价集合
	ApprizeIdList string `json:"apprize_id_list,omitempty" xml:"apprize_id_list,omitempty"`
	// 多笔估价场景，订单纬度对应的多笔spuId集合
	SpuIdList string `json:"spu_id_list,omitempty" xml:"spu_id_list,omitempty"`
	// 返现金额，单位分
	CashFee string `json:"cash_fee,omitempty" xml:"cash_fee,omitempty"`
	// 估计版本，第几次估价
	QuoteVersion string `json:"quote_version,omitempty" xml:"quote_version,omitempty"`
	// 回收场景
	SceneType string `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	// 取件截止时间
	ShipTimeEnd string `json:"ship_time_end,omitempty" xml:"ship_time_end,omitempty"`
	// 当前代扣状态：104代扣失败、105代扣成功、106代扣逾期
	DkStatus string `json:"dk_status,omitempty" xml:"dk_status,omitempty"`
	// 卖家地址-乡镇/街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 支付宝交易号
	AlipayOrderId string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 外部业务ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 申请扣款金额（单位分）
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 实际支付金额（单位分）
	ReceiptAmount int64 `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 回收订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 买家id,已脱敏处理返回0
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 卖家id,已脱敏处理返回0
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 信用预付金额,单位分
	CreditPayAmount int64 `json:"credit_pay_amount,omitempty" xml:"credit_pay_amount,omitempty"`
	// 估价金额,单位分
	ApprizeAmount int64 `json:"apprize_amount,omitempty" xml:"apprize_amount,omitempty"`
	// 是否信用预付订单
	CreditPay bool `json:"credit_pay,omitempty" xml:"credit_pay,omitempty"`
	// 是否支付宝签约
	ZfbDk bool `json:"zfb_dk,omitempty" xml:"zfb_dk,omitempty"`
	// 是否为返现订单
	CashOrder bool `json:"cash_order,omitempty" xml:"cash_order,omitempty"`
}

var poolSerializable = sync.Pool{
	New: func() any {
		return new(Serializable)
	},
}

// GetSerializable() 从对象池中获取Serializable
func GetSerializable() *Serializable {
	return poolSerializable.Get().(*Serializable)
}

// ReleaseSerializable 释放Serializable
func ReleaseSerializable(v *Serializable) {
	v.PayDetails = v.PayDetails[:0]
	v.PlanIds = v.PlanIds[:0]
	v.BizOrderId = ""
	v.PlanId = ""
	v.AlipayTradeNo = ""
	v.PayTime = ""
	v.PayStatus = ""
	v.PayBizCode = ""
	v.PayStatusDesc = ""
	v.AppKey = ""
	v.BuyerNick = ""
	v.SellerNick = ""
	v.SellerAlipayUserId = ""
	v.SellerAlipayAccount = ""
	v.SellerAddress = ""
	v.SellerPhone = ""
	v.ShipType = ""
	v.ShipTime = ""
	v.SellerRealName = ""
	v.ZfbDkCode = ""
	v.ApprizeId = ""
	v.City = ""
	v.Province = ""
	v.Area = ""
	v.Country = ""
	v.ZmScore = ""
	v.CouponId = ""
	v.CouponFee = ""
	v.CouponRule = ""
	v.Channel = ""
	v.RateGrade = ""
	v.RateContent = ""
	v.CloseReason = ""
	v.BuyerCloseReason = ""
	v.RefundReason = ""
	v.Env = ""
	v.GmtCreate = ""
	v.ZmLevel = ""
	v.IdlePayType = ""
	v.ChannelData = ""
	v.SpuId = ""
	v.DkType = ""
	v.ShipMailNo = ""
	v.ApprizeIdList = ""
	v.SpuIdList = ""
	v.CashFee = ""
	v.QuoteVersion = ""
	v.SceneType = ""
	v.ShipTimeEnd = ""
	v.DkStatus = ""
	v.Town = ""
	v.AlipayOrderId = ""
	v.Amount = ""
	v.OuterId = ""
	v.TotalAmount = 0
	v.ReceiptAmount = 0
	v.OrderStatus = 0
	v.BuyerId = 0
	v.SellerId = 0
	v.CreditPayAmount = 0
	v.ApprizeAmount = 0
	v.CreditPay = false
	v.ZfbDk = false
	v.CashOrder = false
	poolSerializable.Put(v)
}

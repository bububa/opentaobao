package fenxiao

import (
	"sync"
)

// OrderList 结构体
type OrderList struct {
	// 子采购单详情
	SubOrderDetailYphList []SubOrderList `json:"sub_order_detail_yph_list,omitempty" xml:"sub_order_detail_yph_list>sub_order_list,omitempty"`
	// 建议废弃 Feature对象列表目前已有的属性： attr_key为 www，attr_value为1 表示是www子订单； attr_key为 wwwStoreCode，attr_value是www子订单发货的仓库编码； attr_key为 isWt，attr_value为1 表示是网厅子订单； attr_key为wtInfo,attr_value为网厅相关合约信息； attr_key为shipper,attr_value为cn表示菜鸟发货； attr_key为 storeCode，attr_value为仓库信息； attr_key为 erpHold，attr_value为1表示强管控中， attr_value为2表示分单完成；
	Features []FeatureDo `json:"features,omitempty" xml:"features>feature_do,omitempty"`
	// 支付宝交易号，在分销商使用担保交易支付时存在。 来源主订单的payOrderId。
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 运单号
	LogisticsId string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
	// 采购单付款时间。格式:yyyy-MM-dd HH:mm:ss
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 采购单留言
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 采购单主单买家实付金额，当前台订单只有一个供应商时该字段有值。单位：元。
	BuyerActualPayYuan string `json:"buyer_actual_pay_yuan,omitempty" xml:"buyer_actual_pay_yuan,omitempty"`
	// 发货时间，格式：格式:yyyy-MM-dd HH:mm:ss
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 物流公司
	LogisticsCompanyName string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
	// 分销商昵称
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 采购单邮费
	PostageYuan string `json:"postage_yuan,omitempty" xml:"postage_yuan,omitempty"`
	// 供应商备注信息。 只有供应商身份查询输出此新信息
	SupplierMemo string `json:"supplier_memo,omitempty" xml:"supplier_memo,omitempty"`
	// 采购单编号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 采购单最后修改时间（包括订单状态变更和订单操作）。格式:yyyy-MM-dd HH:mm:ss
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 分销商实付金额，单位：元
	DistributorPayPriceYuan string `json:"distributor_pay_price_yuan,omitempty" xml:"distributor_pay_price_yuan,omitempty"`
	// 供应商昵称
	SupplierNick string `json:"supplier_nick,omitempty" xml:"supplier_nick,omitempty"`
	// 确认收货金额，单位：元
	ConfirmPaidFeeYuan string `json:"confirm_paid_fee_yuan,omitempty" xml:"confirm_paid_fee_yuan,omitempty"`
	// 采购单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 分销商来源网站
	DistributorFromSys string `json:"distributor_from_sys,omitempty" xml:"distributor_from_sys,omitempty"`
	// 交易成功的时间，分销商确认收货/付款的时间。格式:yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 采购单总额（不含邮费），单位：元
	TotalPriceYuan string `json:"total_price_yuan,omitempty" xml:"total_price_yuan,omitempty"`
	// 主采购单交易状态，枚举：WAIT_SELLER_SEND_GOODS-已付款，待发货；WAIT_SELLER_CONFIRM_PAY-待确认收款；TRADE_REFUNDING-退款中；WAIT_BUYER_PAY-等待买家付款；WAIT_BUYER_CONFIRM_GOODS-已付款，已发货；TRADE_FINISHED-交易成功了；TRADE_CLOSED-交易关闭；TRADE_FOR_PAY-已付款；TRADE_REFUNDED-已退款
	OrderStatusCode string `json:"order_status_code,omitempty" xml:"order_status_code,omitempty"`
	// 主采购单退款状态，枚举：ORDER_RF_NO_REFUND-没有过申请退款；ORDER_RF_REFUNDING-退款活动中,至少有一个子单在退款中；ORDER_RF_END_REFUND-退款结束，即曾经发生过退款且所有退款都完结了
	RefundStatusCode string `json:"refund_status_code,omitempty" xml:"refund_status_code,omitempty"`
	// 寄售模式文字（商家仓，菜鸟仓，门店发货）
	LogisType string `json:"logis_type,omitempty" xml:"logis_type,omitempty"`
	// tp单创单时间
	TpCreateTime string `json:"tp_create_time,omitempty" xml:"tp_create_time,omitempty"`
	// 支付方式，枚举：101-支付宝担保交易；102-代销分账；103-即时到账；104-预存款；201-线下转账汇款交易；301-资金账户预存款交易；302-资金账户即时到账交易；303-先款后货；401-账期支付；402-国际分账；403-直营无需支付；404-外币分账；10000-无需支付
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 配送方式，枚举：1-平邮；2-快递；3-ems；4-虚拟物品
	Shipping int64 `json:"shipping,omitempty" xml:"shipping,omitempty"`
	// 采购单总额（不含邮费，单位：分）
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 消费者主订单ID （经销模式下，不存在该单号)
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 主采购单交易状态-数字表示，枚举：1-已付款，待发货；2-待确认收款；4-等待买家付款；5-已付款，已发货；6-交易成功；7-交易关闭；8-已付款；
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 确认收货金额，单位：分
	ConfirmPaidFee int64 `json:"confirm_paid_fee,omitempty" xml:"confirm_paid_fee,omitempty"`
	// 查询的经营模式，当不指定时，默认查询代销订单 代销：1 经销：2 寄售（自营寄售）：5 平台寄售：6
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场；302-猫超一盘货；506-虾选一盘货
	ChannelCode int64 `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 采购单主单买家实付金额，当前台订单只有一个供应商时该字段有值。单位：分。
	BuyerActualPay int64 `json:"buyer_actual_pay,omitempty" xml:"buyer_actual_pay,omitempty"`
	// 分销采购单号（分销流水号）
	FenxiaoId int64 `json:"fenxiao_id,omitempty" xml:"fenxiao_id,omitempty"`
	// 采购单邮费，单位：分
	Postage int64 `json:"postage,omitempty" xml:"postage,omitempty"`
	// 供应商来源网站。 入驻时定义供应商来源的外部系统编码, values: taobao, alibaba
	SupplierFromSys int64 `json:"supplier_from_sys,omitempty" xml:"supplier_from_sys,omitempty"`
	// 分销商实付金额，单位：分
	DistributorPayPrice int64 `json:"distributor_pay_price,omitempty" xml:"distributor_pay_price,omitempty"`
	// 主采购单退款状态-数字表示，枚举：9-没有过申请退款；10-退款活动中,至少有一个子单在退款中；11-退款结束，即曾经发生过退款且所有退款都完结了
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 寄售模式code(0-菜鸟仓，1-商家仓，2-门店发货)
	LogisTypeCode int64 `json:"logis_type_code,omitempty" xml:"logis_type_code,omitempty"`
	// 消费者详细信息（收货人姓名、收货人手机、收货人电话、收货人邮编、收货人省份、收货人详细地址）
	DpBuyerDetail *DpBuyerDetail `json:"dp_buyer_detail,omitempty" xml:"dp_buyer_detail,omitempty"`
}

var poolOrderList = sync.Pool{
	New: func() any {
		return new(OrderList)
	},
}

// GetOrderList() 从对象池中获取OrderList
func GetOrderList() *OrderList {
	return poolOrderList.Get().(*OrderList)
}

// ReleaseOrderList 释放OrderList
func ReleaseOrderList(v *OrderList) {
	v.SubOrderDetailYphList = v.SubOrderDetailYphList[:0]
	v.Features = v.Features[:0]
	v.AlipayNo = ""
	v.LogisticsId = ""
	v.PayTime = ""
	v.Memo = ""
	v.BuyerActualPayYuan = ""
	v.ConsignTime = ""
	v.LogisticsCompanyName = ""
	v.DistributorNick = ""
	v.PostageYuan = ""
	v.SupplierMemo = ""
	v.BizOrderId = ""
	v.GmtModified = ""
	v.DistributorPayPriceYuan = ""
	v.SupplierNick = ""
	v.ConfirmPaidFeeYuan = ""
	v.GmtCreate = ""
	v.DistributorFromSys = ""
	v.EndTime = ""
	v.TotalPriceYuan = ""
	v.OrderStatusCode = ""
	v.RefundStatusCode = ""
	v.LogisType = ""
	v.TpCreateTime = ""
	v.PayType = 0
	v.Shipping = 0
	v.TotalPrice = 0
	v.TcOrderId = 0
	v.OrderStatus = 0
	v.ConfirmPaidFee = 0
	v.TradeType = 0
	v.ChannelCode = 0
	v.BuyerActualPay = 0
	v.FenxiaoId = 0
	v.Postage = 0
	v.SupplierFromSys = 0
	v.DistributorPayPrice = 0
	v.RefundStatus = 0
	v.LogisTypeCode = 0
	v.DpBuyerDetail = nil
	poolOrderList.Put(v)
}

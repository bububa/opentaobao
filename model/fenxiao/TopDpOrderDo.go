package fenxiao

import (
	"sync"
)

// TopDpOrderDo 结构体
type TopDpOrderDo struct {
	// ISV自定义key，通过taobao.fenxiao.order.customfield.update 写入。目前禁用
	IsvCustomKey []string `json:"isv_custom_key,omitempty" xml:"isv_custom_key>string,omitempty"`
	// ISV自定义值，通过taobao.fenxiao.order.customfield.update 写入。目前禁用
	IsvCustomValue []string `json:"isv_custom_value,omitempty" xml:"isv_custom_value>string,omitempty"`
	// 建议废弃 Feature对象列表目前已有的属性： attr_key为 www，attr_value为1 表示是www子订单； attr_key为 wwwStoreCode，attr_value是www子订单发货的仓库编码； attr_key为 isWt，attr_value为1 表示是网厅子订单； attr_key为wtInfo,attr_value为网厅相关合约信息； attr_key为 storeCode，attr_value为仓库信息；  attr_key为 erpHold，attr_value为1表示强管控中， attr_value为2表示分单完成；
	Features []FeatureDo `json:"features,omitempty" xml:"features>feature_do,omitempty"`
	// 子订单的详细信息列表
	SubPurchaseOrders []SubOrderDetail `json:"sub_purchase_orders,omitempty" xml:"sub_purchase_orders>sub_order_detail,omitempty"`
	// 采购单留言列表，如果是代销，包含消费者留言。
	OrderMessages []OrderMessages `json:"order_messages,omitempty" xml:"order_messages>order_messages,omitempty"`
	// [架海金梁独有字段，非架海金梁用户请勿关心]子单物流发货信息
	LogisticsInfos []ErpLogisticsInfo `json:"logistics_infos,omitempty" xml:"logistics_infos>erp_logistics_info,omitempty"`
	// 供应商来源网站。 入驻时定义供应商来源的外部系统编码, values: taobao, alibaba
	SupplierFrom string `json:"supplier_from,omitempty" xml:"supplier_from,omitempty"`
	// 供应商昵称（在来源网站的帐号名）
	SupplierUsername string `json:"supplier_username,omitempty" xml:"supplier_username,omitempty"`
	// 建议废弃： 分销商来源网站。 入驻时定义分销商来源的外部系统编码, 目前固定值: taobao
	DistributorFrom string `json:"distributor_from,omitempty" xml:"distributor_from,omitempty"`
	// 分销商昵称（在来源网站的帐号名）。
	DistributorUsername string `json:"distributor_username,omitempty" xml:"distributor_username,omitempty"`
	// 加密后的消费者淘宝ID，长度为32
	BuyerTaobaoId string `json:"buyer_taobao_id,omitempty" xml:"buyer_taobao_id,omitempty"`
	// 消费者（买家）nick，供应商查询不会返回买家昵称，分销商查询才会返回。
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 交易模式（分销方式）：AGENT（代销）、DEALER（经销）、INSTEAD_SALE（代售）、CONSIGNMENT（自营寄售）、PLATFORM_CONSIGNMENT（平台寄售）
	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 采购单创建时间。格式:yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 采购单留言，默认是分销商留言。在代销模式下信息包括消费者（买家）留言和分销商留言信息
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 供应商备注信息。 只有供应商身份查询输出此新信息
	SupplierMemo string `json:"supplier_memo,omitempty" xml:"supplier_memo,omitempty"`
	// 支付宝交易号，在分销商使用担保交易支付时存在。 来源主订单的payOrderId。
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 应付采购单总额（不含邮费,精确到2位小数;单位:元。如:200.07，表示:200元7分 ) 计算公式： SUM(子单采购价 * 子单采购数量 - 子单折扣）+ 主单调整金额 - 主单折扣
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 采购单邮费。(改价后最新邮费，精确到2位小数;单位:元。如:8.07，表示:8元7分 )
	PostFee string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 分销商实付金额。(精确到2位小数;单位:元。如:208.14，表示:208元1角4分 )
	DistributorPayment string `json:"distributor_payment,omitempty" xml:"distributor_payment,omitempty"`
	// 消费者（买家）支付给分销商的主单的实际总金额。 注意买家购买的商品可能不是全部来自同一供货商，所以此金额不代表供应商销售商品消费者的实付款。请同时参考子单上的相关金额。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	BuyerPayment string `json:"buyer_payment,omitempty" xml:"buyer_payment,omitempty"`
	// 交易订单快照URL（该字段废弃）
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 采购单交易状态。 可选值： WAIT_BUYER_PAY(等待付款) WAIT_SELLER_SEND_GOODS(已付款，待发货） WAIT_BUYER_CONFIRM_GOODS(已付款，已发货) PAID_FORBID_CONSIGN(已付款，禁止发货 ps:只有大快消行业的才有) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) WAIT_BUYER_CONFIRM_GOODS_ACOUNTED(已付款（已分账），已发货。只对代销分账支持) PAY_ACOUNTED_GOODS_CONFIRM （已分账发货成功） PAY_WAIT_ACOUNT_GOODS_CONFIRM（已付款，确认收货）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 采购单付款时间。格式:yyyy-MM-dd HH:mm:ss
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 交易结束时间。 如果交易成功此时间表示交易成功时间，如果交易关闭，此时间表示交易关闭时间。
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 采购单物流发货时间。格式:yyyy-MM-dd HH:mm:ss
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 采购单最后修改时间（包括订单状态变更和订单操作）。格式:yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 支付方式：ALIPAY_SURETY（支付宝担保交易）、ALIPAY_CHAIN（分账交易）、TRANSFER（线下转账）、PREPAY（预存款）、IMMEDIATELY（即时到账）、CASHGOODS（先款后货）、ACCOUNT_PERIOD（账期支付）
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 物流配送方式，发货后有此数据。 对应的值：FAST(快速)、EMS、ORDINARY(平邮)、SELLER(卖家包邮)
	Shipping string `json:"shipping,omitempty" xml:"shipping,omitempty"`
	// 物流公司名称
	LogisticsCompanyName string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
	// 运单号
	LogisticsId string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
	// 支付宝交易号，担保交易使用。
	AlipayOrderNo string `json:"alipay_order_no,omitempty" xml:"alipay_order_no,omitempty"`
	// 已执行确认收货的金额，单位：元
	ConfirmPaidFeeYuan string `json:"confirm_paid_fee_yuan,omitempty" xml:"confirm_paid_fee_yuan,omitempty"`
	// 消费者淘宝id被加密后的唯一标识
	OpenBuyerUid string `json:"open_buyer_uid,omitempty" xml:"open_buyer_uid,omitempty"`
	// 计划出库时间
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 计划送达时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 配送服务,201为驿站送货上门服务,202为顺丰配送服务,里面多个值时用英文逗号隔开
	AsdpAds string `json:"asdp_ads,omitempty" xml:"asdp_ads,omitempty"`
	// 推荐配送公司编码
	DeliveryCps string `json:"delivery_cps,omitempty" xml:"delivery_cps,omitempty"`
	// 消费者主订单ID （经销模式下，不存在该单号）
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 废弃：如果是担保交易此值是TP800的交易单号，其他情况是采购单单号
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 分销采购单号（分销流水号）
	FenxiaoId int64 `json:"fenxiao_id,omitempty" xml:"fenxiao_id,omitempty"`
	// 供应商备注旗帜vlaue在1-5之间。 非1-5之间，都采用1作为默认。 1:红色 2:黄色 3:绿色 4:蓝色 5:粉红色 供应商角色查询返回
	SupplierFlag int64 `json:"supplier_flag,omitempty" xml:"supplier_flag,omitempty"`
	// 买家详细信息
	Receiver *TopReceiverDo `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 渠道(市场)编码，例如消费电子的编码为200002
	ChannelCode int64 `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}

var poolTopDpOrderDo = sync.Pool{
	New: func() any {
		return new(TopDpOrderDo)
	},
}

// GetTopDpOrderDo() 从对象池中获取TopDpOrderDo
func GetTopDpOrderDo() *TopDpOrderDo {
	return poolTopDpOrderDo.Get().(*TopDpOrderDo)
}

// ReleaseTopDpOrderDo 释放TopDpOrderDo
func ReleaseTopDpOrderDo(v *TopDpOrderDo) {
	v.IsvCustomKey = v.IsvCustomKey[:0]
	v.IsvCustomValue = v.IsvCustomValue[:0]
	v.Features = v.Features[:0]
	v.SubPurchaseOrders = v.SubPurchaseOrders[:0]
	v.OrderMessages = v.OrderMessages[:0]
	v.LogisticsInfos = v.LogisticsInfos[:0]
	v.SupplierFrom = ""
	v.SupplierUsername = ""
	v.DistributorFrom = ""
	v.DistributorUsername = ""
	v.BuyerTaobaoId = ""
	v.BuyerNick = ""
	v.TradeType = ""
	v.Created = ""
	v.Memo = ""
	v.SupplierMemo = ""
	v.AlipayNo = ""
	v.TotalFee = ""
	v.PostFee = ""
	v.DistributorPayment = ""
	v.BuyerPayment = ""
	v.SnapshotUrl = ""
	v.Status = ""
	v.PayTime = ""
	v.EndTime = ""
	v.ConsignTime = ""
	v.Modified = ""
	v.PayType = ""
	v.Shipping = ""
	v.LogisticsCompanyName = ""
	v.LogisticsId = ""
	v.AlipayOrderNo = ""
	v.ConfirmPaidFeeYuan = ""
	v.OpenBuyerUid = ""
	v.DeliveryTime = ""
	v.SignTime = ""
	v.AsdpAds = ""
	v.DeliveryCps = ""
	v.TcOrderId = 0
	v.Id = 0
	v.FenxiaoId = 0
	v.SupplierFlag = 0
	v.Receiver = nil
	v.ChannelCode = 0
	poolTopDpOrderDo.Put(v)
}

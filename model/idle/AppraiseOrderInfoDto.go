package idle

import (
	"sync"
)

// AppraiseOrderInfoDto 结构体
type AppraiseOrderInfoDto struct {
	// 鉴定商发货给买家的单号 主状态&gt;=7，验货中心已发货时填写
	Ac2buyerMailNo string `json:"ac2buyer_mail_no,omitempty" xml:"ac2buyer_mail_no,omitempty"`
	// 逆向流程鉴定商发货给卖家的单号 逆向流程中，主状态&gt;=103卖家已发货后发起逆向时填写
	Ac2sellerMailNo string `json:"ac2seller_mail_no,omitempty" xml:"ac2seller_mail_no,omitempty"`
	// 服务商appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 品牌Id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 买家收货地址。主状态&gt;=6，买家确认购买时填写
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 买家支付宝ID，用于卖家责任时赔付买家 逆向流程中，主状态105鉴定为赝品时填写
	BuyerAlipayUserId string `json:"buyer_alipay_user_id,omitempty" xml:"buyer_alipay_user_id,omitempty"`
	// 买家区
	BuyerArea string `json:"buyer_area,omitempty" xml:"buyer_area,omitempty"`
	// 买家城市
	BuyerCity string `json:"buyer_city,omitempty" xml:"buyer_city,omitempty"`
	// 买家关闭原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 买家村庄
	BuyerCountry string `json:"buyer_country,omitempty" xml:"buyer_country,omitempty"`
	// 买家手机号。主状态&gt;=6，买家确认购买时有值
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 买家省份
	BuyerProvince string `json:"buyer_province,omitempty" xml:"buyer_province,omitempty"`
	// 买家收货姓名。主状态&gt;=6，买家确认购买时有值
	BuyerReceiptName string `json:"buyer_receipt_name,omitempty" xml:"buyer_receipt_name,omitempty"`
	// 类目聚合场景,如&#34;YHB_3C&#34;
	CateAggScene string `json:"cate_agg_scene,omitempty" xml:"cate_agg_scene,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 预留 JSON 格式渠道业务数据
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 订单环境 &#39;online&#39;：线上环境；&#39;pre&#39;：测试环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 二手卖家承诺验货项。数组，每条记录的keyId代表验货项id，比如1001可能代表内存。valueId和valueName分别代表验货项值的id和描述。keyId、valueId的取值参考对接文档
	IdleAppraiseCheckpoints string `json:"idle_appraise_checkpoints,omitempty" xml:"idle_appraise_checkpoints,omitempty"`
	// 鉴定场景：1表示新品鉴定，2表示旧品鉴定
	IdleAppraiseScene string `json:"idle_appraise_scene,omitempty" xml:"idle_appraise_scene,omitempty"`
	// 商品详情页
	ItemDetailInfo string `json:"item_detail_info,omitempty" xml:"item_detail_info,omitempty"`
	// 订单主状态。(主状态,子状态,状态说明)示例如下： (&#34;1&#34;, &#34;1&#34;, &#34;买家拍下未付款&#34;) (&#34;2&#34;, &#34;1&#34;, &#34;买家拍下已付款&#34;) (&#34;3&#34;, &#34;1&#34;, &#34;卖家已发货&#34;) 等。详情参考对接文档
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单子状态，说明见order_status
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 卖家发货给验货中心的单号。主状态&gt;=3，卖家已发货时有值
	Seller2acMailNo string `json:"seller2ac_mail_no,omitempty" xml:"seller2ac_mail_no,omitempty"`
	// 卖家区
	SellerArea string `json:"seller_area,omitempty" xml:"seller_area,omitempty"`
	// 卖家城市
	SellerCity string `json:"seller_city,omitempty" xml:"seller_city,omitempty"`
	// 卖家关闭原因
	SellerCloseReason string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
	// 卖家村庄
	SellerCountry string `json:"seller_country,omitempty" xml:"seller_country,omitempty"`
	// 卖家手机号码。逆向流程中，主状态&gt;=102卖家已发货后发起逆向时有值
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 卖家省份
	SellerProvince string `json:"seller_province,omitempty" xml:"seller_province,omitempty"`
	// 逆向流程卖家收货地址。逆向流程中，主状态&gt;=102卖家已发货后发起逆向时有值
	SellerReceiptAddress string `json:"seller_receipt_address,omitempty" xml:"seller_receipt_address,omitempty"`
	// 卖家收货姓名，逆向退货用。逆向流程中，主状态&gt;=102卖家已发货后发起逆向时有值
	SellerReceiptName string `json:"seller_receipt_name,omitempty" xml:"seller_receipt_name,omitempty"`
	// spuId
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// spu信息jsonStr
	SpuInfo string `json:"spu_info,omitempty" xml:"spu_info,omitempty"`
	// 2022-10-01 00:00:00
	BuyerConfirmReceiptTimeout string `json:"buyer_confirm_receipt_timeout,omitempty" xml:"buyer_confirm_receipt_timeout,omitempty"`
	// 2022-10-01 00:00:00
	SellerConfirmBackAddressTime string `json:"seller_confirm_back_address_time,omitempty" xml:"seller_confirm_back_address_time,omitempty"`
	// 验货宝订单流程版本
	YhbVersion string `json:"yhb_version,omitempty" xml:"yhb_version,omitempty"`
	// 买家验货费支付宝流水单号（仅在买家需要承担时有效）
	BuyerInspectionZfbId string `json:"buyer_inspection_zfb_id,omitempty" xml:"buyer_inspection_zfb_id,omitempty"`
	// 买家验货费支付单号
	BuyerInspectionOrderId string `json:"buyer_inspection_order_id,omitempty" xml:"buyer_inspection_order_id,omitempty"`
	// 卖家验货费支付宝流水单号（仅在卖家需要承担时有效）
	SellerInspectionZfbId string `json:"seller_inspection_zfb_id,omitempty" xml:"seller_inspection_zfb_id,omitempty"`
	// 卖家验货费支付单号
	SellerInspectionOrderId string `json:"seller_inspection_order_id,omitempty" xml:"seller_inspection_order_id,omitempty"`
	// 纠纷单判责结论，只在产生纠纷单并在判责结论出具后有效
	DisputeConclusion string `json:"dispute_conclusion,omitempty" xml:"dispute_conclusion,omitempty"`
	// 卖家支付宝 pid
	SellerAlipayPid string `json:"seller_alipay_pid,omitempty" xml:"seller_alipay_pid,omitempty"`
	// 仓库 id
	AcStorageId string `json:"ac_storage_id,omitempty" xml:"ac_storage_id,omitempty"`
	// 有值代表使用了验货费券，表示金额
	InspectionCouponFee string `json:"inspection_coupon_fee,omitempty" xml:"inspection_coupon_fee,omitempty"`
	// 需付给买家的赔付金额，单位分。状态为：已付款后发货超时（101-1 101-2）、卖家取消订单（101-3 101-4）、鉴定为赝品（主状态105）时有值。
	Compensation2buyer int64 `json:"compensation2buyer,omitempty" xml:"compensation2buyer,omitempty"`
	// 服务商应收实时分账金额，分。仅当交易成功或者卖家无责买家不买 这两个状态有效，其余为0。当保证金被罚没的状态时，不会实时分账而是线下结算，所以只有上述两个状态有值。
	SupplierChargeFeeCent int64 `json:"supplier_charge_fee_cent,omitempty" xml:"supplier_charge_fee_cent,omitempty"`
	// 买家实付金额，单位：分
	ActualPaidFeeCent int64 `json:"actual_paid_fee_cent,omitempty" xml:"actual_paid_fee_cent,omitempty"`
	// 确定买家验货费金额，单位：分
	BuyerInspectionFee int64 `json:"buyer_inspection_fee,omitempty" xml:"buyer_inspection_fee,omitempty"`
	// 确定买家验货费金额，单位：分
	SellerInspectionFee int64 `json:"seller_inspection_fee,omitempty" xml:"seller_inspection_fee,omitempty"`
	// 是否是经主发发布/编辑的验货宝商品
	AppraiseFromNewPublisher bool `json:"appraise_from_new_publisher,omitempty" xml:"appraise_from_new_publisher,omitempty"`
	// true 表示交易成功卖家包验货费模式
	TradeSuccessSellerAssume bool `json:"trade_success_seller_assume,omitempty" xml:"trade_success_seller_assume,omitempty"`
}

var poolAppraiseOrderInfoDto = sync.Pool{
	New: func() any {
		return new(AppraiseOrderInfoDto)
	},
}

// GetAppraiseOrderInfoDto() 从对象池中获取AppraiseOrderInfoDto
func GetAppraiseOrderInfoDto() *AppraiseOrderInfoDto {
	return poolAppraiseOrderInfoDto.Get().(*AppraiseOrderInfoDto)
}

// ReleaseAppraiseOrderInfoDto 释放AppraiseOrderInfoDto
func ReleaseAppraiseOrderInfoDto(v *AppraiseOrderInfoDto) {
	v.Ac2buyerMailNo = ""
	v.Ac2sellerMailNo = ""
	v.AppKey = ""
	v.BizOrderId = ""
	v.BrandId = ""
	v.BuyerAddress = ""
	v.BuyerAlipayUserId = ""
	v.BuyerArea = ""
	v.BuyerCity = ""
	v.BuyerCloseReason = ""
	v.BuyerCountry = ""
	v.BuyerPhone = ""
	v.BuyerProvince = ""
	v.BuyerReceiptName = ""
	v.CateAggScene = ""
	v.Channel = ""
	v.ChannelData = ""
	v.Env = ""
	v.GmtCreate = ""
	v.IdleAppraiseCheckpoints = ""
	v.IdleAppraiseScene = ""
	v.ItemDetailInfo = ""
	v.OrderStatus = ""
	v.OrderSubStatus = ""
	v.Seller2acMailNo = ""
	v.SellerArea = ""
	v.SellerCity = ""
	v.SellerCloseReason = ""
	v.SellerCountry = ""
	v.SellerPhone = ""
	v.SellerProvince = ""
	v.SellerReceiptAddress = ""
	v.SellerReceiptName = ""
	v.SpuId = ""
	v.SpuInfo = ""
	v.BuyerConfirmReceiptTimeout = ""
	v.SellerConfirmBackAddressTime = ""
	v.YhbVersion = ""
	v.BuyerInspectionZfbId = ""
	v.BuyerInspectionOrderId = ""
	v.SellerInspectionZfbId = ""
	v.SellerInspectionOrderId = ""
	v.DisputeConclusion = ""
	v.SellerAlipayPid = ""
	v.AcStorageId = ""
	v.InspectionCouponFee = ""
	v.Compensation2buyer = 0
	v.SupplierChargeFeeCent = 0
	v.ActualPaidFeeCent = 0
	v.BuyerInspectionFee = 0
	v.SellerInspectionFee = 0
	v.AppraiseFromNewPublisher = false
	v.TradeSuccessSellerAssume = false
	poolAppraiseOrderInfoDto.Put(v)
}

package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeGetAPIResponse 零售+平台查询订单 API返回值
// alibaba.nlife.b2c.trade.get
//
// 查询零售+平台创建出来的订单详情
type AlibabaNlifeB2cTradeGetAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradeGetAPIResponseModel
}

// AlibabaNlifeB2cTradeGetAPIResponseModel is 零售+平台查询订单 成功返回结果
type AlibabaNlifeB2cTradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 购买的商品列表
	GoodsList []Goods `json:"goods_list,omitempty" xml:"goods_list>goods,omitempty"`
	// 支付资金渠道列表
	FundBillList []FundBill `json:"fund_bill_list,omitempty" xml:"fund_bill_list>fund_bill,omitempty"`
	// 退款列表
	RefundList []Refund `json:"refund_list,omitempty" xml:"refund_list>refund,omitempty"`
	// 物流状态列表
	LogisticsStatusList []LogisticsStatus `json:"logistics_status_list,omitempty" xml:"logistics_status_list>logistics_status,omitempty"`
	// 订单创建渠道，   API("经由零售+平台API创建"),     POS("经由零售+平台收银系统创建"),     TAOBAO("经由淘宝/天猫创建");
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 渠道数据，JSON
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 买家ID
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 买家类型： TAOBAO_OPENID("经阿里百川淘宝登陆获取到的用户ID"),     ALIPAY_OPENID("经蚂蚁金服开放平台支付宝App登陆获得的用户ID"),     WECHAT_OPENID("经微信开放平台微信App登陆获得的用户id"),     PHONE_NUMBER("通过手机号码登陆"),     APP_USERID("商户自由的用户ID"),     ANONYMOUS_USER("匿名用户");
	BuyerIdType string `json:"buyer_id_type,omitempty" xml:"buyer_id_type,omitempty"`
	// 商户自有的会员卡号
	OutCardNo string `json:"out_card_no,omitempty" xml:"out_card_no,omitempty"`
	// 订单描述
	OrderBody string `json:"order_body,omitempty" xml:"order_body,omitempty"`
	// 导购员ID
	SalesId string `json:"sales_id,omitempty" xml:"sales_id,omitempty"`
	// 提货方式:   LOGISTICS("物流发货"),     SELF_DELIVERY("门店自提");
	PickingUp string `json:"picking_up,omitempty" xml:"picking_up,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 支付时间
	GmtPay string `json:"gmt_pay,omitempty" xml:"gmt_pay,omitempty"`
	// 订单取消时间
	GmtCancel string `json:"gmt_cancel,omitempty" xml:"gmt_cancel,omitempty"`
	// 退款状态:    REFUNED("已全额退款"),     REFUNDED_PART("已部分退货");
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 物流状态:    WAIT_FOR_CONSIGN("有待发货商品"),     WAIT_FOR_SIGN("全部商品已发货"),     SIGNED("全部商品已签收"),     REJECTED("全部商品已拒收");
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// 扩展参数 JSON
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 透传参数，格式自定，查询时原样返回
	Attachment string `json:"attachment,omitempty" xml:"attachment,omitempty"`
	// payStatus
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 零售+订单号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 外部业务方订单号
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 全渠道订单号
	OmniTradeNo string `json:"omni_trade_no,omitempty" xml:"omni_trade_no,omitempty"`
	// 订单所在的零售+门店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 收银员名称
	SalesName string `json:"sales_name,omitempty" xml:"sales_name,omitempty"`
	// 订单总金额 基础单位，人民币 分
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 0：门店订单；1：全渠道订单；3：网直供订单
	TradeBizType *model.File `json:"trade_biz_type,omitempty" xml:"trade_biz_type,omitempty"`
}

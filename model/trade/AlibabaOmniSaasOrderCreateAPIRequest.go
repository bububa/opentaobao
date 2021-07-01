package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOmniSaasOrderCreateAPIRequest
订单创建接口 API请求
alibaba.omni.saas.order.create

服务商利用现有的saas系统和阿里完成交易系统的对接 */
type AlibabaOmniSaasOrderCreateAPIRequest struct {
	model.Params
	// 商品列表
	_goodsDetails []GoodsDetail
	// 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
	_buyerId string
	// 门店ID
	_storeId string
	// 收银设备类型
	_device string
	// 收银设备号
	_deviceNo string
	// 收银员标识
	_operatorId string
	// ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
	_buyerIdType string
	// ALIPAY：支付宝付款；BANK_CARD：刷卡
	_payChannel string
	// 商家自有优惠
	_couponInfos []CouponInfo
	// PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
	_storeIdType string
	// 请求号，用于标识一次请求
	_requestNo string
}

// New

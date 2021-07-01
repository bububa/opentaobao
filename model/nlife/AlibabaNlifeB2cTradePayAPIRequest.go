package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradePayAPIRequest
零售+平台支付订单 API请求
alibaba.nlife.b2c.trade.pay

零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步 */
type AlibabaNlifeB2cTradePayAPIRequest struct {
	model.Params
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 提货方式：    LOGISTICS("物流发货"),     SELF_DELIVERY("门店自提");
	_pickingUp string
	// 收货人
	_consignee string
	// 收货人电话
	_consigneePhoneNum string
	// 收货人地址
	_consigneeAddress string
	// ISV处支付时间
	_gmtPayment string
	// 支付资金各渠道列表
	_fundBillList []FundBill
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 实付金额，单位人民币分；该字段实际为必选，为兼容已经接入的isv设置成可选
	_actualPayFee int64
	// 只传out_trade_no时候，零售+门店号一定要传
	_storeId string
}

// New

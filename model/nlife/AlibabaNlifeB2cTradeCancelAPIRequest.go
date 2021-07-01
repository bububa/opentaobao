package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeCancelAPIRequest
零售+平台取消订单 API请求
alibaba.nlife.b2c.trade.cancel

零售+平台取消订单接口 */
type AlibabaNlifeB2cTradeCancelAPIRequest struct {
	model.Params
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店号
	_storeId string
}

// New

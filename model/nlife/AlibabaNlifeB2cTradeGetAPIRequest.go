package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeGetAPIRequest
零售+平台查询订单 API请求
alibaba.nlife.b2c.trade.get

查询零售+平台创建出来的订单详情 */
type AlibabaNlifeB2cTradeGetAPIRequest struct {
	model.Params
	// 零售+平台订单号,和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
	_storeId string
}

// New

package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderExchangeCreateAPIRequest
代理商积分兑换接口 API请求
alibaba.alicom.order.exchange.create

代理商调用该接口来进行积分兑换 */
type AlibabaAlicomOrderExchangeCreateAPIRequest struct {
	model.Params
	// 入参
	_exchangeModel *ExchangeModel
}

// New

package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTradeAlianceCreateAPIRequest
推客平台订单回流 API请求
alibaba.trade.aliance.create

推客平台订单回流 */
type AlibabaTradeAlianceCreateAPIRequest struct {
	model.Params
	// 下单请求
	_paramIsvCreateOrderParam *IsvCreateOrderParam
}

// New

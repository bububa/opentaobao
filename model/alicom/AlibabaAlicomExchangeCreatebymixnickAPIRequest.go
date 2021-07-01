package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomExchangeCreatebymixnickAPIRequest
代理商积分兑换接口tae API请求
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae */
type AlibabaAlicomExchangeCreatebymixnickAPIRequest struct {
	model.Params
	// 入参
	_exchangeModel *ExchangeModel
}

// New

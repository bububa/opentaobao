package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOnsiteTradeOldrefundAPIRequest
线下新退款接口（专为老退款接口调用） API请求
alibaba.mos.onsite.trade.oldrefund

线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回 */
type AlibabaMosOnsiteTradeOldrefundAPIRequest struct {
	model.Params
	// 交易退款请求
	_onsiteRefundRequest *OnsiteRefundRequest
}

// New

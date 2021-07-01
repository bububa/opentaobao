package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWtOrderExchangePartnerChecktbuserAPIRequest
积分兑换校验淘宝账号是否存在 API请求
alibaba.wt.order.exchange.partner.checktbuser

积分兑换之前校验淘宝账号是否存在 */
type AlibabaWtOrderExchangePartnerChecktbuserAPIRequest struct {
	model.Params
	// model入参
	_outExchangeModel *OutExchangeModel
}

// New

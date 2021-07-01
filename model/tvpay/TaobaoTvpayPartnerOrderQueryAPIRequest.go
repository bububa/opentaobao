package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayPartnerOrderQueryAPIRequest
商户查询订单 API请求
taobao.tvpay.partner.order.query

给商户提供的查询订单状态的API */
type TaobaoTvpayPartnerOrderQueryAPIRequest struct {
	model.Params
	// 商户订单号
	_orderNo string
}

// New

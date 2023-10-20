package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayPartnerOrderQuery 商户查询订单
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
func TaobaoTvpayPartnerOrderQuery(clt *core.SDKClient, req *tvpay.TaobaoTvpayPartnerOrderQueryAPIRequest, resp *tvpay.TaobaoTvpayPartnerOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

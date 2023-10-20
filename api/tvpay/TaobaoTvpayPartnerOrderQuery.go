package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpaypartnerorderquery 商户查询订单
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
func Taobaotvpaypartnerorderquery(clt *core.SDKClient, req *tvpay.TaobaotvpaypartnerorderqueryAPIRequest, session string) (*tvpay.TaobaotvpaypartnerorderqueryAPIResponse, error) {
	var resp tvpay.TaobaotvpaypartnerorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

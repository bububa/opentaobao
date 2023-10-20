package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpayorderquery tv支付查询订单状态
// taobao.tvpay.order.query
//
// tv支付查询订单状态
func Taobaotvpayorderquery(clt *core.SDKClient, req *tvpay.TaobaotvpayorderqueryAPIRequest, session string) (*tvpay.TaobaotvpayorderqueryAPIResponse, error) {
	var resp tvpay.TaobaotvpayorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

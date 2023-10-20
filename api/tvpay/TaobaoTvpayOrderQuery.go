package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayOrderQuery tv支付查询订单状态
// taobao.tvpay.order.query
//
// tv支付查询订单状态
func TaobaoTvpayOrderQuery(clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderQueryAPIRequest, resp *tvpay.TaobaoTvpayOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

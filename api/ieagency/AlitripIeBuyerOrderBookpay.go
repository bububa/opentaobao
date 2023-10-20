package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// AlitripIeBuyerOrderBookpay 【国际机票】下单预定支付
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
func AlitripIeBuyerOrderBookpay(clt *core.SDKClient, req *ieagency.AlitripIeBuyerOrderBookpayAPIRequest, resp *ieagency.AlitripIeBuyerOrderBookpayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

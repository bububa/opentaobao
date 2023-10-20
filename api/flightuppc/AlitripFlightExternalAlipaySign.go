package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipaySign 支付宝小程序验签
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
func AlitripFlightExternalAlipaySign(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipaySignAPIRequest, resp *flightuppc.AlitripFlightExternalAlipaySignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayVerify 支付宝小程序验签
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
func AlitripFlightExternalAlipayVerify(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayVerifyAPIRequest, resp *flightuppc.AlitripFlightExternalAlipayVerifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

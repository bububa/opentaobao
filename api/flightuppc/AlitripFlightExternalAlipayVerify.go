package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayVerify 支付宝小程序验签
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
func AlitripFlightExternalAlipayVerify(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayVerifyAPIRequest, session string) (*flightuppc.AlitripFlightExternalAlipayVerifyAPIResponse, error) {
	var resp flightuppc.AlitripFlightExternalAlipayVerifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

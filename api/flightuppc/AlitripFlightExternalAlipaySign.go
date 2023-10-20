package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipaySign 支付宝小程序验签
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
func AlitripFlightExternalAlipaySign(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipaySignAPIRequest, session string) (*flightuppc.AlitripFlightExternalAlipaySignAPIResponse, error) {
	var resp flightuppc.AlitripFlightExternalAlipaySignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

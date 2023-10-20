package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayEncrypt 支付宝小程序明文加密
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
func AlitripFlightExternalAlipayEncrypt(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayEncryptAPIRequest, session string) (*flightuppc.AlitripFlightExternalAlipayEncryptAPIResponse, error) {
	var resp flightuppc.AlitripFlightExternalAlipayEncryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayDecrypt 支付宝小程序密文解密
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
func AlitripFlightExternalAlipayDecrypt(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayDecryptAPIRequest, session string) (*flightuppc.AlitripFlightExternalAlipayDecryptAPIResponse, error) {
	var resp flightuppc.AlitripFlightExternalAlipayDecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightexternalalipayencrypt 支付宝小程序明文加密
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
func Alitripflightexternalalipayencrypt(clt *core.SDKClient, req *flightuppc.AlitripflightexternalalipayencryptAPIRequest, session string) (*flightuppc.AlitripflightexternalalipayencryptAPIResponse, error) {
	var resp flightuppc.AlitripflightexternalalipayencryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

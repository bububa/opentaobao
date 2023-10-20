package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightexternalalipayverify 支付宝小程序验签
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
func Alitripflightexternalalipayverify(clt *core.SDKClient, req *flightuppc.AlitripflightexternalalipayverifyAPIRequest, session string) (*flightuppc.AlitripflightexternalalipayverifyAPIResponse, error) {
	var resp flightuppc.AlitripflightexternalalipayverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

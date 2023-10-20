package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightexternalalipaysign 支付宝小程序验签
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
func Alitripflightexternalalipaysign(clt *core.SDKClient, req *flightuppc.AlitripflightexternalalipaysignAPIRequest, session string) (*flightuppc.AlitripflightexternalalipaysignAPIResponse, error) {
	var resp flightuppc.AlitripflightexternalalipaysignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

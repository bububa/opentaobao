package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightexternalalipaysnquery 支付宝小程序查询证书序列号
// alitrip.flight.external.alipay.sn.query
//
// 支付宝小程序查询证书序列号
func Alitripflightexternalalipaysnquery(clt *core.SDKClient, req *flightuppc.AlitripflightexternalalipaysnqueryAPIRequest, session string) (*flightuppc.AlitripflightexternalalipaysnqueryAPIResponse, error) {
	var resp flightuppc.AlitripflightexternalalipaysnqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

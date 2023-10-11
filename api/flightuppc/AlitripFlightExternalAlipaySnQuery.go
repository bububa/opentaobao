package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipaySnQuery 支付宝小程序查询证书序列号
// alitrip.flight.external.alipay.sn.query
//
// 支付宝小程序查询证书序列号
func AlitripFlightExternalAlipaySnQuery(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipaySnQueryAPIRequest, session string) (*flightuppc.AlitripFlightExternalAlipaySnQueryAPIResponse, error) {
	var resp flightuppc.AlitripFlightExternalAlipaySnQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

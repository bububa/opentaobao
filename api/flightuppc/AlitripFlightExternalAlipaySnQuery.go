package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipaySnQuery 支付宝小程序查询证书序列号
// alitrip.flight.external.alipay.sn.query
//
// 支付宝小程序查询证书序列号
func AlitripFlightExternalAlipaySnQuery(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipaySnQueryAPIRequest, resp *flightuppc.AlitripFlightExternalAlipaySnQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

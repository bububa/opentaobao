package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayEncrypt 支付宝小程序明文加密
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
func AlitripFlightExternalAlipayEncrypt(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayEncryptAPIRequest, resp *flightuppc.AlitripFlightExternalAlipayEncryptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

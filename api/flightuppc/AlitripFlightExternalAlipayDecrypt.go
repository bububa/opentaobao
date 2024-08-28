package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayDecrypt 支付宝小程序密文解密
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
func AlitripFlightExternalAlipayDecrypt(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayDecryptAPIRequest, resp *flightuppc.AlitripFlightExternalAlipayDecryptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

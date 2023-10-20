package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipayDecrypt 支付宝小程序密文解密
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
func AlitripFlightExternalAlipayDecrypt(clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipayDecryptAPIRequest, resp *flightuppc.AlitripFlightExternalAlipayDecryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

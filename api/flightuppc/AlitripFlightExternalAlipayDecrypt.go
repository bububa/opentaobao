package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightexternalalipaydecrypt 支付宝小程序密文解密
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
func Alitripflightexternalalipaydecrypt(clt *core.SDKClient, req *flightuppc.AlitripflightexternalalipaydecryptAPIRequest, session string) (*flightuppc.AlitripflightexternalalipaydecryptAPIResponse, error) {
	var resp flightuppc.AlitripflightexternalalipaydecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

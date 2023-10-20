package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopOrderWithhold 代扣支付
// cainiao.endpoint.locker.top.order.withhold
//
// 提供代扣，允许有一笔欠款。
func CainiaoEndpointLockerTopOrderWithhold(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderWithholdAPIRequest, resp *cainiaolocker.CainiaoEndpointLockerTopOrderWithholdAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

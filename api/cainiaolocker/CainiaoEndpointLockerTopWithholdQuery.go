package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopWithholdQuery 查询能否代扣
// cainiao.endpoint.locker.top.withhold.query
//
// 查询是否有代扣欠款，是否签署代扣协议。
func CainiaoEndpointLockerTopWithholdQuery(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopWithholdQueryAPIRequest, resp *cainiaolocker.CainiaoEndpointLockerTopWithholdQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

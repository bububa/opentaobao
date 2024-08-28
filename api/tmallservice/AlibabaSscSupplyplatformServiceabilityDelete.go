package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceabilityDelete 删除服务能力
// alibaba.ssc.supplyplatform.serviceability.delete
//
// 删除服务能力
func AlibabaSscSupplyplatformServiceabilityDelete(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

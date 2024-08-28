package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicecapacityDelete 服务容量删除
// alibaba.ssc.supplyplatform.servicecapacity.delete
//
// 服务容量删除
func AlibabaSscSupplyplatformServicecapacityDelete(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

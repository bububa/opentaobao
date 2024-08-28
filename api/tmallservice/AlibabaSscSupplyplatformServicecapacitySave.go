package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicecapacitySave 保存服务容量
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
func AlibabaSscSupplyplatformServicecapacitySave(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicecapacitySaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

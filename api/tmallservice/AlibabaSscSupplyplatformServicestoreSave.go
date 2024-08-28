package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicestoreSave 保存网点
// alibaba.ssc.supplyplatform.servicestore.save
//
// 网点创建、修改
func AlibabaSscSupplyplatformServicestoreSave(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreSaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicestoreSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicedefinitionQuerysku 服务sku查询
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
func AlibabaSscSupplyplatformServicedefinitionQuerysku(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

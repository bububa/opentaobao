package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicestoreOffline 网点下线
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
func AlibabaSscSupplyplatformServicestoreOffline(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

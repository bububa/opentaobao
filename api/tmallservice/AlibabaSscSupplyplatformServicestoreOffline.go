package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServicestoreOffline 网点下线
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
func AlibabaSscSupplyplatformServicestoreOffline(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

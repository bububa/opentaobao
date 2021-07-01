package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* AlibabaSscSupplyplatformServicestoreOffline
网点下线
alibaba.ssc.supplyplatform.servicestore.offline

网点下线功能 */
func AlibabaSscSupplyplatformServicestoreOffline(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
网点下线 
alibaba.ssc.supplyplatform.servicestore.offline

网点下线功能
*/
func AlibabaSscSupplyplatformServicestoreOffline(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreOfflineRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServicestoreOfflineResponse, error) {
    var resp tmallservice.AlibabaSscSupplyplatformServicestoreOfflineAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务容量删除 
alibaba.ssc.supplyplatform.servicecapacity.delete

服务容量删除
*/
func AlibabaSscSupplyplatformServicecapacityDelete(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteResponse, error) {
    var resp tmallservice.AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
保存网点 
alibaba.ssc.supplyplatform.servicestore.save

网点创建、修改
*/
func AlibabaSscSupplyplatformServicestoreSave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServicestoreSaveRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServicestoreSaveAPIResponse, error) {
    var resp tmallservice.AlibabaSscSupplyplatformServicestoreSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

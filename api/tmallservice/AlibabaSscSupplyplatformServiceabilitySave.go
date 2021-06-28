package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
保存服务能力 
alibaba.ssc.supplyplatform.serviceability.save

保存服务能力
*/
func AlibabaSscSupplyplatformServiceabilitySave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceabilitySaveRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceabilitySaveAPIResponse, error) {
    var resp tmallservice.AlibabaSscSupplyplatformServiceabilitySaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

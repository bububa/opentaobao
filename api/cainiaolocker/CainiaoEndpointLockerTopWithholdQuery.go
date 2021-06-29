package cainiaolocker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaolocker"
)

/* 
查询能否代扣 
cainiao.endpoint.locker.top.withhold.query

查询是否有代扣欠款，是否签署代扣协议。
*/
func CainiaoEndpointLockerTopWithholdQuery(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopWithholdQueryRequest, session string) (*cainiaolocker.CainiaoEndpointLockerTopWithholdQueryAPIResponse, error) {
    var resp cainiaolocker.CainiaoEndpointLockerTopWithholdQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

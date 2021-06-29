package cainiaolocker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaolocker"
)

/* 
代扣支付 
cainiao.endpoint.locker.top.order.withhold

提供代扣，允许有一笔欠款。
*/
func CainiaoEndpointLockerTopOrderWithhold(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderWithholdRequest, session string) (*cainiaolocker.CainiaoEndpointLockerTopOrderWithholdAPIResponse, error) {
    var resp cainiaolocker.CainiaoEndpointLockerTopOrderWithholdAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

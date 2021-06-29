package cainiaolocker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaolocker"
)

/* 
查询订单是否由裹裹发送消息 
cainiao.endpoint.locker.top.order.noticesend.query

合作公司查询消息发送的接口，判断是否裹裹发送消息
*/
func CainiaoEndpointLockerTopOrderNoticesendQuery(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderNoticesendQueryRequest, session string) (*cainiaolocker.CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse, error) {
    var resp cainiaolocker.CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

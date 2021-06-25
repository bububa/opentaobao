package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
查询面单服务订购及面单使用情况 
cainiao.waybill.ii.search

获取发货地&CP开通状态&账户的使用情况
*/
func CainiaoWaybillIiSearch(clt *core.SDKClient, req *waybill.CainiaoWaybillIiSearchRequest, session string) (*waybill.CainiaoWaybillIiSearchResponse, error) {
    var resp waybill.CainiaoWaybillIiSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

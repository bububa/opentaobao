package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
查询面单服务订购及面单使用情况v1.0 
taobao.wlb.waybill.i.search

获取发货地&CP开通状态&账户的使用情况
*/
func TaobaoWlbWaybillISearch(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillISearchRequest, session string) (*waybill.TaobaoWlbWaybillISearchResponse, error) {
    var resp waybill.TaobaoWlbWaybillISearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

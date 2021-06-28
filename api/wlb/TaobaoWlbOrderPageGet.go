package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
分页查询物流宝订单 
taobao.wlb.order.page.get

分页查询物流宝订单
*/
func TaobaoWlbOrderPageGet(clt *core.SDKClient, req *wlb.TaobaoWlbOrderPageGetRequest, session string) (*wlb.TaobaoWlbOrderPageGetAPIResponse, error) {
    var resp wlb.TaobaoWlbOrderPageGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

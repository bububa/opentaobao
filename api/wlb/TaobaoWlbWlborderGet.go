package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
根据物流宝订单编号查询物流宝订单概要信息 
taobao.wlb.wlborder.get

根据物流宝订单编号查询物流宝订单概要信息
*/
func TaobaoWlbWlborderGet(clt *core.SDKClient, req *wlb.TaobaoWlbWlborderGetRequest, session string) (*wlb.TaobaoWlbWlborderGetResponse, error) {
    var resp wlb.TaobaoWlbWlborderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

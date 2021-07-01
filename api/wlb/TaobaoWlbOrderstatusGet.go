package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
物流宝订单流转状态查询 
taobao.wlb.orderstatus.get

根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
*/
func TaobaoWlbOrderstatusGet(clt *core.SDKClient, req *wlb.TaobaoWlbOrderstatusGetAPIRequest, session string) (*wlb.TaobaoWlbOrderstatusGetAPIResponse, error) {
    var resp wlb.TaobaoWlbOrderstatusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

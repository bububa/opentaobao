package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
查询库存详情 
taobao.wlb.inventory.detail.get

查询库存详情，通过商品ID获取发送请求的卖家的库存详情
*/
func TaobaoWlbInventoryDetailGet(clt *core.SDKClient, req *wlb.TaobaoWlbInventoryDetailGetRequest, session string) (*wlb.TaobaoWlbInventoryDetailGetResponse, error) {
    var resp wlb.TaobaoWlbInventoryDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
分页查询商品 
taobao.wlb.item.query

根据状态、卖家、SKU等信息查询商品列表
*/
func TaobaoWlbItemQuery(clt *core.SDKClient, req *wlb.TaobaoWlbItemQueryRequest, session string) (*wlb.TaobaoWlbItemQueryAPIResponse, error) {
    var resp wlb.TaobaoWlbItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

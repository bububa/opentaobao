package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
根据商品ID查询所有库存变更记录 
taobao.wlb.inventorylog.query

通过商品ID等几个条件来分页查询库存变更记录
*/
func TaobaoWlbInventorylogQuery(clt *core.SDKClient, req *wlb.TaobaoWlbInventorylogQueryRequest, session string) (*wlb.TaobaoWlbInventorylogQueryResponse, error) {
    var resp wlb.TaobaoWlbInventorylogQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

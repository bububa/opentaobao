package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
通过订单列表批量获取库存损益单信息 
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息
*/
func TaobaoWlbWmsInventoryProfitlossGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsInventoryProfitlossGetRequest, session string) (*wms.TaobaoWlbWmsInventoryProfitlossGetResponse, error) {
    var resp wms.TaobaoWlbWmsInventoryProfitlossGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

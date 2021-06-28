package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
查询组合商品的组合关系 
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系
*/
func TaobaoWlbWmsItemCombinationGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsItemCombinationGetRequest, session string) (*wms.TaobaoWlbWmsItemCombinationGetAPIResponse, error) {
    var resp wms.TaobaoWlbWmsItemCombinationGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

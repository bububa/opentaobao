package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
获取全渠道门店商品 
taobao.omniitem.item.get

通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
*/
func TaobaoOmniitemItemGet(clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemGetRequest, session string) (*omniorder.TaobaoOmniitemItemGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniitemItemGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
获取前台展示的店铺内卖家自定义商品类目 
taobao.sellercats.list.get

此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
*/
func TaobaoSellercatsListGet(clt *core.SDKClient, req *shop.TaobaoSellercatsListGetAPIRequest, session string) (*shop.TaobaoSellercatsListGetAPIResponse, error) {
    var resp shop.TaobaoSellercatsListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

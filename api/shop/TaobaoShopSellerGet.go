package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
卖家店铺基础信息查询 
taobao.shop.seller.get

获取卖家店铺的基本信息
*/
func TaobaoShopSellerGet(clt *core.SDKClient, req *shop.TaobaoShopSellerGetRequest, session string) (*shop.TaobaoShopSellerGetAPIResponse, error) {
    var resp shop.TaobaoShopSellerGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

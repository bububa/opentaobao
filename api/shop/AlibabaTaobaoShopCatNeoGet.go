package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
店铺mtop接口鉴权虚拟api 
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
func AlibabaTaobaoShopCatNeoGet(clt *core.SDKClient, req *shop.AlibabaTaobaoShopCatNeoGetRequest, session string) (*shop.AlibabaTaobaoShopCatNeoGetAPIResponse, error) {
    var resp shop.AlibabaTaobaoShopCatNeoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

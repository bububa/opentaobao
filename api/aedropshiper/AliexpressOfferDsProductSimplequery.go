package aedropshiper

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aedropshiper"
)

/* 
Dropshipper查询单个商品的简易信息 
aliexpress.offer.ds.product.simplequery

提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用
*/
func AliexpressOfferDsProductSimplequery(clt *core.SDKClient, req *aedropshiper.AliexpressOfferDsProductSimplequeryRequest, session string) (*aedropshiper.AliexpressOfferDsProductSimplequeryAPIResponse, error) {
    var resp aedropshiper.AliexpressOfferDsProductSimplequeryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

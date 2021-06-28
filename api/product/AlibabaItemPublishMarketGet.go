package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取商家可发布商品的市场信息 
alibaba.item.publish.market.get

获取商家可发布商品的市场信息
*/
func AlibabaItemPublishMarketGet(clt *core.SDKClient, req *product.AlibabaItemPublishMarketGetRequest, session string) (*product.AlibabaItemPublishMarketGetAPIResponse, error) {
    var resp product.AlibabaItemPublishMarketGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

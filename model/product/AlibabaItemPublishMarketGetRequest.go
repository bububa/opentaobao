package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家可发布商品的市场信息 APIRequest
alibaba.item.publish.market.get

获取商家可发布商品的市场信息
*/
type AlibabaItemPublishMarketGetRequest struct {
    model.Params

}

func NewAlibabaItemPublishMarketGetRequest() *AlibabaItemPublishMarketGetRequest{
    return &AlibabaItemPublishMarketGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemPublishMarketGetRequest) GetApiMethodName() string {
    return "alibaba.item.publish.market.get"
}

func (r AlibabaItemPublishMarketGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}



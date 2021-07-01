package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家可发布商品的市场信息 API请求
alibaba.item.publish.market.get

获取商家可发布商品的市场信息
*/
type AlibabaItemPublishMarketGetAPIRequest struct {
    model.Params
}

// 初始化AlibabaItemPublishMarketGetAPIRequest对象
func NewAlibabaItemPublishMarketGetRequest() *AlibabaItemPublishMarketGetAPIRequest{
    return &AlibabaItemPublishMarketGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishMarketGetAPIRequest) GetApiMethodName() string {
    return "alibaba.item.publish.market.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishMarketGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}

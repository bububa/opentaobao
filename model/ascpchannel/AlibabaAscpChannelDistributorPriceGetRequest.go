package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
链渠道中心淘外分销价格查询(分销商专用) API请求
alibaba.ascp.channel.distributor.price.get

此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorPriceGetRequest struct {
    model.Params
    // 价格入参
    _priceRequest   *Pricerequest
}

// 初始化AlibabaAscpChannelDistributorPriceGetRequest对象
func NewAlibabaAscpChannelDistributorPriceGetRequest() *AlibabaAscpChannelDistributorPriceGetRequest{
    return &AlibabaAscpChannelDistributorPriceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorPriceGetRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.price.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PriceRequest Setter
// 价格入参
func (r *AlibabaAscpChannelDistributorPriceGetRequest) SetPriceRequest(_priceRequest *Pricerequest) error {
    r._priceRequest = _priceRequest
    r.Set("price_request", _priceRequest)
    return nil
}

// PriceRequest Getter
func (r AlibabaAscpChannelDistributorPriceGetRequest) GetPriceRequest() *Pricerequest {
    return r._priceRequest
}

package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
链渠道中心淘外分销价格查询(分销商专用) APIRequest
alibaba.ascp.channel.distributor.price.get

此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorPriceGetRequest struct {
    model.Params

    // 价格入参
    priceRequest   *Pricerequest 

}

func NewAlibabaAscpChannelDistributorPriceGetRequest() *AlibabaAscpChannelDistributorPriceGetRequest{
    return &AlibabaAscpChannelDistributorPriceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelDistributorPriceGetRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.price.get"
}

func (r AlibabaAscpChannelDistributorPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelDistributorPriceGetRequest) SetPriceRequest(priceRequest *Pricerequest) error {
    r.priceRequest = priceRequest
    r.Set("price_request", priceRequest)
    return nil
}

func (r AlibabaAscpChannelDistributorPriceGetRequest) GetPriceRequest() *Pricerequest {
    return r.priceRequest
}


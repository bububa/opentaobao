package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取供应链渠道中心品的详情接口（淘外分销商专用） APIRequest
alibaba.ascp.channel.distributor.product.detail

此api为淘外分销的品批量查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorProductDetailRequest struct {
    model.Params

    // 产品详情查询入参
    productDetailRequest   *ProductDetailQueryRequestForDistributor 

}

func NewAlibabaAscpChannelDistributorProductDetailRequest() *AlibabaAscpChannelDistributorProductDetailRequest{
    return &AlibabaAscpChannelDistributorProductDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelDistributorProductDetailRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.product.detail"
}

func (r AlibabaAscpChannelDistributorProductDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelDistributorProductDetailRequest) SetProductDetailRequest(productDetailRequest *ProductDetailQueryRequestForDistributor) error {
    r.productDetailRequest = productDetailRequest
    r.Set("product_detail_request", productDetailRequest)
    return nil
}

func (r AlibabaAscpChannelDistributorProductDetailRequest) GetProductDetailRequest() *ProductDetailQueryRequestForDistributor {
    return r.productDetailRequest
}


package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴dropshipping 产品信息获取 APIRequest
alibaba.dropshipping.product.get

阿里巴巴dropshipping 产品信息获取
*/
type AlibabaDropshippingProductGetRequest struct {
    model.Params

    // {}
    paramDistributionSaleProductRequest   *DistributionSaleProductRequest 

}

func NewAlibabaDropshippingProductGetRequest() *AlibabaDropshippingProductGetRequest{
    return &AlibabaDropshippingProductGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDropshippingProductGetRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.product.get"
}

func (r AlibabaDropshippingProductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDropshippingProductGetRequest) SetParamDistributionSaleProductRequest(paramDistributionSaleProductRequest *DistributionSaleProductRequest) error {
    r.paramDistributionSaleProductRequest = paramDistributionSaleProductRequest
    r.Set("param_distribution_sale_product_request", paramDistributionSaleProductRequest)
    return nil
}

func (r AlibabaDropshippingProductGetRequest) GetParamDistributionSaleProductRequest() *DistributionSaleProductRequest {
    return r.paramDistributionSaleProductRequest
}


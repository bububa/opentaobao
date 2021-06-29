package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴dropshipping 产品信息获取 API请求
alibaba.dropshipping.product.get

阿里巴巴dropshipping 产品信息获取
*/
type AlibabaDropshippingProductGetRequest struct {
    model.Params
    // {}
    _paramDistributionSaleProductRequest   *DistributionSaleProductRequest
}

// 初始化AlibabaDropshippingProductGetRequest对象
func NewAlibabaDropshippingProductGetRequest() *AlibabaDropshippingProductGetRequest{
    return &AlibabaDropshippingProductGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingProductGetRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.product.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingProductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDistributionSaleProductRequest Setter
// {}
func (r *AlibabaDropshippingProductGetRequest) SetParamDistributionSaleProductRequest(_paramDistributionSaleProductRequest *DistributionSaleProductRequest) error {
    r._paramDistributionSaleProductRequest = _paramDistributionSaleProductRequest
    r.Set("param_distribution_sale_product_request", _paramDistributionSaleProductRequest)
    return nil
}

// ParamDistributionSaleProductRequest Getter
func (r AlibabaDropshippingProductGetRequest) GetParamDistributionSaleProductRequest() *DistributionSaleProductRequest {
    return r._paramDistributionSaleProductRequest
}

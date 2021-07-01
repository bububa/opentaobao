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
type AlibabaDropshippingProductGetAPIRequest struct {
    model.Params
    // {}
    _paramDistributionSaleProductRequest   *DistributionSaleProductRequest
}

// 初始化AlibabaDropshippingProductGetAPIRequest对象
func NewAlibabaDropshippingProductGetRequest() *AlibabaDropshippingProductGetAPIRequest{
    return &AlibabaDropshippingProductGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingProductGetAPIRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.product.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingProductGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDistributionSaleProductRequest Setter
// {}
func (r *AlibabaDropshippingProductGetAPIRequest) SetParamDistributionSaleProductRequest(_paramDistributionSaleProductRequest *DistributionSaleProductRequest) error {
    r._paramDistributionSaleProductRequest = _paramDistributionSaleProductRequest
    r.Set("param_distribution_sale_product_request", _paramDistributionSaleProductRequest)
    return nil
}

// ParamDistributionSaleProductRequest Getter
func (r AlibabaDropshippingProductGetAPIRequest) GetParamDistributionSaleProductRequest() *DistributionSaleProductRequest {
    return r._paramDistributionSaleProductRequest
}

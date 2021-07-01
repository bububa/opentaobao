package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Single Product Info API请求
aliexpress.solution.product.info.get

Get Single Product Info
*/
type AliexpressSolutionProductInfoGetAPIRequest struct {
    model.Params
    // product ID
    _productId   int64
}

// 初始化AliexpressSolutionProductInfoGetAPIRequest对象
func NewAliexpressSolutionProductInfoGetRequest() *AliexpressSolutionProductInfoGetAPIRequest{
    return &AliexpressSolutionProductInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductInfoGetAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// product ID
func (r *AliexpressSolutionProductInfoGetAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AliexpressSolutionProductInfoGetAPIRequest) GetProductId() int64 {
    return r._productId
}

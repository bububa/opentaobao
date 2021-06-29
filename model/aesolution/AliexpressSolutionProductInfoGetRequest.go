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
type AliexpressSolutionProductInfoGetRequest struct {
    model.Params
    // product ID
    productId   int64
}

// 初始化AliexpressSolutionProductInfoGetRequest对象
func NewAliexpressSolutionProductInfoGetRequest() *AliexpressSolutionProductInfoGetRequest{
    return &AliexpressSolutionProductInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductInfoGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// product ID
func (r *AliexpressSolutionProductInfoGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AliexpressSolutionProductInfoGetRequest) GetProductId() int64 {
    return r.productId
}

package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Single Product Info APIRequest
aliexpress.solution.product.info.get

Get Single Product Info
*/
type AliexpressSolutionProductInfoGetRequest struct {
    model.Params

    // product ID
    productId   int64 

}

func NewAliexpressSolutionProductInfoGetRequest() *AliexpressSolutionProductInfoGetRequest{
    return &AliexpressSolutionProductInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionProductInfoGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.info.get"
}

func (r AliexpressSolutionProductInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionProductInfoGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AliexpressSolutionProductInfoGetRequest) GetProductId() int64 {
    return r.productId
}


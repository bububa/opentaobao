package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.delete APIRequest
aliexpress.solution.batch.product.delete

Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
*/
type AliexpressSolutionBatchProductDeleteRequest struct {
    model.Params

    // maximum 100
    productIds   []int64 

}

func NewAliexpressSolutionBatchProductDeleteRequest() *AliexpressSolutionBatchProductDeleteRequest{
    return &AliexpressSolutionBatchProductDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionBatchProductDeleteRequest) GetApiMethodName() string {
    return "aliexpress.solution.batch.product.delete"
}

func (r AliexpressSolutionBatchProductDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionBatchProductDeleteRequest) SetProductIds(productIds []int64) error {
    r.productIds = productIds
    r.Set("product_ids", productIds)
    return nil
}

func (r AliexpressSolutionBatchProductDeleteRequest) GetProductIds() []int64 {
    return r.productIds
}


package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.price.update APIRequest
aliexpress.solution.batch.product.price.update

batch product price update operation for oversea sellers
*/
type AliexpressSolutionBatchProductPriceUpdateRequest struct {
    model.Params

    // The product list, in which the price needs to be updated. Maximum length:20
    mutipleProductUpdateList   []SynchronizeProductRequestDto 

}

func NewAliexpressSolutionBatchProductPriceUpdateRequest() *AliexpressSolutionBatchProductPriceUpdateRequest{
    return &AliexpressSolutionBatchProductPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionBatchProductPriceUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.batch.product.price.update"
}

func (r AliexpressSolutionBatchProductPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionBatchProductPriceUpdateRequest) SetMutipleProductUpdateList(mutipleProductUpdateList []SynchronizeProductRequestDto) error {
    r.mutipleProductUpdateList = mutipleProductUpdateList
    r.Set("mutiple_product_update_list", mutipleProductUpdateList)
    return nil
}

func (r AliexpressSolutionBatchProductPriceUpdateRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
    return r.mutipleProductUpdateList
}


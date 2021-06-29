package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.price.update API请求
aliexpress.solution.batch.product.price.update

batch product price update operation for oversea sellers
*/
type AliexpressSolutionBatchProductPriceUpdateRequest struct {
    model.Params
    // The product list, in which the price needs to be updated. Maximum length:20
    _mutipleProductUpdateList   []SynchronizeProductRequestDTO
}

// 初始化AliexpressSolutionBatchProductPriceUpdateRequest对象
func NewAliexpressSolutionBatchProductPriceUpdateRequest() *AliexpressSolutionBatchProductPriceUpdateRequest{
    return &AliexpressSolutionBatchProductPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductPriceUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.batch.product.price.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MutipleProductUpdateList Setter
// The product list, in which the price needs to be updated. Maximum length:20
func (r *AliexpressSolutionBatchProductPriceUpdateRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDTO) error {
    r._mutipleProductUpdateList = _mutipleProductUpdateList
    r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
    return nil
}

// MutipleProductUpdateList Getter
func (r AliexpressSolutionBatchProductPriceUpdateRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDTO {
    return r._mutipleProductUpdateList
}

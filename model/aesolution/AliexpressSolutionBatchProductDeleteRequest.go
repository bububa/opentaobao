package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.delete API请求
aliexpress.solution.batch.product.delete

Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
*/
type AliexpressSolutionBatchProductDeleteRequest struct {
    model.Params
    // maximum 100
    _productIds   []int64
}

// 初始化AliexpressSolutionBatchProductDeleteRequest对象
func NewAliexpressSolutionBatchProductDeleteRequest() *AliexpressSolutionBatchProductDeleteRequest{
    return &AliexpressSolutionBatchProductDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductDeleteRequest) GetApiMethodName() string {
    return "aliexpress.solution.batch.product.delete"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductIds Setter
// maximum 100
func (r *AliexpressSolutionBatchProductDeleteRequest) SetProductIds(_productIds []int64) error {
    r._productIds = _productIds
    r.Set("product_ids", _productIds)
    return nil
}

// ProductIds Getter
func (r AliexpressSolutionBatchProductDeleteRequest) GetProductIds() []int64 {
    return r._productIds
}

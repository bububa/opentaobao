package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Query the sku attribute information belonged to a specific category API请求
aliexpress.solution.sku.attribute.query

Query the sku attribute information belonged to a specific category, customized for oversea merchants.
*/
type AliexpressSolutionSkuAttributeQueryRequest struct {
    model.Params
    // input parameters
    _querySkuAttributeInfoRequest   *SkuAttributeInfoQueryRequest
}

// 初始化AliexpressSolutionSkuAttributeQueryRequest对象
func NewAliexpressSolutionSkuAttributeQueryRequest() *AliexpressSolutionSkuAttributeQueryRequest{
    return &AliexpressSolutionSkuAttributeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSkuAttributeQueryRequest) GetApiMethodName() string {
    return "aliexpress.solution.sku.attribute.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSkuAttributeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QuerySkuAttributeInfoRequest Setter
// input parameters
func (r *AliexpressSolutionSkuAttributeQueryRequest) SetQuerySkuAttributeInfoRequest(_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest) error {
    r._querySkuAttributeInfoRequest = _querySkuAttributeInfoRequest
    r.Set("query_sku_attribute_info_request", _querySkuAttributeInfoRequest)
    return nil
}

// QuerySkuAttributeInfoRequest Getter
func (r AliexpressSolutionSkuAttributeQueryRequest) GetQuerySkuAttributeInfoRequest() *SkuAttributeInfoQueryRequest {
    return r._querySkuAttributeInfoRequest
}

package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Get product list API请求
aliexpress.solution.product.list.get

Get product list
*/
type AliexpressSolutionProductListGetRequest struct {
    model.Params
    // request parameters to query
    _aeopAEProductListQuery   *ItemListQuery
}

// 初始化AliexpressSolutionProductListGetRequest对象
func NewAliexpressSolutionProductListGetRequest() *AliexpressSolutionProductListGetRequest{
    return &AliexpressSolutionProductListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductListGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AeopAEProductListQuery Setter
// request parameters to query
func (r *AliexpressSolutionProductListGetRequest) SetAeopAEProductListQuery(_aeopAEProductListQuery *ItemListQuery) error {
    r._aeopAEProductListQuery = _aeopAEProductListQuery
    r.Set("aeop_a_e_product_list_query", _aeopAEProductListQuery)
    return nil
}

// AeopAEProductListQuery Getter
func (r AliexpressSolutionProductListGetRequest) GetAeopAEProductListQuery() *ItemListQuery {
    return r._aeopAEProductListQuery
}

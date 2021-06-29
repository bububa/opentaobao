package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Get product list APIRequest
aliexpress.solution.product.list.get

Get product list
*/
type AliexpressSolutionProductListGetRequest struct {
    model.Params

    // request parameters to query
    aeopAEProductListQuery   *ItemListQuery 

}

func NewAliexpressSolutionProductListGetRequest() *AliexpressSolutionProductListGetRequest{
    return &AliexpressSolutionProductListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionProductListGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.list.get"
}

func (r AliexpressSolutionProductListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionProductListGetRequest) SetAeopAEProductListQuery(aeopAEProductListQuery *ItemListQuery) error {
    r.aeopAEProductListQuery = aeopAEProductListQuery
    r.Set("aeop_a_e_product_list_query", aeopAEProductListQuery)
    return nil
}

func (r AliexpressSolutionProductListGetRequest) GetAeopAEProductListQuery() *ItemListQuery {
    return r.aeopAEProductListQuery
}


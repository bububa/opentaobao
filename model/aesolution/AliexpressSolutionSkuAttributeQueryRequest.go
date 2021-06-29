package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Query the sku attribute information belonged to a specific category APIRequest
aliexpress.solution.sku.attribute.query

Query the sku attribute information belonged to a specific category, customized for oversea merchants.
*/
type AliexpressSolutionSkuAttributeQueryRequest struct {
    model.Params

    // input parameters
    querySkuAttributeInfoRequest   *SkuAttributeInfoQueryRequest 

}

func NewAliexpressSolutionSkuAttributeQueryRequest() *AliexpressSolutionSkuAttributeQueryRequest{
    return &AliexpressSolutionSkuAttributeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionSkuAttributeQueryRequest) GetApiMethodName() string {
    return "aliexpress.solution.sku.attribute.query"
}

func (r AliexpressSolutionSkuAttributeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionSkuAttributeQueryRequest) SetQuerySkuAttributeInfoRequest(querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest) error {
    r.querySkuAttributeInfoRequest = querySkuAttributeInfoRequest
    r.Set("query_sku_attribute_info_request", querySkuAttributeInfoRequest)
    return nil
}

func (r AliexpressSolutionSkuAttributeQueryRequest) GetQuerySkuAttributeInfoRequest() *SkuAttributeInfoQueryRequest {
    return r.querySkuAttributeInfoRequest
}


package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
get product schema APIRequest
aliexpress.solution.product.schema.get

provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
*/
type AliexpressSolutionProductSchemaGetRequest struct {
    model.Params

    // aliexpress category id. You can get it from category API
    aliexpressCategoryId   int64 

}

func NewAliexpressSolutionProductSchemaGetRequest() *AliexpressSolutionProductSchemaGetRequest{
    return &AliexpressSolutionProductSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionProductSchemaGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.schema.get"
}

func (r AliexpressSolutionProductSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionProductSchemaGetRequest) SetAliexpressCategoryId(aliexpressCategoryId int64) error {
    r.aliexpressCategoryId = aliexpressCategoryId
    r.Set("aliexpress_category_id", aliexpressCategoryId)
    return nil
}

func (r AliexpressSolutionProductSchemaGetRequest) GetAliexpressCategoryId() int64 {
    return r.aliexpressCategoryId
}


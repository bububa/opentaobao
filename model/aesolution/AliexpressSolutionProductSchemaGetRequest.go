package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
get product schema API请求
aliexpress.solution.product.schema.get

provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
*/
type AliexpressSolutionProductSchemaGetRequest struct {
    model.Params
    // aliexpress category id. You can get it from category API
    _aliexpressCategoryId   int64
}

// 初始化AliexpressSolutionProductSchemaGetRequest对象
func NewAliexpressSolutionProductSchemaGetRequest() *AliexpressSolutionProductSchemaGetRequest{
    return &AliexpressSolutionProductSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductSchemaGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliexpressCategoryId Setter
// aliexpress category id. You can get it from category API
func (r *AliexpressSolutionProductSchemaGetRequest) SetAliexpressCategoryId(_aliexpressCategoryId int64) error {
    r._aliexpressCategoryId = _aliexpressCategoryId
    r.Set("aliexpress_category_id", _aliexpressCategoryId)
    return nil
}

// AliexpressCategoryId Getter
func (r AliexpressSolutionProductSchemaGetRequest) GetAliexpressCategoryId() int64 {
    return r._aliexpressCategoryId
}

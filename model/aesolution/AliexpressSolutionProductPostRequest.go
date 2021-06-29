package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Product posting API API请求
aliexpress.solution.product.post

Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
*/
type AliexpressSolutionProductPostRequest struct {
    model.Params
    // input param
    _postProductRequest   *PostProductRequestDto
}

// 初始化AliexpressSolutionProductPostRequest对象
func NewAliexpressSolutionProductPostRequest() *AliexpressSolutionProductPostRequest{
    return &AliexpressSolutionProductPostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductPostRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.post"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PostProductRequest Setter
// input param
func (r *AliexpressSolutionProductPostRequest) SetPostProductRequest(_postProductRequest *PostProductRequestDto) error {
    r._postProductRequest = _postProductRequest
    r.Set("post_product_request", _postProductRequest)
    return nil
}

// PostProductRequest Getter
func (r AliexpressSolutionProductPostRequest) GetPostProductRequest() *PostProductRequestDto {
    return r._postProductRequest
}

package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Product posting API APIRequest
aliexpress.solution.product.post

Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
*/
type AliexpressSolutionProductPostRequest struct {
    model.Params

    // input param
    postProductRequest   *PostProductRequestDto 

}

func NewAliexpressSolutionProductPostRequest() *AliexpressSolutionProductPostRequest{
    return &AliexpressSolutionProductPostRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionProductPostRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.post"
}

func (r AliexpressSolutionProductPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionProductPostRequest) SetPostProductRequest(postProductRequest *PostProductRequestDto) error {
    r.postProductRequest = postProductRequest
    r.Set("post_product_request", postProductRequest)
    return nil
}

func (r AliexpressSolutionProductPostRequest) GetPostProductRequest() *PostProductRequestDto {
    return r.postProductRequest
}


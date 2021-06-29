package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Edit Product API APIRequest
aliexpress.solution.product.edit

API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
*/
type AliexpressSolutionProductEditRequest struct {
    model.Params

    // input param
    editProductRequest   *PostProductRequestDto 

}

func NewAliexpressSolutionProductEditRequest() *AliexpressSolutionProductEditRequest{
    return &AliexpressSolutionProductEditRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionProductEditRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.edit"
}

func (r AliexpressSolutionProductEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionProductEditRequest) SetEditProductRequest(editProductRequest *PostProductRequestDto) error {
    r.editProductRequest = editProductRequest
    r.Set("edit_product_request", editProductRequest)
    return nil
}

func (r AliexpressSolutionProductEditRequest) GetEditProductRequest() *PostProductRequestDto {
    return r.editProductRequest
}


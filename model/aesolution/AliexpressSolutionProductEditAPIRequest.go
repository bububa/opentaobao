package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Edit Product API API请求
aliexpress.solution.product.edit

API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
*/
type AliexpressSolutionProductEditAPIRequest struct {
    model.Params
    // input param
    _editProductRequest   *PostProductRequestDto
}

// 初始化AliexpressSolutionProductEditAPIRequest对象
func NewAliexpressSolutionProductEditRequest() *AliexpressSolutionProductEditAPIRequest{
    return &AliexpressSolutionProductEditAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductEditAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.product.edit"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductEditAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EditProductRequest Setter
// input param
func (r *AliexpressSolutionProductEditAPIRequest) SetEditProductRequest(_editProductRequest *PostProductRequestDto) error {
    r._editProductRequest = _editProductRequest
    r.Set("edit_product_request", _editProductRequest)
    return nil
}

// EditProductRequest Getter
func (r AliexpressSolutionProductEditAPIRequest) GetEditProductRequest() *PostProductRequestDto {
    return r._editProductRequest
}

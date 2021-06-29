package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Edit Product API APIResponse
aliexpress.solution.product.edit

API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
*/
type AliexpressSolutionProductEditAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionProductEditResponse
}

type AliexpressSolutionProductEditResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_product_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PostItemResponseDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

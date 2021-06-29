package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Single Product Info APIResponse
aliexpress.solution.product.info.get

Get Single Product Info
*/
type AliexpressSolutionProductInfoGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionProductInfoGetResponse
}

type AliexpressSolutionProductInfoGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_product_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *GlobalAeopFindProductResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

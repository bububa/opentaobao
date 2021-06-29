package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
get product schema APIResponse
aliexpress.solution.product.schema.get

provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
*/
type AliexpressSolutionProductSchemaGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionProductSchemaGetResponse
}

type AliexpressSolutionProductSchemaGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_product_schema_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ProductSchemaDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

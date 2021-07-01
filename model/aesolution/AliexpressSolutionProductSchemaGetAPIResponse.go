package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
get product schema API返回值 
aliexpress.solution.product.schema.get

provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
*/
type AliexpressSolutionProductSchemaGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionProductSchemaGetAPIResponseModel
}

// get product schema 成功返回结果
type AliexpressSolutionProductSchemaGetAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_solution_product_schema_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ProductSchemaDto `json:"result,omitempty" xml:"result,omitempty"`
}

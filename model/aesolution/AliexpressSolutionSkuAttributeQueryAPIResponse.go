package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Query the sku attribute information belonged to a specific category API返回值 
aliexpress.solution.sku.attribute.query

Query the sku attribute information belonged to a specific category, customized for oversea merchants.
*/
type AliexpressSolutionSkuAttributeQueryAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionSkuAttributeQueryAPIResponseModel
}

// Query the sku attribute information belonged to a specific category 成功返回结果
type AliexpressSolutionSkuAttributeQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_solution_sku_attribute_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *SkuAttributeInfoQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

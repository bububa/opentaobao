package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Product posting API API返回值 
aliexpress.solution.product.post

Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
*/
type AliexpressSolutionProductPostAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionProductPostResponse
}

// Product posting API 成功返回结果
type AliexpressSolutionProductPostResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_product_post_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PostItemResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}

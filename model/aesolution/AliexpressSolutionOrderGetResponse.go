package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
get order list API返回值 
aliexpress.solution.order.get

Get Order List from AliExpress
*/
type AliexpressSolutionOrderGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionOrderGetResponse
}

// get order list 成功返回结果
type AliexpressSolutionOrderGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_order_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`
}

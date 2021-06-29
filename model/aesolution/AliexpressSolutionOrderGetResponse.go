package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
get order list APIResponse
aliexpress.solution.order.get

Get Order List from AliExpress
*/
type AliexpressSolutionOrderGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionOrderGetResponse
}

type AliexpressSolutionOrderGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

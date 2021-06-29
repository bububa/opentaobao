package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
fulfill order APIResponse
aliexpress.solution.order.fulfill

fulfill order for seller
*/
type AliexpressSolutionOrderFulfillAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionOrderFulfillResponse
}

type AliexpressSolutionOrderFulfillResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_order_fulfill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

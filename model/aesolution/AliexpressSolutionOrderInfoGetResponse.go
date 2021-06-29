package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
get order detail info APIResponse
aliexpress.solution.order.info.get

get order detail info
*/
type AliexpressSolutionOrderInfoGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionOrderInfoGetResponse
}

type AliexpressSolutionOrderInfoGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_order_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

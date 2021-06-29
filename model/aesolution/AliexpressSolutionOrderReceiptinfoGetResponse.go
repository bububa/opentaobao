package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Order Receipt Info APIResponse
aliexpress.solution.order.receiptinfo.get

Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
*/
type AliexpressSolutionOrderReceiptinfoGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionOrderReceiptinfoGetResponse
}

type AliexpressSolutionOrderReceiptinfoGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_order_receiptinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *OrderAddressDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

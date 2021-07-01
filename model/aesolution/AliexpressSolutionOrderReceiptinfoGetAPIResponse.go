package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Get Order Receipt Info API返回值 
aliexpress.solution.order.receiptinfo.get

Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
*/
type AliexpressSolutionOrderReceiptinfoGetAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionOrderReceiptinfoGetAPIResponseModel
}

// Get Order Receipt Info 成功返回结果
type AliexpressSolutionOrderReceiptinfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_solution_order_receiptinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *OrderAddressDto `json:"result,omitempty" xml:"result,omitempty"`
}

package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户可用发票列表 APIResponse
alitrip.btrip.invoice.get

差旅申请用户获取可用发票列表
*/
type AlitripBtripInvoiceGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripInvoiceGetResponse
}

type AlitripBtripInvoiceGetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_invoice_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

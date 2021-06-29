package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据发票抬头搜索发票 APIResponse
alitrip.btrip.invoice.search

用户根据发票抬头搜索发票信息
*/
type AlitripBtripInvoiceSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripInvoiceSearchResponse
}

type AlitripBtripInvoiceSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_invoice_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

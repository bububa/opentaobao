package tax

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方开票回调API APIResponse
alibaba.tax.invoice.sync

该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
*/
type AlibabaTaxInvoiceSyncAPIResponse struct {
    model.CommonResponse
    AlibabaTaxInvoiceSyncResponse
}

type AlibabaTaxInvoiceSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_tax_invoice_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ThirdPartyInvoiceCallBackResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

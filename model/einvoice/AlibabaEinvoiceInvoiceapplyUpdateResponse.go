package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家开票申请单状态回传 APIResponse
alibaba.einvoice.invoiceapply.update

开票服务商更新商家开票申请单状态
*/
type AlibabaEinvoiceInvoiceapplyUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceInvoiceapplyUpdateResponse
}

type AlibabaEinvoiceInvoiceapplyUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_invoiceapply_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}

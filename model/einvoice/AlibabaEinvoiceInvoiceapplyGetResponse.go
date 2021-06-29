package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的开票申请 APIResponse
alibaba.einvoice.invoiceapply.get

开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
*/
type AlibabaEinvoiceInvoiceapplyGetAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceInvoiceapplyGetResponse
}

type AlibabaEinvoiceInvoiceapplyGetResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_invoiceapply_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 开票申请查询结果
    
    Result   *UserInvoiceApplyDTO `json:"result,omitempty" xml:"result,omitempty"`

    
    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}

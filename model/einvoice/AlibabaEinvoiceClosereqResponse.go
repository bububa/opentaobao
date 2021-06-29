package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭开票失败请求（失败列表可重试） APIResponse
alibaba.einvoice.closereq

关闭失败开票请求，避免造成重复开票
*/
type AlibabaEinvoiceClosereqAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceClosereqResponse
}

type AlibabaEinvoiceClosereqResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_closereq_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 关闭是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}

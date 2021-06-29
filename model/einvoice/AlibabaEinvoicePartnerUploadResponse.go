package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商发票上传接口（非授权） APIResponse
alibaba.einvoice.partner.upload

服务商发票上传接口（非授权）
*/
type AlibabaEinvoicePartnerUploadAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoicePartnerUploadResponse
}

type AlibabaEinvoicePartnerUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_partner_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 上传结果
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}

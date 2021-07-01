package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商发票上传接口（非授权） API返回值 
alibaba.einvoice.partner.upload

服务商发票上传接口（非授权）
*/
type AlibabaEinvoicePartnerUploadAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoicePartnerUploadAPIResponseModel
}

// 服务商发票上传接口（非授权） 成功返回结果
type AlibabaEinvoicePartnerUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_einvoice_partner_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 上传结果
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

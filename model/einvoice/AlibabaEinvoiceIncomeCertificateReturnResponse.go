package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传进项认证结果 APIResponse
alibaba.einvoice.income.certificate.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪
*/
type AlibabaEinvoiceIncomeCertificateReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeCertificateReturnResponse
}

type AlibabaEinvoiceIncomeCertificateReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_certificate_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}

package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeCertificateReturnAPIResponse 服务商回传进项认证结果 API返回值
// alibaba.einvoice.income.certificate.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
type AlibabaEinvoiceIncomeCertificateReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceIncomeCertificateReturnAPIResponseModel
}

// AlibabaEinvoiceIncomeCertificateReturnAPIResponseModel is 服务商回传进项认证结果 成功返回结果
type AlibabaEinvoiceIncomeCertificateReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_income_certificate_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

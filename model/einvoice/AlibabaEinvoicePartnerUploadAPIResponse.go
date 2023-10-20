package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePartnerUploadAPIResponse 服务商发票上传接口（非授权） API返回值
// alibaba.einvoice.partner.upload
//
// 服务商发票上传接口（非授权）
type AlibabaEinvoicePartnerUploadAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePartnerUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePartnerUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePartnerUploadAPIResponseModel).Reset()
}

// AlibabaEinvoicePartnerUploadAPIResponseModel is 服务商发票上传接口（非授权） 成功返回结果
type AlibabaEinvoicePartnerUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_partner_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePartnerUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoicePartnerUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePartnerUploadAPIResponse)
	},
}

// GetAlibabaEinvoicePartnerUploadAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePartnerUploadAPIResponse
func GetAlibabaEinvoicePartnerUploadAPIResponse() *AlibabaEinvoicePartnerUploadAPIResponse {
	return poolAlibabaEinvoicePartnerUploadAPIResponse.Get().(*AlibabaEinvoicePartnerUploadAPIResponse)
}

// ReleaseAlibabaEinvoicePartnerUploadAPIResponse 将 AlibabaEinvoicePartnerUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePartnerUploadAPIResponse(v *AlibabaEinvoicePartnerUploadAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePartnerUploadAPIResponse.Put(v)
}

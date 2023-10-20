package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthStoreCertificateCreateAPIResponse 仓库换证审批 API返回值
// alibaba.alihealth.store.certificate.create
//
// 仓库侧换证发起审批
type AlibabaAlihealthStoreCertificateCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthStoreCertificateCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthStoreCertificateCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthStoreCertificateCreateAPIResponseModel).Reset()
}

// AlibabaAlihealthStoreCertificateCreateAPIResponseModel is 仓库换证审批 成功返回结果
type AlibabaAlihealthStoreCertificateCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_store_certificate_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaAlihealthStoreCertificateCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthStoreCertificateCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthStoreCertificateCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthStoreCertificateCreateAPIResponse)
	},
}

// GetAlibabaAlihealthStoreCertificateCreateAPIResponse 从 sync.Pool 获取 AlibabaAlihealthStoreCertificateCreateAPIResponse
func GetAlibabaAlihealthStoreCertificateCreateAPIResponse() *AlibabaAlihealthStoreCertificateCreateAPIResponse {
	return poolAlibabaAlihealthStoreCertificateCreateAPIResponse.Get().(*AlibabaAlihealthStoreCertificateCreateAPIResponse)
}

// ReleaseAlibabaAlihealthStoreCertificateCreateAPIResponse 将 AlibabaAlihealthStoreCertificateCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthStoreCertificateCreateAPIResponse(v *AlibabaAlihealthStoreCertificateCreateAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthStoreCertificateCreateAPIResponse.Put(v)
}

package retail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailElectronicCertificateConfirmAPIResponse 确认核销接口 API返回值
// alibaba.retail.electronic.certificate.confirm
//
// 确认核销接口
type AlibabaRetailElectronicCertificateConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaRetailElectronicCertificateConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailElectronicCertificateConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailElectronicCertificateConfirmAPIResponseModel).Reset()
}

// AlibabaRetailElectronicCertificateConfirmAPIResponseModel is 确认核销接口 成功返回结果
type AlibabaRetailElectronicCertificateConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_electronic_certificate_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailElectronicCertificateConfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailElectronicCertificateConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailElectronicCertificateConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailElectronicCertificateConfirmAPIResponse)
	},
}

// GetAlibabaRetailElectronicCertificateConfirmAPIResponse 从 sync.Pool 获取 AlibabaRetailElectronicCertificateConfirmAPIResponse
func GetAlibabaRetailElectronicCertificateConfirmAPIResponse() *AlibabaRetailElectronicCertificateConfirmAPIResponse {
	return poolAlibabaRetailElectronicCertificateConfirmAPIResponse.Get().(*AlibabaRetailElectronicCertificateConfirmAPIResponse)
}

// ReleaseAlibabaRetailElectronicCertificateConfirmAPIResponse 将 AlibabaRetailElectronicCertificateConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailElectronicCertificateConfirmAPIResponse(v *AlibabaRetailElectronicCertificateConfirmAPIResponse) {
	v.Reset()
	poolAlibabaRetailElectronicCertificateConfirmAPIResponse.Put(v)
}

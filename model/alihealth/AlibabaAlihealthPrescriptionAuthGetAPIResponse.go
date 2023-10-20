package alihealth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPrescriptionAuthGetAPIResponse 阿里健康处方平台获取授权码 API返回值
// alibaba.alihealth.prescription.auth.get
//
// 获取处方用户授权
type AlibabaAlihealthPrescriptionAuthGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPrescriptionAuthGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPrescriptionAuthGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPrescriptionAuthGetAPIResponseModel).Reset()
}

// AlibabaAlihealthPrescriptionAuthGetAPIResponseModel is 阿里健康处方平台获取授权码 成功返回结果
type AlibabaAlihealthPrescriptionAuthGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_prescription_auth_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPrescriptionAuthGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthPrescriptionAuthGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPrescriptionAuthGetAPIResponse)
	},
}

// GetAlibabaAlihealthPrescriptionAuthGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPrescriptionAuthGetAPIResponse
func GetAlibabaAlihealthPrescriptionAuthGetAPIResponse() *AlibabaAlihealthPrescriptionAuthGetAPIResponse {
	return poolAlibabaAlihealthPrescriptionAuthGetAPIResponse.Get().(*AlibabaAlihealthPrescriptionAuthGetAPIResponse)
}

// ReleaseAlibabaAlihealthPrescriptionAuthGetAPIResponse 将 AlibabaAlihealthPrescriptionAuthGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPrescriptionAuthGetAPIResponse(v *AlibabaAlihealthPrescriptionAuthGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPrescriptionAuthGetAPIResponse.Put(v)
}

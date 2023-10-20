package alihealthmedical

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalItemPublishAPIResponse 三方入驻-开通服务 API返回值
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
type AlibabaAlihealthMedicalItemPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalItemPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalItemPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalItemPublishAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalItemPublishAPIResponseModel is 三方入驻-开通服务 成功返回结果
type AlibabaAlihealthMedicalItemPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_item_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统返回的通用结果类
	Result1 *ServiceResult `json:"result1,omitempty" xml:"result1,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalItemPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result1 = nil
}

var poolAlibabaAlihealthMedicalItemPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalItemPublishAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalItemPublishAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalItemPublishAPIResponse
func GetAlibabaAlihealthMedicalItemPublishAPIResponse() *AlibabaAlihealthMedicalItemPublishAPIResponse {
	return poolAlibabaAlihealthMedicalItemPublishAPIResponse.Get().(*AlibabaAlihealthMedicalItemPublishAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalItemPublishAPIResponse 将 AlibabaAlihealthMedicalItemPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalItemPublishAPIResponse(v *AlibabaAlihealthMedicalItemPublishAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalItemPublishAPIResponse.Put(v)
}

package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse 三方评论信息同步 API返回值
// alibaba.alihealth.medicalbase.third.evaluate.sync
//
// 三方评论信息同步
type AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel is 三方评论信息同步 成功返回结果
type AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_third_evaluate_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse
func GetAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse() *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse {
	return poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse.Get().(*AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse 将 AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse(v *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse.Put(v)
}

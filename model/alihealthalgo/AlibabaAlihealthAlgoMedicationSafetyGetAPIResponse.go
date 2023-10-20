package alihealthalgo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse 合理用药api API返回值
// alibaba.alihealth.algo.medication.safety.get
//
// 合理用药规则引擎服务
type AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel).Reset()
}

// AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel is 合理用药api 成功返回结果
type AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_algo_medication_safety_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *SolutionResultTopSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse)
	},
}

// GetAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse
func GetAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse() *AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse {
	return poolAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse.Get().(*AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse)
}

// ReleaseAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse 将 AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse(v *AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthAlgoMedicationSafetyGetAPIResponse.Put(v)
}

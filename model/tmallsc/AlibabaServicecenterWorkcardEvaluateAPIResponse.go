package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardEvaluateAPIResponse 服务商售后鉴定服务 API返回值
// alibaba.servicecenter.workcard.evaluate
//
// 服务商售后鉴定服务,提供给服务商针对售后场景上门鉴定服务，鉴定成功则服务商完成履约，鉴定失败则取消工单
type AlibabaServicecenterWorkcardEvaluateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardEvaluateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardEvaluateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardEvaluateAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardEvaluateAPIResponseModel is 服务商售后鉴定服务 成功返回结果
type AlibabaServicecenterWorkcardEvaluateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_evaluate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	SscResult *AlibabaServicecenterWorkcardEvaluateResult `json:"ssc_result,omitempty" xml:"ssc_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardEvaluateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SscResult = nil
}

var poolAlibabaServicecenterWorkcardEvaluateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardEvaluateAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardEvaluateAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardEvaluateAPIResponse
func GetAlibabaServicecenterWorkcardEvaluateAPIResponse() *AlibabaServicecenterWorkcardEvaluateAPIResponse {
	return poolAlibabaServicecenterWorkcardEvaluateAPIResponse.Get().(*AlibabaServicecenterWorkcardEvaluateAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardEvaluateAPIResponse 将 AlibabaServicecenterWorkcardEvaluateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardEvaluateAPIResponse(v *AlibabaServicecenterWorkcardEvaluateAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardEvaluateAPIResponse.Put(v)
}

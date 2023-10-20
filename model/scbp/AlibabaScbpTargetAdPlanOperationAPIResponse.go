package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanOperationAPIResponse 定向推广-计划开启/暂停/删除 API返回值
// alibaba.scbp.target.ad.plan.operation
//
// 定向推广-计划开启/暂停/删除
type AlibabaScbpTargetAdPlanOperationAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanOperationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanOperationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanOperationAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanOperationAPIResponseModel is 定向推广-计划开启/暂停/删除 成功返回结果
type AlibabaScbpTargetAdPlanOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改成功记录数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanOperationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpTargetAdPlanOperationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanOperationAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanOperationAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanOperationAPIResponse
func GetAlibabaScbpTargetAdPlanOperationAPIResponse() *AlibabaScbpTargetAdPlanOperationAPIResponse {
	return poolAlibabaScbpTargetAdPlanOperationAPIResponse.Get().(*AlibabaScbpTargetAdPlanOperationAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanOperationAPIResponse 将 AlibabaScbpTargetAdPlanOperationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanOperationAPIResponse(v *AlibabaScbpTargetAdPlanOperationAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanOperationAPIResponse.Put(v)
}

package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse 物流管控作业状态回传 API返回值
// alibaba.wdk.fulfill.batch.on.task.status.changed
//
// 物流管控作业状态回传
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponseModel).Reset()
}

// AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponseModel is 物流管控作业状态回传 成功返回结果
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_batch_on_task_status_changed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// fulfillLogisticSingleResult
	FulfillLogisticSingleResult *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty" xml:"fulfill_logistic_single_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FulfillLogisticSingleResult = nil
}

var poolAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse)
	},
}

// GetAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse
func GetAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse() *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse {
	return poolAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse.Get().(*AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse)
}

// ReleaseAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse 将 AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse(v *AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse.Put(v)
}

package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse 作业小票查询接口 API返回值
// alibaba.wdk.fulfill.batch.query.by.batchids
//
// 根据节点等条件查询履约单小票信息
type AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillBatchQueryByBatchidsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillBatchQueryByBatchidsAPIResponseModel).Reset()
}

// AlibabaWdkFulfillBatchQueryByBatchidsAPIResponseModel is 作业小票查询接口 成功返回结果
type AlibabaWdkFulfillBatchQueryByBatchidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_batch_query_by_batchids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果对象
	FulfillLogisticListResult *FulfillLogisticListResult `json:"fulfill_logistic_list_result,omitempty" xml:"fulfill_logistic_list_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBatchQueryByBatchidsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FulfillLogisticListResult = nil
}

var poolAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse)
	},
}

// GetAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse
func GetAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse() *AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse {
	return poolAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse.Get().(*AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse)
}

// ReleaseAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse 将 AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse(v *AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillBatchQueryByBatchidsAPIResponse.Put(v)
}

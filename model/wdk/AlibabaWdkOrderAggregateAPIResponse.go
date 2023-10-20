package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderAggregateAPIResponse 淘鲜达订单按门店机台号聚合查询 API返回值
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
type AlibabaWdkOrderAggregateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderAggregateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOrderAggregateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOrderAggregateAPIResponseModel).Reset()
}

// AlibabaWdkOrderAggregateAPIResponseModel is 淘鲜达订单按门店机台号聚合查询 成功返回结果
type AlibabaWdkOrderAggregateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_aggregate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderAggregateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOrderAggregateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkOrderAggregateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderAggregateAPIResponse)
	},
}

// GetAlibabaWdkOrderAggregateAPIResponse 从 sync.Pool 获取 AlibabaWdkOrderAggregateAPIResponse
func GetAlibabaWdkOrderAggregateAPIResponse() *AlibabaWdkOrderAggregateAPIResponse {
	return poolAlibabaWdkOrderAggregateAPIResponse.Get().(*AlibabaWdkOrderAggregateAPIResponse)
}

// ReleaseAlibabaWdkOrderAggregateAPIResponse 将 AlibabaWdkOrderAggregateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOrderAggregateAPIResponse(v *AlibabaWdkOrderAggregateAPIResponse) {
	v.Reset()
	poolAlibabaWdkOrderAggregateAPIResponse.Put(v)
}

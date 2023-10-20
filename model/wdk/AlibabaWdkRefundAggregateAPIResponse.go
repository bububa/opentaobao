package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkRefundAggregateAPIResponse 淘鲜达退款单按门店聚合查询 API返回值
// alibaba.wdk.refund.aggregate
//
// 淘鲜达退款单按门店聚合查询
type AlibabaWdkRefundAggregateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkRefundAggregateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkRefundAggregateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkRefundAggregateAPIResponseModel).Reset()
}

// AlibabaWdkRefundAggregateAPIResponseModel is 淘鲜达退款单按门店聚合查询 成功返回结果
type AlibabaWdkRefundAggregateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_refund_aggregate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundAggregateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkRefundAggregateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkRefundAggregateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkRefundAggregateAPIResponse)
	},
}

// GetAlibabaWdkRefundAggregateAPIResponse 从 sync.Pool 获取 AlibabaWdkRefundAggregateAPIResponse
func GetAlibabaWdkRefundAggregateAPIResponse() *AlibabaWdkRefundAggregateAPIResponse {
	return poolAlibabaWdkRefundAggregateAPIResponse.Get().(*AlibabaWdkRefundAggregateAPIResponse)
}

// ReleaseAlibabaWdkRefundAggregateAPIResponse 将 AlibabaWdkRefundAggregateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkRefundAggregateAPIResponse(v *AlibabaWdkRefundAggregateAPIResponse) {
	v.Reset()
	poolAlibabaWdkRefundAggregateAPIResponse.Put(v)
}

package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderRefundListAPIResponse 五道口交易退款批量查询 API返回值
// alibaba.wdk.order.refund.list
//
// 按照条件查询退款数据
type AlibabaWdkOrderRefundListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderRefundListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOrderRefundListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOrderRefundListAPIResponseModel).Reset()
}

// AlibabaWdkOrderRefundListAPIResponseModel is 五道口交易退款批量查询 成功返回结果
type AlibabaWdkOrderRefundListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_refund_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果内容
	Result *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOrderRefundListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkOrderRefundListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderRefundListAPIResponse)
	},
}

// GetAlibabaWdkOrderRefundListAPIResponse 从 sync.Pool 获取 AlibabaWdkOrderRefundListAPIResponse
func GetAlibabaWdkOrderRefundListAPIResponse() *AlibabaWdkOrderRefundListAPIResponse {
	return poolAlibabaWdkOrderRefundListAPIResponse.Get().(*AlibabaWdkOrderRefundListAPIResponse)
}

// ReleaseAlibabaWdkOrderRefundListAPIResponse 将 AlibabaWdkOrderRefundListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOrderRefundListAPIResponse(v *AlibabaWdkOrderRefundListAPIResponse) {
	v.Reset()
	poolAlibabaWdkOrderRefundListAPIResponse.Put(v)
}

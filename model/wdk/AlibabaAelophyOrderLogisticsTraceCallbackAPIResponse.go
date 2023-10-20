package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse 配送轨迹回传 API返回值
// alibaba.aelophy.order.logistics.trace.callback
//
// 配送轨迹回传
type AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyOrderLogisticsTraceCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyOrderLogisticsTraceCallbackAPIResponseModel).Reset()
}

// AlibabaAelophyOrderLogisticsTraceCallbackAPIResponseModel is 配送轨迹回传 成功返回结果
type AlibabaAelophyOrderLogisticsTraceCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_logistics_trace_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 配送轨迹回传响应
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderLogisticsTraceCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse)
	},
}

// GetAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse 从 sync.Pool 获取 AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse
func GetAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse() *AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse {
	return poolAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse.Get().(*AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse)
}

// ReleaseAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse 将 AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse(v *AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse) {
	v.Reset()
	poolAlibabaAelophyOrderLogisticsTraceCallbackAPIResponse.Put(v)
}

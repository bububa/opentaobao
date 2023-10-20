package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderSuccessCreateAPIResponse 五道口终态订单创建 API返回值
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
type AlibabaWdkTradeOrderSuccessCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeOrderSuccessCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderSuccessCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeOrderSuccessCreateAPIResponseModel).Reset()
}

// AlibabaWdkTradeOrderSuccessCreateAPIResponseModel is 五道口终态订单创建 成功返回结果
type AlibabaWdkTradeOrderSuccessCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_success_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单返回结果
	OrderResult *OrderQueryResult `json:"order_result,omitempty" xml:"order_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderSuccessCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderResult = nil
}

var poolAlibabaWdkTradeOrderSuccessCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeOrderSuccessCreateAPIResponse)
	},
}

// GetAlibabaWdkTradeOrderSuccessCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeOrderSuccessCreateAPIResponse
func GetAlibabaWdkTradeOrderSuccessCreateAPIResponse() *AlibabaWdkTradeOrderSuccessCreateAPIResponse {
	return poolAlibabaWdkTradeOrderSuccessCreateAPIResponse.Get().(*AlibabaWdkTradeOrderSuccessCreateAPIResponse)
}

// ReleaseAlibabaWdkTradeOrderSuccessCreateAPIResponse 将 AlibabaWdkTradeOrderSuccessCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeOrderSuccessCreateAPIResponse(v *AlibabaWdkTradeOrderSuccessCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeOrderSuccessCreateAPIResponse.Put(v)
}

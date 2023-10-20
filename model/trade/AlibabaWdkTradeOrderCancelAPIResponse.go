package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderCancelAPIResponse 外部交易订单取消接口 API返回值
// alibaba.wdk.trade.order.cancel
//
// 通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeOrderCancelAPIResponseModel).Reset()
}

// AlibabaWdkTradeOrderCancelAPIResponseModel is 外部交易订单取消接口 成功返回结果
type AlibabaWdkTradeOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result *OrderResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkTradeOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeOrderCancelAPIResponse)
	},
}

// GetAlibabaWdkTradeOrderCancelAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeOrderCancelAPIResponse
func GetAlibabaWdkTradeOrderCancelAPIResponse() *AlibabaWdkTradeOrderCancelAPIResponse {
	return poolAlibabaWdkTradeOrderCancelAPIResponse.Get().(*AlibabaWdkTradeOrderCancelAPIResponse)
}

// ReleaseAlibabaWdkTradeOrderCancelAPIResponse 将 AlibabaWdkTradeOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeOrderCancelAPIResponse(v *AlibabaWdkTradeOrderCancelAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeOrderCancelAPIResponse.Put(v)
}

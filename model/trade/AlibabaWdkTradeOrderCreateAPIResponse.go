package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderCreateAPIResponse 外部交易订单创单接口 API返回值
// alibaba.wdk.trade.order.create
//
// 通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeOrderCreateAPIResponseModel).Reset()
}

// AlibabaWdkTradeOrderCreateAPIResponseModel is 外部交易订单创单接口 成功返回结果
type AlibabaWdkTradeOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result *OrderResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkTradeOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeOrderCreateAPIResponse)
	},
}

// GetAlibabaWdkTradeOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeOrderCreateAPIResponse
func GetAlibabaWdkTradeOrderCreateAPIResponse() *AlibabaWdkTradeOrderCreateAPIResponse {
	return poolAlibabaWdkTradeOrderCreateAPIResponse.Get().(*AlibabaWdkTradeOrderCreateAPIResponse)
}

// ReleaseAlibabaWdkTradeOrderCreateAPIResponse 将 AlibabaWdkTradeOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeOrderCreateAPIResponse(v *AlibabaWdkTradeOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeOrderCreateAPIResponse.Put(v)
}

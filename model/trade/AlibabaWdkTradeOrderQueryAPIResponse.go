package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderQueryAPIResponse 查询外部交易订单接口 API返回值
// alibaba.wdk.trade.order.query
//
// 通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeOrderQueryAPIResponseModel).Reset()
}

// AlibabaWdkTradeOrderQueryAPIResponseModel is 查询外部交易订单接口 成功返回结果
type AlibabaWdkTradeOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *TradeOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkTradeOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeOrderQueryAPIResponse)
	},
}

// GetAlibabaWdkTradeOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeOrderQueryAPIResponse
func GetAlibabaWdkTradeOrderQueryAPIResponse() *AlibabaWdkTradeOrderQueryAPIResponse {
	return poolAlibabaWdkTradeOrderQueryAPIResponse.Get().(*AlibabaWdkTradeOrderQueryAPIResponse)
}

// ReleaseAlibabaWdkTradeOrderQueryAPIResponse 将 AlibabaWdkTradeOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeOrderQueryAPIResponse(v *AlibabaWdkTradeOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeOrderQueryAPIResponse.Put(v)
}

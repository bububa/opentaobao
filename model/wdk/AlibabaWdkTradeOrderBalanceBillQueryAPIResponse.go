package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderBalanceBillQueryAPIResponse 分页拉取订单数据 API返回值
// alibaba.wdk.trade.order.balance.bill.query
//
// 提供接口供外部调用，分页拉取订单数据
type AlibabaWdkTradeOrderBalanceBillQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeOrderBalanceBillQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderBalanceBillQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeOrderBalanceBillQueryAPIResponseModel).Reset()
}

// AlibabaWdkTradeOrderBalanceBillQueryAPIResponseModel is 分页拉取订单数据 成功返回结果
type AlibabaWdkTradeOrderBalanceBillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_balance_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ApiResult
	ApiResult *AlibabaWdkTradeOrderBalanceBillQueryApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeOrderBalanceBillQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkTradeOrderBalanceBillQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeOrderBalanceBillQueryAPIResponse)
	},
}

// GetAlibabaWdkTradeOrderBalanceBillQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeOrderBalanceBillQueryAPIResponse
func GetAlibabaWdkTradeOrderBalanceBillQueryAPIResponse() *AlibabaWdkTradeOrderBalanceBillQueryAPIResponse {
	return poolAlibabaWdkTradeOrderBalanceBillQueryAPIResponse.Get().(*AlibabaWdkTradeOrderBalanceBillQueryAPIResponse)
}

// ReleaseAlibabaWdkTradeOrderBalanceBillQueryAPIResponse 将 AlibabaWdkTradeOrderBalanceBillQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeOrderBalanceBillQueryAPIResponse(v *AlibabaWdkTradeOrderBalanceBillQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeOrderBalanceBillQueryAPIResponse.Put(v)
}

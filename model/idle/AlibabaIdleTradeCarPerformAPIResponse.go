package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTradeCarPerformAPIResponse 二手车寄卖履约接口 API返回值
// alibaba.idle.trade.car.perform
//
// 二手车寄卖履约接口
type AlibabaIdleTradeCarPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTradeCarPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTradeCarPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTradeCarPerformAPIResponseModel).Reset()
}

// AlibabaIdleTradeCarPerformAPIResponseModel is 二手车寄卖履约接口 成功返回结果
type AlibabaIdleTradeCarPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_trade_car_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CarConsignmentResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTradeCarPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTradeCarPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTradeCarPerformAPIResponse)
	},
}

// GetAlibabaIdleTradeCarPerformAPIResponse 从 sync.Pool 获取 AlibabaIdleTradeCarPerformAPIResponse
func GetAlibabaIdleTradeCarPerformAPIResponse() *AlibabaIdleTradeCarPerformAPIResponse {
	return poolAlibabaIdleTradeCarPerformAPIResponse.Get().(*AlibabaIdleTradeCarPerformAPIResponse)
}

// ReleaseAlibabaIdleTradeCarPerformAPIResponse 将 AlibabaIdleTradeCarPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTradeCarPerformAPIResponse(v *AlibabaIdleTradeCarPerformAPIResponse) {
	v.Reset()
	poolAlibabaIdleTradeCarPerformAPIResponse.Put(v)
}

package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeRefundAPIResponse 退款 API返回值
// alibaba.mos.onsite.trade.refund
//
// 当交易发生之后一段时间内，由于消费者或者商户的原因需退款，商户可通过退款接口将支付款退还给消费者，喵街将在收到退款请求并验证成功后，按退款规则将支付款按原路退到消费者账号上。
//
// 1. 交易超过可退款时间（签约时设置的可退款时间）的订单无法进行退款。
// 2. 只支持全额退款。
type AlibabaMosOnsiteTradeRefundAPIResponse struct {
	model.CommonResponse
	AlibabaMosOnsiteTradeRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOnsiteTradeRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOnsiteTradeRefundAPIResponseModel).Reset()
}

// AlibabaMosOnsiteTradeRefundAPIResponseModel is 退款 成功返回结果
type AlibabaMosOnsiteTradeRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_onsite_trade_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易退款响应
	Result *AlibabaMosOnsiteTradeRefundResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOnsiteTradeRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosOnsiteTradeRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOnsiteTradeRefundAPIResponse)
	},
}

// GetAlibabaMosOnsiteTradeRefundAPIResponse 从 sync.Pool 获取 AlibabaMosOnsiteTradeRefundAPIResponse
func GetAlibabaMosOnsiteTradeRefundAPIResponse() *AlibabaMosOnsiteTradeRefundAPIResponse {
	return poolAlibabaMosOnsiteTradeRefundAPIResponse.Get().(*AlibabaMosOnsiteTradeRefundAPIResponse)
}

// ReleaseAlibabaMosOnsiteTradeRefundAPIResponse 将 AlibabaMosOnsiteTradeRefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOnsiteTradeRefundAPIResponse(v *AlibabaMosOnsiteTradeRefundAPIResponse) {
	v.Reset()
	poolAlibabaMosOnsiteTradeRefundAPIResponse.Put(v)
}

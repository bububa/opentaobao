package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarEticketConsumeAPIResponse 天猫汽车电子凭证核销 API返回值
// tmall.aliauto.trade.car.eticket.consume
//
// 为商家提供电子凭证核销接口，支持分账
type TmallAliautoTradeCarEticketConsumeAPIResponse struct {
	model.CommonResponse
	TmallAliautoTradeCarEticketConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoTradeCarEticketConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoTradeCarEticketConsumeAPIResponseModel).Reset()
}

// TmallAliautoTradeCarEticketConsumeAPIResponseModel is 天猫汽车电子凭证核销 成功返回结果
type TmallAliautoTradeCarEticketConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_car_eticket_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoTradeCarEticketConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoTradeCarEticketConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoTradeCarEticketConsumeAPIResponse)
	},
}

// GetTmallAliautoTradeCarEticketConsumeAPIResponse 从 sync.Pool 获取 TmallAliautoTradeCarEticketConsumeAPIResponse
func GetTmallAliautoTradeCarEticketConsumeAPIResponse() *TmallAliautoTradeCarEticketConsumeAPIResponse {
	return poolTmallAliautoTradeCarEticketConsumeAPIResponse.Get().(*TmallAliautoTradeCarEticketConsumeAPIResponse)
}

// ReleaseTmallAliautoTradeCarEticketConsumeAPIResponse 将 TmallAliautoTradeCarEticketConsumeAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoTradeCarEticketConsumeAPIResponse(v *TmallAliautoTradeCarEticketConsumeAPIResponse) {
	v.Reset()
	poolTmallAliautoTradeCarEticketConsumeAPIResponse.Put(v)
}

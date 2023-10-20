package alitripcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderAgentCancelAPIResponse 司机或客服取消订单 API返回值
// taobao.alitrip.car.order.agent.cancel
//
// 司机或客服取消订单后通知飞猪订单取消
type TaobaoAlitripCarOrderAgentCancelAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderAgentCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderAgentCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderAgentCancelAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderAgentCancelAPIResponseModel is 司机或客服取消订单 成功返回结果
type TaobaoAlitripCarOrderAgentCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_agent_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
	// 结果对象
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderAgentCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MessageCode = 0
	m.Model = false
}

var poolTaobaoAlitripCarOrderAgentCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderAgentCancelAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderAgentCancelAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderAgentCancelAPIResponse
func GetTaobaoAlitripCarOrderAgentCancelAPIResponse() *TaobaoAlitripCarOrderAgentCancelAPIResponse {
	return poolTaobaoAlitripCarOrderAgentCancelAPIResponse.Get().(*TaobaoAlitripCarOrderAgentCancelAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderAgentCancelAPIResponse 将 TaobaoAlitripCarOrderAgentCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderAgentCancelAPIResponse(v *TaobaoAlitripCarOrderAgentCancelAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderAgentCancelAPIResponse.Put(v)
}

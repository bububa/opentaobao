package alitripcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderRefundAPIResponse 用户投诉达成一致后给用户退款 API返回值
// taobao.alitrip.car.order.refund
//
// 用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
type TaobaoAlitripCarOrderRefundAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderRefundAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderRefundAPIResponseModel is 用户投诉达成一致后给用户退款 成功返回结果
type TaobaoAlitripCarOrderRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_refund_response"`
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
func (m *TaobaoAlitripCarOrderRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MessageCode = 0
	m.Model = false
}

var poolTaobaoAlitripCarOrderRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderRefundAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderRefundAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderRefundAPIResponse
func GetTaobaoAlitripCarOrderRefundAPIResponse() *TaobaoAlitripCarOrderRefundAPIResponse {
	return poolTaobaoAlitripCarOrderRefundAPIResponse.Get().(*TaobaoAlitripCarOrderRefundAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderRefundAPIResponse 将 TaobaoAlitripCarOrderRefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderRefundAPIResponse(v *TaobaoAlitripCarOrderRefundAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderRefundAPIResponse.Put(v)
}

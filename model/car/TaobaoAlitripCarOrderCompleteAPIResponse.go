package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderCompleteAPIResponse 服务完成API API返回值
// taobao.alitrip.car.order.complete
//
// 用来接收服务商订单流程完成信息
type TaobaoAlitripCarOrderCompleteAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderCompleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderCompleteAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderCompleteAPIResponseModel is 服务完成API 成功返回结果
type TaobaoAlitripCarOrderCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 其它数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.MessageCode = 0
}

var poolTaobaoAlitripCarOrderCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderCompleteAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderCompleteAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderCompleteAPIResponse
func GetTaobaoAlitripCarOrderCompleteAPIResponse() *TaobaoAlitripCarOrderCompleteAPIResponse {
	return poolTaobaoAlitripCarOrderCompleteAPIResponse.Get().(*TaobaoAlitripCarOrderCompleteAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderCompleteAPIResponse 将 TaobaoAlitripCarOrderCompleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderCompleteAPIResponse(v *TaobaoAlitripCarOrderCompleteAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderCompleteAPIResponse.Put(v)
}

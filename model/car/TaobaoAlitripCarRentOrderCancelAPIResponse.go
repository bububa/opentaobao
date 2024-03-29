package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarRentOrderCancelAPIResponse 租车-取消订单 API返回值
// taobao.alitrip.car.rent.order.cancel
//
// 服务商主动取消用户订单或者拒绝取消订单.
type TaobaoAlitripCarRentOrderCancelAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarRentOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarRentOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarRentOrderCancelAPIResponseModel).Reset()
}

// TaobaoAlitripCarRentOrderCancelAPIResponseModel is 租车-取消订单 成功返回结果
type TaobaoAlitripCarRentOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_rent_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 扩展对象
	Models string `json:"models,omitempty" xml:"models,omitempty"`
	// 结果对象
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 结果码
	C int64 `json:"c,omitempty" xml:"c,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarRentOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Models = ""
	m.Model = ""
	m.C = 0
}

var poolTaobaoAlitripCarRentOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarRentOrderCancelAPIResponse)
	},
}

// GetTaobaoAlitripCarRentOrderCancelAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarRentOrderCancelAPIResponse
func GetTaobaoAlitripCarRentOrderCancelAPIResponse() *TaobaoAlitripCarRentOrderCancelAPIResponse {
	return poolTaobaoAlitripCarRentOrderCancelAPIResponse.Get().(*TaobaoAlitripCarRentOrderCancelAPIResponse)
}

// ReleaseTaobaoAlitripCarRentOrderCancelAPIResponse 将 TaobaoAlitripCarRentOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarRentOrderCancelAPIResponse(v *TaobaoAlitripCarRentOrderCancelAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarRentOrderCancelAPIResponse.Put(v)
}

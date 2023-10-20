package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightFlightchangeOrderQueryAPIResponse 订单维度航变查询 API返回值
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
type AlitripFlightFlightchangeOrderQueryAPIResponse struct {
	model.CommonResponse
	AlitripFlightFlightchangeOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightFlightchangeOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightFlightchangeOrderQueryAPIResponseModel).Reset()
}

// AlitripFlightFlightchangeOrderQueryAPIResponseModel is 订单维度航变查询 成功返回结果
type AlitripFlightFlightchangeOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_flightchange_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 航变信息查询结果
	Result []FlightChangeOrderDto `json:"result,omitempty" xml:"result>flight_change_order_dto,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用成功标志
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightFlightchangeOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolAlitripFlightFlightchangeOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightFlightchangeOrderQueryAPIResponse)
	},
}

// GetAlitripFlightFlightchangeOrderQueryAPIResponse 从 sync.Pool 获取 AlitripFlightFlightchangeOrderQueryAPIResponse
func GetAlitripFlightFlightchangeOrderQueryAPIResponse() *AlitripFlightFlightchangeOrderQueryAPIResponse {
	return poolAlitripFlightFlightchangeOrderQueryAPIResponse.Get().(*AlitripFlightFlightchangeOrderQueryAPIResponse)
}

// ReleaseAlitripFlightFlightchangeOrderQueryAPIResponse 将 AlitripFlightFlightchangeOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightFlightchangeOrderQueryAPIResponse(v *AlitripFlightFlightchangeOrderQueryAPIResponse) {
	v.Reset()
	poolAlitripFlightFlightchangeOrderQueryAPIResponse.Put(v)
}

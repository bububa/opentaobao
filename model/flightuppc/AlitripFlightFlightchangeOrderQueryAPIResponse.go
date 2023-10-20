package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightflightchangeorderqueryAPIResponse 订单维度航变查询 API返回值
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
type AlitripflightflightchangeorderqueryAPIResponse struct {
	model.CommonResponse
	AlitripflightflightchangeorderqueryAPIResponseModel
}

// AlitripflightflightchangeorderqueryAPIResponseModel is 订单维度航变查询 成功返回结果
type AlitripflightflightchangeorderqueryAPIResponseModel struct {
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

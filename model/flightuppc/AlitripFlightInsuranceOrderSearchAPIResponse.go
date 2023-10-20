package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderSearchAPIResponse 查询保险订单详情 API返回值
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
type AlitripFlightInsuranceOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripFlightInsuranceOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightInsuranceOrderSearchAPIResponseModel).Reset()
}

// AlitripFlightInsuranceOrderSearchAPIResponseModel is 查询保险订单详情 成功返回结果
type AlitripFlightInsuranceOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 保险订单
	InsOrders []InsOrderOpenDto `json:"ins_orders,omitempty" xml:"ins_orders>ins_order_open_dto,omitempty"`
	// 错误信息，调用成功是该字段为空
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InsOrders = m.InsOrders[:0]
	m.ErrMsgForClient = ""
	m.IsSuccess = false
}

var poolAlitripFlightInsuranceOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightInsuranceOrderSearchAPIResponse)
	},
}

// GetAlitripFlightInsuranceOrderSearchAPIResponse 从 sync.Pool 获取 AlitripFlightInsuranceOrderSearchAPIResponse
func GetAlitripFlightInsuranceOrderSearchAPIResponse() *AlitripFlightInsuranceOrderSearchAPIResponse {
	return poolAlitripFlightInsuranceOrderSearchAPIResponse.Get().(*AlitripFlightInsuranceOrderSearchAPIResponse)
}

// ReleaseAlitripFlightInsuranceOrderSearchAPIResponse 将 AlitripFlightInsuranceOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightInsuranceOrderSearchAPIResponse(v *AlitripFlightInsuranceOrderSearchAPIResponse) {
	v.Reset()
	poolAlitripFlightInsuranceOrderSearchAPIResponse.Put(v)
}

package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderModifyAPIResponse 保险订单批改申请 API返回值
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
type AlitripFlightInsuranceOrderModifyAPIResponse struct {
	model.CommonResponse
	AlitripFlightInsuranceOrderModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceOrderModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightInsuranceOrderModifyAPIResponseModel).Reset()
}

// AlitripFlightInsuranceOrderModifyAPIResponseModel is 保险订单批改申请 成功返回结果
type AlitripFlightInsuranceOrderModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_order_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息（调用成功时该字段为空）
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceOrderModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrMsgForClient = ""
	m.IsSuccess = false
}

var poolAlitripFlightInsuranceOrderModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightInsuranceOrderModifyAPIResponse)
	},
}

// GetAlitripFlightInsuranceOrderModifyAPIResponse 从 sync.Pool 获取 AlitripFlightInsuranceOrderModifyAPIResponse
func GetAlitripFlightInsuranceOrderModifyAPIResponse() *AlitripFlightInsuranceOrderModifyAPIResponse {
	return poolAlitripFlightInsuranceOrderModifyAPIResponse.Get().(*AlitripFlightInsuranceOrderModifyAPIResponse)
}

// ReleaseAlitripFlightInsuranceOrderModifyAPIResponse 将 AlitripFlightInsuranceOrderModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightInsuranceOrderModifyAPIResponse(v *AlitripFlightInsuranceOrderModifyAPIResponse) {
	v.Reset()
	poolAlitripFlightInsuranceOrderModifyAPIResponse.Put(v)
}

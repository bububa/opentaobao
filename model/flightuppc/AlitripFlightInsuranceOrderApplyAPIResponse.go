package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderApplyAPIResponse 多险种批量投保 API返回值
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
type AlitripFlightInsuranceOrderApplyAPIResponse struct {
	model.CommonResponse
	AlitripFlightInsuranceOrderApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceOrderApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightInsuranceOrderApplyAPIResponseModel).Reset()
}

// AlitripFlightInsuranceOrderApplyAPIResponseModel is 多险种批量投保 成功返回结果
type AlitripFlightInsuranceOrderApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_order_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	InsProductResultList []InsProductResult `json:"ins_product_result_list,omitempty" xml:"ins_product_result_list>ins_product_result,omitempty"`
	// 错误信息，调用成功时该字段为空
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceOrderApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InsProductResultList = m.InsProductResultList[:0]
	m.ErrMsgForClient = ""
	m.OutOrderId = 0
	m.IsSuccess = false
}

var poolAlitripFlightInsuranceOrderApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightInsuranceOrderApplyAPIResponse)
	},
}

// GetAlitripFlightInsuranceOrderApplyAPIResponse 从 sync.Pool 获取 AlitripFlightInsuranceOrderApplyAPIResponse
func GetAlitripFlightInsuranceOrderApplyAPIResponse() *AlitripFlightInsuranceOrderApplyAPIResponse {
	return poolAlitripFlightInsuranceOrderApplyAPIResponse.Get().(*AlitripFlightInsuranceOrderApplyAPIResponse)
}

// ReleaseAlitripFlightInsuranceOrderApplyAPIResponse 将 AlitripFlightInsuranceOrderApplyAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightInsuranceOrderApplyAPIResponse(v *AlitripFlightInsuranceOrderApplyAPIResponse) {
	v.Reset()
	poolAlitripFlightInsuranceOrderApplyAPIResponse.Put(v)
}

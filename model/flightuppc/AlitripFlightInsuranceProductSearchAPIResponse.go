package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceProductSearchAPIResponse 搜索保险产品 API返回值
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
type AlitripFlightInsuranceProductSearchAPIResponse struct {
	model.CommonResponse
	AlitripFlightInsuranceProductSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceProductSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightInsuranceProductSearchAPIResponseModel).Reset()
}

// AlitripFlightInsuranceProductSearchAPIResponseModel is 搜索保险产品 成功返回结果
type AlitripFlightInsuranceProductSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_product_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 保险产品列表
	InsProducts []InsProductDto `json:"ins_products,omitempty" xml:"ins_products>ins_product_dto,omitempty"`
	// 错误信息（调用成功时该字段为空）
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightInsuranceProductSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InsProducts = m.InsProducts[:0]
	m.ErrMsgForClient = ""
	m.IsSuccess = false
}

var poolAlitripFlightInsuranceProductSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightInsuranceProductSearchAPIResponse)
	},
}

// GetAlitripFlightInsuranceProductSearchAPIResponse 从 sync.Pool 获取 AlitripFlightInsuranceProductSearchAPIResponse
func GetAlitripFlightInsuranceProductSearchAPIResponse() *AlitripFlightInsuranceProductSearchAPIResponse {
	return poolAlitripFlightInsuranceProductSearchAPIResponse.Get().(*AlitripFlightInsuranceProductSearchAPIResponse)
}

// ReleaseAlitripFlightInsuranceProductSearchAPIResponse 将 AlitripFlightInsuranceProductSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightInsuranceProductSearchAPIResponse(v *AlitripFlightInsuranceProductSearchAPIResponse) {
	v.Reset()
	poolAlitripFlightInsuranceProductSearchAPIResponse.Put(v)
}

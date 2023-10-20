package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipaySignAPIResponse 支付宝小程序验签 API返回值
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
type AlitripFlightExternalAlipaySignAPIResponse struct {
	model.CommonResponse
	AlitripFlightExternalAlipaySignAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipaySignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightExternalAlipaySignAPIResponseModel).Reset()
}

// AlitripFlightExternalAlipaySignAPIResponseModel is 支付宝小程序验签 成功返回结果
type AlitripFlightExternalAlipaySignAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_external_alipay_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 签名结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 请求失败描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否请求成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipaySignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolAlitripFlightExternalAlipaySignAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightExternalAlipaySignAPIResponse)
	},
}

// GetAlitripFlightExternalAlipaySignAPIResponse 从 sync.Pool 获取 AlitripFlightExternalAlipaySignAPIResponse
func GetAlitripFlightExternalAlipaySignAPIResponse() *AlitripFlightExternalAlipaySignAPIResponse {
	return poolAlitripFlightExternalAlipaySignAPIResponse.Get().(*AlitripFlightExternalAlipaySignAPIResponse)
}

// ReleaseAlitripFlightExternalAlipaySignAPIResponse 将 AlitripFlightExternalAlipaySignAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightExternalAlipaySignAPIResponse(v *AlitripFlightExternalAlipaySignAPIResponse) {
	v.Reset()
	poolAlitripFlightExternalAlipaySignAPIResponse.Put(v)
}

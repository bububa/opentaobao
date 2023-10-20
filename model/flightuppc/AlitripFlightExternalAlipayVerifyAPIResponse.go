package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipayVerifyAPIResponse 支付宝小程序验签 API返回值
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
type AlitripFlightExternalAlipayVerifyAPIResponse struct {
	model.CommonResponse
	AlitripFlightExternalAlipayVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipayVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightExternalAlipayVerifyAPIResponseModel).Reset()
}

// AlitripFlightExternalAlipayVerifyAPIResponseModel is 支付宝小程序验签 成功返回结果
type AlitripFlightExternalAlipayVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_external_alipay_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求失败描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否验签成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否请求成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipayVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = false
	m.IsSuccess = false
}

var poolAlitripFlightExternalAlipayVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightExternalAlipayVerifyAPIResponse)
	},
}

// GetAlitripFlightExternalAlipayVerifyAPIResponse 从 sync.Pool 获取 AlitripFlightExternalAlipayVerifyAPIResponse
func GetAlitripFlightExternalAlipayVerifyAPIResponse() *AlitripFlightExternalAlipayVerifyAPIResponse {
	return poolAlitripFlightExternalAlipayVerifyAPIResponse.Get().(*AlitripFlightExternalAlipayVerifyAPIResponse)
}

// ReleaseAlitripFlightExternalAlipayVerifyAPIResponse 将 AlitripFlightExternalAlipayVerifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightExternalAlipayVerifyAPIResponse(v *AlitripFlightExternalAlipayVerifyAPIResponse) {
	v.Reset()
	poolAlitripFlightExternalAlipayVerifyAPIResponse.Put(v)
}

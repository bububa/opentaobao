package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipayEncryptAPIResponse 支付宝小程序明文加密 API返回值
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
type AlitripFlightExternalAlipayEncryptAPIResponse struct {
	model.CommonResponse
	AlitripFlightExternalAlipayEncryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipayEncryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightExternalAlipayEncryptAPIResponseModel).Reset()
}

// AlitripFlightExternalAlipayEncryptAPIResponseModel is 支付宝小程序明文加密 成功返回结果
type AlitripFlightExternalAlipayEncryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_external_alipay_encrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 明文加密后的密文
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 请求失败描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否请求成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipayEncryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolAlitripFlightExternalAlipayEncryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightExternalAlipayEncryptAPIResponse)
	},
}

// GetAlitripFlightExternalAlipayEncryptAPIResponse 从 sync.Pool 获取 AlitripFlightExternalAlipayEncryptAPIResponse
func GetAlitripFlightExternalAlipayEncryptAPIResponse() *AlitripFlightExternalAlipayEncryptAPIResponse {
	return poolAlitripFlightExternalAlipayEncryptAPIResponse.Get().(*AlitripFlightExternalAlipayEncryptAPIResponse)
}

// ReleaseAlitripFlightExternalAlipayEncryptAPIResponse 将 AlitripFlightExternalAlipayEncryptAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightExternalAlipayEncryptAPIResponse(v *AlitripFlightExternalAlipayEncryptAPIResponse) {
	v.Reset()
	poolAlitripFlightExternalAlipayEncryptAPIResponse.Put(v)
}

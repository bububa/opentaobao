package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipayDecryptAPIResponse 支付宝小程序密文解密 API返回值
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
type AlitripFlightExternalAlipayDecryptAPIResponse struct {
	model.CommonResponse
	AlitripFlightExternalAlipayDecryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipayDecryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightExternalAlipayDecryptAPIResponseModel).Reset()
}

// AlitripFlightExternalAlipayDecryptAPIResponseModel is 支付宝小程序密文解密 成功返回结果
type AlitripFlightExternalAlipayDecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_external_alipay_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 解密后的明文
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 请求失败描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否请求成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipayDecryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolAlitripFlightExternalAlipayDecryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightExternalAlipayDecryptAPIResponse)
	},
}

// GetAlitripFlightExternalAlipayDecryptAPIResponse 从 sync.Pool 获取 AlitripFlightExternalAlipayDecryptAPIResponse
func GetAlitripFlightExternalAlipayDecryptAPIResponse() *AlitripFlightExternalAlipayDecryptAPIResponse {
	return poolAlitripFlightExternalAlipayDecryptAPIResponse.Get().(*AlitripFlightExternalAlipayDecryptAPIResponse)
}

// ReleaseAlitripFlightExternalAlipayDecryptAPIResponse 将 AlitripFlightExternalAlipayDecryptAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightExternalAlipayDecryptAPIResponse(v *AlitripFlightExternalAlipayDecryptAPIResponse) {
	v.Reset()
	poolAlitripFlightExternalAlipayDecryptAPIResponse.Put(v)
}

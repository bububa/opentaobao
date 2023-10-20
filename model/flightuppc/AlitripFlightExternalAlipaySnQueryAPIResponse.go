package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipaySnQueryAPIResponse 支付宝小程序查询证书序列号 API返回值
// alitrip.flight.external.alipay.sn.query
//
// 支付宝小程序查询证书序列号
type AlitripFlightExternalAlipaySnQueryAPIResponse struct {
	model.CommonResponse
	AlitripFlightExternalAlipaySnQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipaySnQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFlightExternalAlipaySnQueryAPIResponseModel).Reset()
}

// AlitripFlightExternalAlipaySnQueryAPIResponseModel is 支付宝小程序查询证书序列号 成功返回结果
type AlitripFlightExternalAlipaySnQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_external_alipay_sn_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求失败描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 证书编码结构体
	Result *AlipayCertSnDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否请求成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFlightExternalAlipaySnQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = nil
	m.IsSuccess = false
}

var poolAlitripFlightExternalAlipaySnQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFlightExternalAlipaySnQueryAPIResponse)
	},
}

// GetAlitripFlightExternalAlipaySnQueryAPIResponse 从 sync.Pool 获取 AlitripFlightExternalAlipaySnQueryAPIResponse
func GetAlitripFlightExternalAlipaySnQueryAPIResponse() *AlitripFlightExternalAlipaySnQueryAPIResponse {
	return poolAlitripFlightExternalAlipaySnQueryAPIResponse.Get().(*AlitripFlightExternalAlipaySnQueryAPIResponse)
}

// ReleaseAlitripFlightExternalAlipaySnQueryAPIResponse 将 AlitripFlightExternalAlipaySnQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripFlightExternalAlipaySnQueryAPIResponse(v *AlitripFlightExternalAlipaySnQueryAPIResponse) {
	v.Reset()
	poolAlitripFlightExternalAlipaySnQueryAPIResponse.Put(v)
}

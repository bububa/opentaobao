package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipayencryptAPIResponse 支付宝小程序明文加密 API返回值
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
type AlitripflightexternalalipayencryptAPIResponse struct {
	model.CommonResponse
	AlitripflightexternalalipayencryptAPIResponseModel
}

// AlitripflightexternalalipayencryptAPIResponseModel is 支付宝小程序明文加密 成功返回结果
type AlitripflightexternalalipayencryptAPIResponseModel struct {
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

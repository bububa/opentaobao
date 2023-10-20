package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgeniethirdtelecomautoauthAPIResponse 电信iot自动授权 API返回值
// alibaba.ailabs.tmallgenie.third.telecom.autoauth
//
// 电信iot自动授权
type AlibabaailabstmallgeniethirdtelecomautoauthAPIResponse struct {
	model.CommonResponse
	AlibabaailabstmallgeniethirdtelecomautoauthAPIResponseModel
}

// AlibabaailabstmallgeniethirdtelecomautoauthAPIResponseModel is 电信iot自动授权 成功返回结果
type AlibabaailabstmallgeniethirdtelecomautoauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_third_telecom_autoauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

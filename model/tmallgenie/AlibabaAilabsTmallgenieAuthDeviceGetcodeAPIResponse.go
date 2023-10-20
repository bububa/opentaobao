package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicegetcodeAPIResponse 获取authcode API返回值
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
type AlibabaailabstmallgenieauthdevicegetcodeAPIResponse struct {
	model.CommonResponse
	AlibabaailabstmallgenieauthdevicegetcodeAPIResponseModel
}

// AlibabaailabstmallgenieauthdevicegetcodeAPIResponseModel is 获取authcode 成功返回结果
type AlibabaailabstmallgenieauthdevicegetcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_getcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// authcode
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

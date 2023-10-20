package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthgetcodeAPIResponse 获取token API返回值
// alibaba.ailabs.tmallgenie.auth.getcode
//
// 获取天猫精灵authCode
type AlibabaailabstmallgenieauthgetcodeAPIResponse struct {
	model.CommonResponse
	AlibabaailabstmallgenieauthgetcodeAPIResponseModel
}

// AlibabaailabstmallgenieauthgetcodeAPIResponseModel is 获取token 成功返回结果
type AlibabaailabstmallgenieauthgetcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_getcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权码 json 串，包含授权码和二维码URL
	AuthCode string `json:"auth_code,omitempty" xml:"auth_code,omitempty"`
}

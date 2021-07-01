package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse
sso_token刷新 API返回值
alibaba.alisports.passport.account.ssotokenrefresh

sso_token刷新 */
type AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel
}

// AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel is sso_token刷新 成功返回结果
type AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_ssotokenrefresh_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_data
	AlispData *AlispData `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// 状态码 200表示操作成功
	AlispCode string `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 提示信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
}

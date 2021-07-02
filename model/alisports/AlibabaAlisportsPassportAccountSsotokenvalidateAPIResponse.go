package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse sso_token验证 API返回值
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
type AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel
}

// AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel is sso_token验证 成功返回结果
type AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_ssotokenvalidate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码 200表示操作成功
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 提示信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 阿里体育用户ID
	Aliuid string `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// 第三方用户ID
	Appuid string `json:"appuid,omitempty" xml:"appuid,omitempty"`
}

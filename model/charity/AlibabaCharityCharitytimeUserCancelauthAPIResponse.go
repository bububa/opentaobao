package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimeusercancelauthAPIResponse 取消用户授权 API返回值
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
type AlibabacharitycharitytimeusercancelauthAPIResponse struct {
	model.CommonResponse
	AlibabacharitycharitytimeusercancelauthAPIResponseModel
}

// AlibabacharitycharitytimeusercancelauthAPIResponseModel is 取消用户授权 成功返回结果
type AlibabacharitycharitytimeusercancelauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_user_cancelauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}

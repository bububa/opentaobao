package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserExternalAuthCancelAPIResponse 取消外部授权 API返回值
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
type AlibabaCharityUserExternalAuthCancelAPIResponse struct {
	model.CommonResponse
	AlibabaCharityUserExternalAuthCancelAPIResponseModel
}

// AlibabaCharityUserExternalAuthCancelAPIResponseModel is 取消外部授权 成功返回结果
type AlibabaCharityUserExternalAuthCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_user_external_auth_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}

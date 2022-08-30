package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserExternalAuthAPIResponse 外部用户授权 API返回值
// alibaba.charity.user.external.auth
//
// 外部用户授权
type AlibabaCharityUserExternalAuthAPIResponse struct {
	model.CommonResponse
	AlibabaCharityUserExternalAuthAPIResponseModel
}

// AlibabaCharityUserExternalAuthAPIResponseModel is 外部用户授权 成功返回结果
type AlibabaCharityUserExternalAuthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_user_external_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}

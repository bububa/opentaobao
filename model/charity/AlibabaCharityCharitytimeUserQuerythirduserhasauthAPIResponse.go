package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse 查询是否绑定3小时账号 API返回值
// alibaba.charity.charitytime.user.querythirduserhasauth
//
// 查询是否绑定3小时账号
type AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel
}

// AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel is 查询是否绑定3小时账号 成功返回结果
type AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_user_querythirduserhasauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}

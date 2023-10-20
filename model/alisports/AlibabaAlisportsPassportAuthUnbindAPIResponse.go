package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportauthunbindAPIResponse 三方关系解绑接口 API返回值
// alibaba.alisports.passport.auth.unbind
//
// 解除阿里体育openId和三方合作方的openId的绑定关系
type AlibabaalisportspassportauthunbindAPIResponse struct {
	model.CommonResponse
	AlibabaalisportspassportauthunbindAPIResponseModel
}

// AlibabaalisportspassportauthunbindAPIResponseModel is 三方关系解绑接口 成功返回结果
type AlibabaalisportspassportauthunbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_auth_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 体育返回实体对象
	Result *AlispResult `json:"result,omitempty" xml:"result,omitempty"`
}

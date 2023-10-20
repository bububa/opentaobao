package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportauthbindAPIResponse 授权绑定关系接口 API返回值
// alibaba.alisports.passport.auth.bind
//
// 授权回调绑定关系接口，建立阿里体育openId和三方openId的绑定关系
type AlibabaalisportspassportauthbindAPIResponse struct {
	model.CommonResponse
	AlibabaalisportspassportauthbindAPIResponseModel
}

// AlibabaalisportspassportauthbindAPIResponseModel is 授权绑定关系接口 成功返回结果
type AlibabaalisportspassportauthbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_auth_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果实体
	Result *AlispResult `json:"result,omitempty" xml:"result,omitempty"`
}

package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportauthaccountinfoAPIResponse 授权账号信息 API返回值
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
type AlibabaalisportspassportauthaccountinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalisportspassportauthaccountinfoAPIResponseModel
}

// AlibabaalisportspassportauthaccountinfoAPIResponseModel is 授权账号信息 成功返回结果
type AlibabaalisportspassportauthaccountinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_auth_accountinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlispResult `json:"result,omitempty" xml:"result,omitempty"`
}

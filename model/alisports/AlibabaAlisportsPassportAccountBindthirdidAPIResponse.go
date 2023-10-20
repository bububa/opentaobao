package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccountbindthirdidAPIResponse 阿里体育三方ID绑定接口 API返回值
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
type AlibabaalisportspassportaccountbindthirdidAPIResponse struct {
	model.CommonResponse
	AlibabaalisportspassportaccountbindthirdidAPIResponseModel
}

// AlibabaalisportspassportaccountbindthirdidAPIResponseModel is 阿里体育三方ID绑定接口 成功返回结果
type AlibabaalisportspassportaccountbindthirdidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_bindthirdid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口响应码
	AlispCode string `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 描述
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
}

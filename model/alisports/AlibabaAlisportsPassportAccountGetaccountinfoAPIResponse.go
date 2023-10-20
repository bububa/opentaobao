package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccountgetaccountinfoAPIResponse 获取会员信息 API返回值
// alibaba.alisports.passport.account.getaccountinfo
//
// 获取阿里体育会员信息
type AlibabaalisportspassportaccountgetaccountinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalisportspassportaccountgetaccountinfoAPIResponseModel
}

// AlibabaalisportspassportaccountgetaccountinfoAPIResponseModel is 获取会员信息 成功返回结果
type AlibabaalisportspassportaccountgetaccountinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_getaccountinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提示信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 返回值
	AlispData string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// 状态码 200表示操作成功
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

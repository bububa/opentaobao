package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountTokenvalidateAPIResponse 阿里体育会员系统帐号登录注册token验证接口 API返回值
// alibaba.alisports.passport.account.tokenvalidate
//
// 阿里体育会员系统帐号登录注册token验证接口
type AlibabaAlisportsPassportAccountTokenvalidateAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountTokenvalidateAPIResponseModel
}

// AlibabaAlisportsPassportAccountTokenvalidateAPIResponseModel is 阿里体育会员系统帐号登录注册token验证接口 成功返回结果
type AlibabaAlisportsPassportAccountTokenvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_tokenvalidate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回状态信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 返回数据结果
	AlispData string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// 返回状态码
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

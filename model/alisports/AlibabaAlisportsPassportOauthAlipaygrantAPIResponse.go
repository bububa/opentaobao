package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportOauthAlipaygrantAPIResponse 阿里体育会员系统-支付宝授权接口 API返回值
// alibaba.alisports.passport.oauth.alipaygrant
//
// 开放给乐心运动使用的支付宝授权接口
type AlibabaAlisportsPassportOauthAlipaygrantAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel
}

// AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel is 阿里体育会员系统-支付宝授权接口 成功返回结果
type AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_oauth_alipaygrant_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回状态码，200标识成功
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 返回信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 返回数据
	AlispData *AlispData `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
}

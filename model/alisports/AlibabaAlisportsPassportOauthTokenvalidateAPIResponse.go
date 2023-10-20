package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportoauthtokenvalidateAPIResponse 阿里体育会员系统--获取登录信息接口 API返回值
// alibaba.alisports.passport.oauth.tokenvalidate
//
// 阿里体育会员系统--获取登录信息接口
type AlibabaalisportspassportoauthtokenvalidateAPIResponse struct {
	model.CommonResponse
	AlibabaalisportspassportoauthtokenvalidateAPIResponseModel
}

// AlibabaalisportspassportoauthtokenvalidateAPIResponseModel is 阿里体育会员系统--获取登录信息接口 成功返回结果
type AlibabaalisportspassportoauthtokenvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_oauth_tokenvalidate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_data
	AlispData string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

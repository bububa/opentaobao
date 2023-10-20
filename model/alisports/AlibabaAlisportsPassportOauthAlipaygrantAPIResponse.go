package alisports

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlisportsPassportOauthAlipaygrantAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel is 阿里体育会员系统-支付宝授权接口 成功返回结果
type AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_oauth_alipaygrant_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 返回状态码，200标识成功
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 返回数据
	AlispData *AlispData `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportOauthAlipaygrantAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispCode = 0
	m.AlispData = nil
}

var poolAlibabaAlisportsPassportOauthAlipaygrantAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportOauthAlipaygrantAPIResponse)
	},
}

// GetAlibabaAlisportsPassportOauthAlipaygrantAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportOauthAlipaygrantAPIResponse
func GetAlibabaAlisportsPassportOauthAlipaygrantAPIResponse() *AlibabaAlisportsPassportOauthAlipaygrantAPIResponse {
	return poolAlibabaAlisportsPassportOauthAlipaygrantAPIResponse.Get().(*AlibabaAlisportsPassportOauthAlipaygrantAPIResponse)
}

// ReleaseAlibabaAlisportsPassportOauthAlipaygrantAPIResponse 将 AlibabaAlisportsPassportOauthAlipaygrantAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportOauthAlipaygrantAPIResponse(v *AlibabaAlisportsPassportOauthAlipaygrantAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportOauthAlipaygrantAPIResponse.Put(v)
}

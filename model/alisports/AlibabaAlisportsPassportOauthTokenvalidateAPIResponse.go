package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportOauthTokenvalidateAPIResponse 阿里体育会员系统--获取登录信息接口 API返回值
// alibaba.alisports.passport.oauth.tokenvalidate
//
// 阿里体育会员系统--获取登录信息接口
type AlibabaAlisportsPassportOauthTokenvalidateAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportOauthTokenvalidateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportOauthTokenvalidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportOauthTokenvalidateAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportOauthTokenvalidateAPIResponseModel is 阿里体育会员系统--获取登录信息接口 成功返回结果
type AlibabaAlisportsPassportOauthTokenvalidateAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlisportsPassportOauthTokenvalidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispData = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsPassportOauthTokenvalidateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportOauthTokenvalidateAPIResponse)
	},
}

// GetAlibabaAlisportsPassportOauthTokenvalidateAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportOauthTokenvalidateAPIResponse
func GetAlibabaAlisportsPassportOauthTokenvalidateAPIResponse() *AlibabaAlisportsPassportOauthTokenvalidateAPIResponse {
	return poolAlibabaAlisportsPassportOauthTokenvalidateAPIResponse.Get().(*AlibabaAlisportsPassportOauthTokenvalidateAPIResponse)
}

// ReleaseAlibabaAlisportsPassportOauthTokenvalidateAPIResponse 将 AlibabaAlisportsPassportOauthTokenvalidateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportOauthTokenvalidateAPIResponse(v *AlibabaAlisportsPassportOauthTokenvalidateAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportOauthTokenvalidateAPIResponse.Put(v)
}

package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAuthUnbindAPIResponse 三方关系解绑接口 API返回值
// alibaba.alisports.passport.auth.unbind
//
// 解除阿里体育openId和三方合作方的openId的绑定关系
type AlibabaAlisportsPassportAuthUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAuthUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAuthUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAuthUnbindAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAuthUnbindAPIResponseModel is 三方关系解绑接口 成功返回结果
type AlibabaAlisportsPassportAuthUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_auth_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 体育返回实体对象
	Result *AlispResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAuthUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlisportsPassportAuthUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAuthUnbindAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAuthUnbindAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAuthUnbindAPIResponse
func GetAlibabaAlisportsPassportAuthUnbindAPIResponse() *AlibabaAlisportsPassportAuthUnbindAPIResponse {
	return poolAlibabaAlisportsPassportAuthUnbindAPIResponse.Get().(*AlibabaAlisportsPassportAuthUnbindAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAuthUnbindAPIResponse 将 AlibabaAlisportsPassportAuthUnbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAuthUnbindAPIResponse(v *AlibabaAlisportsPassportAuthUnbindAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAuthUnbindAPIResponse.Put(v)
}

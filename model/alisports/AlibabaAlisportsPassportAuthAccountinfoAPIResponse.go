package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAuthAccountinfoAPIResponse 授权账号信息 API返回值
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
type AlibabaAlisportsPassportAuthAccountinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAuthAccountinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAuthAccountinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAuthAccountinfoAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAuthAccountinfoAPIResponseModel is 授权账号信息 成功返回结果
type AlibabaAlisportsPassportAuthAccountinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_auth_accountinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlispResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAuthAccountinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlisportsPassportAuthAccountinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAuthAccountinfoAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAuthAccountinfoAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAuthAccountinfoAPIResponse
func GetAlibabaAlisportsPassportAuthAccountinfoAPIResponse() *AlibabaAlisportsPassportAuthAccountinfoAPIResponse {
	return poolAlibabaAlisportsPassportAuthAccountinfoAPIResponse.Get().(*AlibabaAlisportsPassportAuthAccountinfoAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAuthAccountinfoAPIResponse 将 AlibabaAlisportsPassportAuthAccountinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAuthAccountinfoAPIResponse(v *AlibabaAlisportsPassportAuthAccountinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAuthAccountinfoAPIResponse.Put(v)
}

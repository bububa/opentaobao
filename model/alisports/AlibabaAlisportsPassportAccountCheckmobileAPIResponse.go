package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountCheckmobileAPIResponse 阿里体育会员系统--手机号验证接口 API返回值
// alibaba.alisports.passport.account.checkmobile
//
// 验证三方用户的手机号，获取对应的阿里体育会员id
type AlibabaAlisportsPassportAccountCheckmobileAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountCheckmobileAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel is 阿里体育会员系统--手机号验证接口 成功返回结果
type AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_checkmobile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// alisp_data
	AlispData *AlispData `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispCode = 0
	m.AlispData = nil
}

var poolAlibabaAlisportsPassportAccountCheckmobileAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAccountCheckmobileAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAccountCheckmobileAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAccountCheckmobileAPIResponse
func GetAlibabaAlisportsPassportAccountCheckmobileAPIResponse() *AlibabaAlisportsPassportAccountCheckmobileAPIResponse {
	return poolAlibabaAlisportsPassportAccountCheckmobileAPIResponse.Get().(*AlibabaAlisportsPassportAccountCheckmobileAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAccountCheckmobileAPIResponse 将 AlibabaAlisportsPassportAccountCheckmobileAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAccountCheckmobileAPIResponse(v *AlibabaAlisportsPassportAccountCheckmobileAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAccountCheckmobileAPIResponse.Put(v)
}

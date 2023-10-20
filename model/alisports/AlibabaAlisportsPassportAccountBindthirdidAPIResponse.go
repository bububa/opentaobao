package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountBindthirdidAPIResponse 阿里体育三方ID绑定接口 API返回值
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
type AlibabaAlisportsPassportAccountBindthirdidAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountBindthirdidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountBindthirdidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAccountBindthirdidAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAccountBindthirdidAPIResponseModel is 阿里体育三方ID绑定接口 成功返回结果
type AlibabaAlisportsPassportAccountBindthirdidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_bindthirdid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口响应码
	AlispCode string `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 描述
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountBindthirdidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispCode = ""
	m.AlispMsg = ""
}

var poolAlibabaAlisportsPassportAccountBindthirdidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAccountBindthirdidAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAccountBindthirdidAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAccountBindthirdidAPIResponse
func GetAlibabaAlisportsPassportAccountBindthirdidAPIResponse() *AlibabaAlisportsPassportAccountBindthirdidAPIResponse {
	return poolAlibabaAlisportsPassportAccountBindthirdidAPIResponse.Get().(*AlibabaAlisportsPassportAccountBindthirdidAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAccountBindthirdidAPIResponse 将 AlibabaAlisportsPassportAccountBindthirdidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAccountBindthirdidAPIResponse(v *AlibabaAlisportsPassportAccountBindthirdidAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAccountBindthirdidAPIResponse.Put(v)
}

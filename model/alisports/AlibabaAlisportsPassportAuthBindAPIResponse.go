package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAuthBindAPIResponse 授权绑定关系接口 API返回值
// alibaba.alisports.passport.auth.bind
//
// 授权回调绑定关系接口，建立阿里体育openId和三方openId的绑定关系
type AlibabaAlisportsPassportAuthBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAuthBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAuthBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAuthBindAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAuthBindAPIResponseModel is 授权绑定关系接口 成功返回结果
type AlibabaAlisportsPassportAuthBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_auth_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果实体
	Result *AlispResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAuthBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlisportsPassportAuthBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAuthBindAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAuthBindAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAuthBindAPIResponse
func GetAlibabaAlisportsPassportAuthBindAPIResponse() *AlibabaAlisportsPassportAuthBindAPIResponse {
	return poolAlibabaAlisportsPassportAuthBindAPIResponse.Get().(*AlibabaAlisportsPassportAuthBindAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAuthBindAPIResponse 将 AlibabaAlisportsPassportAuthBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAuthBindAPIResponse(v *AlibabaAlisportsPassportAuthBindAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAuthBindAPIResponse.Put(v)
}

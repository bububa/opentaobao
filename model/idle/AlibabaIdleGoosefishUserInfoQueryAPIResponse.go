package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleGoosefishUserInfoQueryAPIResponse 闲鱼三方容器用户信息获取 API返回值
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
type AlibabaIdleGoosefishUserInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleGoosefishUserInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleGoosefishUserInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleGoosefishUserInfoQueryAPIResponseModel).Reset()
}

// AlibabaIdleGoosefishUserInfoQueryAPIResponseModel is 闲鱼三方容器用户信息获取 成功返回结果
type AlibabaIdleGoosefishUserInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_goosefish_user_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ApiErrorCode string `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 错误说明
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 用户数据
	Data *IdleGooseFishUserInfoVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleGoosefishUserInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiErrorCode = ""
	m.ApiErrorMsg = ""
	m.Data = nil
	m.ApiSuccess = false
}

var poolAlibabaIdleGoosefishUserInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleGoosefishUserInfoQueryAPIResponse)
	},
}

// GetAlibabaIdleGoosefishUserInfoQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleGoosefishUserInfoQueryAPIResponse
func GetAlibabaIdleGoosefishUserInfoQueryAPIResponse() *AlibabaIdleGoosefishUserInfoQueryAPIResponse {
	return poolAlibabaIdleGoosefishUserInfoQueryAPIResponse.Get().(*AlibabaIdleGoosefishUserInfoQueryAPIResponse)
}

// ReleaseAlibabaIdleGoosefishUserInfoQueryAPIResponse 将 AlibabaIdleGoosefishUserInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleGoosefishUserInfoQueryAPIResponse(v *AlibabaIdleGoosefishUserInfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleGoosefishUserInfoQueryAPIResponse.Put(v)
}

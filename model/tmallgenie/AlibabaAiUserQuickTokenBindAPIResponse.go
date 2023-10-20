package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiUserQuickTokenBindAPIResponse 人工智能实验室精灵用户绑定第三方Token接口 API返回值
// alibaba.ai.user.quick.token.bind
//
// 人工智能实验室精灵用户绑定第三方Token接口
type AlibabaAiUserQuickTokenBindAPIResponse struct {
	model.CommonResponse
	AlibabaAiUserQuickTokenBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiUserQuickTokenBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiUserQuickTokenBindAPIResponseModel).Reset()
}

// AlibabaAiUserQuickTokenBindAPIResponseModel is 人工智能实验室精灵用户绑定第三方Token接口 成功返回结果
type AlibabaAiUserQuickTokenBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_user_quick_token_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAiUserQuickTokenBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.StatusCode = 0
}

var poolAlibabaAiUserQuickTokenBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiUserQuickTokenBindAPIResponse)
	},
}

// GetAlibabaAiUserQuickTokenBindAPIResponse 从 sync.Pool 获取 AlibabaAiUserQuickTokenBindAPIResponse
func GetAlibabaAiUserQuickTokenBindAPIResponse() *AlibabaAiUserQuickTokenBindAPIResponse {
	return poolAlibabaAiUserQuickTokenBindAPIResponse.Get().(*AlibabaAiUserQuickTokenBindAPIResponse)
}

// ReleaseAlibabaAiUserQuickTokenBindAPIResponse 将 AlibabaAiUserQuickTokenBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiUserQuickTokenBindAPIResponse(v *AlibabaAiUserQuickTokenBindAPIResponse) {
	v.Reset()
	poolAlibabaAiUserQuickTokenBindAPIResponse.Put(v)
}

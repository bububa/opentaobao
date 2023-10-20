package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiUserQuickBindAPIResponse 精灵用户绑定第三方账号信息 API返回值
// alibaba.ai.user.quick.bind
//
// 人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口
type AlibabaAiUserQuickBindAPIResponse struct {
	model.CommonResponse
	AlibabaAiUserQuickBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiUserQuickBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiUserQuickBindAPIResponseModel).Reset()
}

// AlibabaAiUserQuickBindAPIResponseModel is 精灵用户绑定第三方账号信息 成功返回结果
type AlibabaAiUserQuickBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_user_quick_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易结果描述（例如： 交易成功、交易失败）
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 扩展字段，无用
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 交易结果详细描述（例如：用户已经存在，交易失败）
	MessageDetail string `json:"message_detail,omitempty" xml:"message_detail,omitempty"`
	// 交易结果码（200：成功、其他：失败）
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAiUserQuickBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Result = ""
	m.MessageDetail = ""
	m.StatusCode = 0
}

var poolAlibabaAiUserQuickBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiUserQuickBindAPIResponse)
	},
}

// GetAlibabaAiUserQuickBindAPIResponse 从 sync.Pool 获取 AlibabaAiUserQuickBindAPIResponse
func GetAlibabaAiUserQuickBindAPIResponse() *AlibabaAiUserQuickBindAPIResponse {
	return poolAlibabaAiUserQuickBindAPIResponse.Get().(*AlibabaAiUserQuickBindAPIResponse)
}

// ReleaseAlibabaAiUserQuickBindAPIResponse 将 AlibabaAiUserQuickBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiUserQuickBindAPIResponse(v *AlibabaAiUserQuickBindAPIResponse) {
	v.Reset()
	poolAlibabaAiUserQuickBindAPIResponse.Put(v)
}

package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieSkillMessagePushAPIResponse 天猫精灵消息推送标准接口 API返回值
// alibaba.ailabs.aligenie.skill.message.push
//
// 用于AliGenie技能开发者在技能内主动向音响推送消息的标准服务接口，只有订阅过该消息的用户才能收到消息；
type AlibabaAilabsAligenieSkillMessagePushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieSkillMessagePushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieSkillMessagePushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieSkillMessagePushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieSkillMessagePushAPIResponseModel is 天猫精灵消息推送标准接口 成功返回结果
type AlibabaAilabsAligenieSkillMessagePushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_skill_message_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息推送错误码
	PushErrorCode string `json:"push_error_code,omitempty" xml:"push_error_code,omitempty"`
	// 消息推送错误信息
	PushErrorMsg string `json:"push_error_msg,omitempty" xml:"push_error_msg,omitempty"`
	// 消息推送是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 接口调用是否成功
	PushResult bool `json:"push_result,omitempty" xml:"push_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieSkillMessagePushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PushErrorCode = ""
	m.PushErrorMsg = ""
	m.Model = false
	m.PushResult = false
}

var poolAlibabaAilabsAligenieSkillMessagePushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieSkillMessagePushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieSkillMessagePushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieSkillMessagePushAPIResponse
func GetAlibabaAilabsAligenieSkillMessagePushAPIResponse() *AlibabaAilabsAligenieSkillMessagePushAPIResponse {
	return poolAlibabaAilabsAligenieSkillMessagePushAPIResponse.Get().(*AlibabaAilabsAligenieSkillMessagePushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieSkillMessagePushAPIResponse 将 AlibabaAilabsAligenieSkillMessagePushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieSkillMessagePushAPIResponse(v *AlibabaAilabsAligenieSkillMessagePushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieSkillMessagePushAPIResponse.Put(v)
}

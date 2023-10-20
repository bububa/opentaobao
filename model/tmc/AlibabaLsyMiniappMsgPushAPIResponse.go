package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyMiniappMsgPushAPIResponse 零售云小程序消息推送 API返回值
// alibaba.lsy.miniapp.msg.push
//
// 零售云小程序消息推送，推送消息至零售云（喵零等）
type AlibabaLsyMiniappMsgPushAPIResponse struct {
	model.CommonResponse
	AlibabaLsyMiniappMsgPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyMiniappMsgPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyMiniappMsgPushAPIResponseModel).Reset()
}

// AlibabaLsyMiniappMsgPushAPIResponseModel is 零售云小程序消息推送 成功返回结果
type AlibabaLsyMiniappMsgPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_miniapp_msg_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyMiniappMsgPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailMsg = ""
	m.FailCode = ""
	m.Succ = false
}

var poolAlibabaLsyMiniappMsgPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyMiniappMsgPushAPIResponse)
	},
}

// GetAlibabaLsyMiniappMsgPushAPIResponse 从 sync.Pool 获取 AlibabaLsyMiniappMsgPushAPIResponse
func GetAlibabaLsyMiniappMsgPushAPIResponse() *AlibabaLsyMiniappMsgPushAPIResponse {
	return poolAlibabaLsyMiniappMsgPushAPIResponse.Get().(*AlibabaLsyMiniappMsgPushAPIResponse)
}

// ReleaseAlibabaLsyMiniappMsgPushAPIResponse 将 AlibabaLsyMiniappMsgPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyMiniappMsgPushAPIResponse(v *AlibabaLsyMiniappMsgPushAPIResponse) {
	v.Reset()
	poolAlibabaLsyMiniappMsgPushAPIResponse.Put(v)
}

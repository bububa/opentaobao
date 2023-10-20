package msgamp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBcChatMessageSendAPIResponse 小程序资源授权-BC客服消息 API返回值
// taobao.bc.chat.message.send
//
// 小程序资源授权-消息订阅
type TaobaoBcChatMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoBcChatMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBcChatMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBcChatMessageSendAPIResponseModel).Reset()
}

// TaobaoBcChatMessageSendAPIResponseModel is 小程序资源授权-BC客服消息 成功返回结果
type TaobaoBcChatMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"bc_chat_message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBcChatMessageSendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBcChatMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBcChatMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBcChatMessageSendAPIResponse)
	},
}

// GetTaobaoBcChatMessageSendAPIResponse 从 sync.Pool 获取 TaobaoBcChatMessageSendAPIResponse
func GetTaobaoBcChatMessageSendAPIResponse() *TaobaoBcChatMessageSendAPIResponse {
	return poolTaobaoBcChatMessageSendAPIResponse.Get().(*TaobaoBcChatMessageSendAPIResponse)
}

// ReleaseTaobaoBcChatMessageSendAPIResponse 将 TaobaoBcChatMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBcChatMessageSendAPIResponse(v *TaobaoBcChatMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoBcChatMessageSendAPIResponse.Put(v)
}

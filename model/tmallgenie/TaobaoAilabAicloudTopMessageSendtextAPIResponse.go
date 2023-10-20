package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageSendtextAPIResponse 故事机发送文本留言 API返回值
// taobao.ailab.aicloud.top.message.sendtext
//
// 故事机文本留言
type TaobaoAilabAicloudTopMessageSendtextAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMessageSendtextAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageSendtextAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMessageSendtextAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMessageSendtextAPIResponseModel is 故事机发送文本留言 成功返回结果
type TaobaoAilabAicloudTopMessageSendtextAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_sendtext_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageSendtextAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMessageSendtextAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessageSendtextAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMessageSendtextAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageSendtextAPIResponse
func GetTaobaoAilabAicloudTopMessageSendtextAPIResponse() *TaobaoAilabAicloudTopMessageSendtextAPIResponse {
	return poolTaobaoAilabAicloudTopMessageSendtextAPIResponse.Get().(*TaobaoAilabAicloudTopMessageSendtextAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMessageSendtextAPIResponse 将 TaobaoAilabAicloudTopMessageSendtextAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageSendtextAPIResponse(v *TaobaoAilabAicloudTopMessageSendtextAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageSendtextAPIResponse.Put(v)
}

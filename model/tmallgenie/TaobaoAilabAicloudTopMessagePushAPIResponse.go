package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessagePushAPIResponse 天猫精灵消息中心广播推送消息接口 API返回值
// taobao.ailab.aicloud.top.message.push
//
// 天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
type TaobaoAilabAicloudTopMessagePushAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMessagePushAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessagePushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMessagePushAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMessagePushAPIResponseModel is 天猫精灵消息中心广播推送消息接口 成功返回结果
type TaobaoAilabAicloudTopMessagePushAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *OpsCommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessagePushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMessagePushAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessagePushAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMessagePushAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMessagePushAPIResponse
func GetTaobaoAilabAicloudTopMessagePushAPIResponse() *TaobaoAilabAicloudTopMessagePushAPIResponse {
	return poolTaobaoAilabAicloudTopMessagePushAPIResponse.Get().(*TaobaoAilabAicloudTopMessagePushAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMessagePushAPIResponse 将 TaobaoAilabAicloudTopMessagePushAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessagePushAPIResponse(v *TaobaoAilabAicloudTopMessagePushAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessagePushAPIResponse.Put(v)
}

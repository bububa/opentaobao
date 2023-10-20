package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagesendaudioAPIResponse 发送语音留言 API返回值
// taobao.ailab.aicloud.top.message.sendaudio
//
// 将语音的二进制byte[]通过TOP接口发送保存
type TaobaoailabaicloudtopmessagesendaudioAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmessagesendaudioAPIResponseModel
}

// TaobaoailabaicloudtopmessagesendaudioAPIResponseModel is 发送语音留言 成功返回结果
type TaobaoailabaicloudtopmessagesendaudioAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_sendaudio_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

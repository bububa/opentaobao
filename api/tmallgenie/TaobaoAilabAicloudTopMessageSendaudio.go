package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMessageSendaudio 发送语音留言
// taobao.ailab.aicloud.top.message.sendaudio
//
// 将语音的二进制byte[]通过TOP接口发送保存
func TaobaoAilabAicloudTopMessageSendaudio(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMessageSendaudioAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMessageSendaudioAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

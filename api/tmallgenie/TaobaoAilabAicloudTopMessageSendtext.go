package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMessageSendtext 故事机发送文本留言
// taobao.ailab.aicloud.top.message.sendtext
//
// 故事机文本留言
func TaobaoAilabAicloudTopMessageSendtext(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMessageSendtextAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMessageSendtextAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package msgamp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/msgamp"
)

// TaobaoBcChatMessageSend 小程序资源授权-BC客服消息
// taobao.bc.chat.message.send
//
// 小程序资源授权-消息订阅
func TaobaoBcChatMessageSend(clt *core.SDKClient, req *msgamp.TaobaoBcChatMessageSendAPIRequest, resp *msgamp.TaobaoBcChatMessageSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

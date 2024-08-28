package msgamp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/msgamp"
)

// TaobaoBcChatMessageSend 小程序资源授权-BC客服消息
// taobao.bc.chat.message.send
//
// 小程序资源授权-消息订阅
func TaobaoBcChatMessageSend(ctx context.Context, clt *core.SDKClient, req *msgamp.TaobaoBcChatMessageSendAPIRequest, resp *msgamp.TaobaoBcChatMessageSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

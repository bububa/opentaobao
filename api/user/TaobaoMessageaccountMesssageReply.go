package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMessageaccountMesssageReply 消息号下行回复接口
// taobao.messageaccount.messsage.reply
//
// 外部 isv 调用该进口来进行消息号消息的回复
func TaobaoMessageaccountMesssageReply(ctx context.Context, clt *core.SDKClient, req *user.TaobaoMessageaccountMesssageReplyAPIRequest, resp *user.TaobaoMessageaccountMesssageReplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

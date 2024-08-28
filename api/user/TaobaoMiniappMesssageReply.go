package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMiniappMesssageReply 轻店铺下行回复接口
// taobao.miniapp.messsage.reply
//
// 外部 isv 调用该进口来进行轻店铺消息的回复
func TaobaoMiniappMesssageReply(ctx context.Context, clt *core.SDKClient, req *user.TaobaoMiniappMesssageReplyAPIRequest, resp *user.TaobaoMiniappMesssageReplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

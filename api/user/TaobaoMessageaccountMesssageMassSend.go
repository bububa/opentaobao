package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMessageaccountMesssageMassSend 消息号开放-消息群发
// taobao.messageaccount.messsage.mass.send
//
// 外部 isv 调用该进口来进行消息号消息的群发
func TaobaoMessageaccountMesssageMassSend(ctx context.Context, clt *core.SDKClient, req *user.TaobaoMessageaccountMesssageMassSendAPIRequest, resp *user.TaobaoMessageaccountMesssageMassSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

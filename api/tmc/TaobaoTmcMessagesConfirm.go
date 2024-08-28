package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcMessagesConfirm 确认消费消息的状态
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
func TaobaoTmcMessagesConfirm(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcMessagesConfirmAPIRequest, resp *tmc.TaobaoTmcMessagesConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

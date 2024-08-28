package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimAppChatlogsGet openim应用聊天记录查询
// taobao.openim.app.chatlogs.get
//
// 查询openim应用的聊天记录
func TaobaoOpenimAppChatlogsGet(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimAppChatlogsGetAPIRequest, resp *openim.TaobaoOpenimAppChatlogsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

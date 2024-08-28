package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeInvite OPENIM群邀请加入
// taobao.openim.tribe.invite
//
// OPENIM群邀请加入接口
func TaobaoOpenimTribeInvite(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeInviteAPIRequest, resp *openim.TaobaoOpenimTribeInviteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

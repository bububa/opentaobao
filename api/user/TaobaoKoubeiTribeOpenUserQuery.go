package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoKoubeiTribeOpenUserQuery 获取用户openId
// taobao.koubei.tribe.open.user.query
//
// 口碑综合体通过手机号码获取加密后的用户openId
func TaobaoKoubeiTribeOpenUserQuery(ctx context.Context, clt *core.SDKClient, req *user.TaobaoKoubeiTribeOpenUserQueryAPIRequest, resp *user.TaobaoKoubeiTribeOpenUserQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package tbuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserAvatarGet 淘宝用户头像查询
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
func TaobaoUserAvatarGet(ctx context.Context, clt *core.SDKClient, req *tbuser.TaobaoUserAvatarGetAPIRequest, resp *tbuser.TaobaoUserAvatarGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

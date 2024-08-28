package game

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// TaobaoAppleNewuserActivateNotify 新用户激活通知接口
// taobao.apple.newuser.activate.notify
//
// 资和信主动通知激活结果
func TaobaoAppleNewuserActivateNotify(ctx context.Context, clt *core.SDKClient, req *game.TaobaoAppleNewuserActivateNotifyAPIRequest, resp *game.TaobaoAppleNewuserActivateNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

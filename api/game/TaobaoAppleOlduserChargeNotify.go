package game

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// TaobaoAppleOlduserChargeNotify 老用户激活并兑换通知接口
// taobao.apple.olduser.charge.notify
//
// 老用户激活并兑换通知接口
func TaobaoAppleOlduserChargeNotify(ctx context.Context, clt *core.SDKClient, req *game.TaobaoAppleOlduserChargeNotifyAPIRequest, resp *game.TaobaoAppleOlduserChargeNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

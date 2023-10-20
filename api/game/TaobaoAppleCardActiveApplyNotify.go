package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// TaobaoAppleCardActiveApplyNotify 苹果卡密申请激活回调接口
// taobao.apple.card.active.apply.notify
//
// 苹果卡密申请激活回调接口
func TaobaoAppleCardActiveApplyNotify(clt *core.SDKClient, req *game.TaobaoAppleCardActiveApplyNotifyAPIRequest, resp *game.TaobaoAppleCardActiveApplyNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

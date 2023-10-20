package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// TaobaoAppleCardActiveCancelNotify 苹果卡密取消激活回调接口
// taobao.apple.card.active.cancel.notify
//
// 苹果卡密取消激活回调接口
func TaobaoAppleCardActiveCancelNotify(clt *core.SDKClient, req *game.TaobaoAppleCardActiveCancelNotifyAPIRequest, resp *game.TaobaoAppleCardActiveCancelNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

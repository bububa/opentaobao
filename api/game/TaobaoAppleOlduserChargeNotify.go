package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// Taobaoappleolduserchargenotify 老用户激活并兑换通知接口
// taobao.apple.olduser.charge.notify
//
// 老用户激活并兑换通知接口
func Taobaoappleolduserchargenotify(clt *core.SDKClient, req *game.TaobaoappleolduserchargenotifyAPIRequest, session string) (*game.TaobaoappleolduserchargenotifyAPIResponse, error) {
	var resp game.TaobaoappleolduserchargenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

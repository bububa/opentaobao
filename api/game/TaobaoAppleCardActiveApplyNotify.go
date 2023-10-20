package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// Taobaoapplecardactiveapplynotify 苹果卡密申请激活回调接口
// taobao.apple.card.active.apply.notify
//
// 苹果卡密申请激活回调接口
func Taobaoapplecardactiveapplynotify(clt *core.SDKClient, req *game.TaobaoapplecardactiveapplynotifyAPIRequest, session string) (*game.TaobaoapplecardactiveapplynotifyAPIResponse, error) {
	var resp game.TaobaoapplecardactiveapplynotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

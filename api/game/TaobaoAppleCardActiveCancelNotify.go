package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// Taobaoapplecardactivecancelnotify 苹果卡密取消激活回调接口
// taobao.apple.card.active.cancel.notify
//
// 苹果卡密取消激活回调接口
func Taobaoapplecardactivecancelnotify(clt *core.SDKClient, req *game.TaobaoapplecardactivecancelnotifyAPIRequest, session string) (*game.TaobaoapplecardactivecancelnotifyAPIResponse, error) {
	var resp game.TaobaoapplecardactivecancelnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

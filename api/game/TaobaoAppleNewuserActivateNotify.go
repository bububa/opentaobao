package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// Taobaoapplenewuseractivatenotify 新用户激活通知接口
// taobao.apple.newuser.activate.notify
//
// 资和信主动通知激活结果
func Taobaoapplenewuseractivatenotify(clt *core.SDKClient, req *game.TaobaoapplenewuseractivatenotifyAPIRequest, session string) (*game.TaobaoapplenewuseractivatenotifyAPIResponse, error) {
	var resp game.TaobaoapplenewuseractivatenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

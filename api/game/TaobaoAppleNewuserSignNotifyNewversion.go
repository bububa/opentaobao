package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// Taobaoapplenewusersignnotifynewversion 新用户签约结果通知接口v2
// taobao.apple.newuser.sign.notify.newversion
//
// 资和信主动通知签约结果
func Taobaoapplenewusersignnotifynewversion(clt *core.SDKClient, req *game.TaobaoapplenewusersignnotifynewversionAPIRequest, session string) (*game.TaobaoapplenewusersignnotifynewversionAPIResponse, error) {
	var resp game.TaobaoapplenewusersignnotifynewversionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

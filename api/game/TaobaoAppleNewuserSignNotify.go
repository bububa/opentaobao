package game

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/game"
)

// TaobaoAppleNewuserSignNotify 新用户签约通知接口
// taobao.apple.newuser.sign.notify
//
// 用户付款成功后，资和信主动通知签约结果
func TaobaoAppleNewuserSignNotify(clt *core.SDKClient, req *game.TaobaoAppleNewuserSignNotifyAPIRequest, resp *game.TaobaoAppleNewuserSignNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

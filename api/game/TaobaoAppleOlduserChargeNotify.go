package game

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/game"
)

/* 
老用户激活并兑换通知接口 
taobao.apple.olduser.charge.notify

老用户激活并兑换通知接口
*/
func TaobaoAppleOlduserChargeNotify(clt *core.SDKClient, req *game.TaobaoAppleOlduserChargeNotifyRequest, session string) (*game.TaobaoAppleOlduserChargeNotifyAPIResponse, error) {
    var resp game.TaobaoAppleOlduserChargeNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

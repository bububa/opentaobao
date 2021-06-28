package game

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/game"
)

/* 
苹果卡密申请激活回调接口 
taobao.apple.card.active.apply.notify

苹果卡密申请激活回调接口
*/
func TaobaoAppleCardActiveApplyNotify(clt *core.SDKClient, req *game.TaobaoAppleCardActiveApplyNotifyRequest, session string) (*game.TaobaoAppleCardActiveApplyNotifyAPIResponse, error) {
    var resp game.TaobaoAppleCardActiveApplyNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

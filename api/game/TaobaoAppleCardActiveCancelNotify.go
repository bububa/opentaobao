package game

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/game"
)

/* 
苹果卡密取消激活回调接口 
taobao.apple.card.active.cancel.notify

苹果卡密取消激活回调接口
*/
func TaobaoAppleCardActiveCancelNotify(clt *core.SDKClient, req *game.TaobaoAppleCardActiveCancelNotifyRequest, session string) (*game.TaobaoAppleCardActiveCancelNotifyAPIResponse, error) {
    var resp game.TaobaoAppleCardActiveCancelNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

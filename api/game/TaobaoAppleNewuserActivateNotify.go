package game

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/game"
)

/* 
新用户激活通知接口 
taobao.apple.newuser.activate.notify

资和信主动通知激活结果
*/
func TaobaoAppleNewuserActivateNotify(clt *core.SDKClient, req *game.TaobaoAppleNewuserActivateNotifyAPIRequest, session string) (*game.TaobaoAppleNewuserActivateNotifyAPIResponse, error) {
    var resp game.TaobaoAppleNewuserActivateNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

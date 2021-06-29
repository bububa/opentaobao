package tmallgenie

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgenie"
)

/* 
天猫精灵酒店播放暂停 
taobao.tmallgenie.hotelplayerpause

酒店推送指令给天猫精灵停止播放音乐
*/
func TaobaoTmallgenieHotelplayerpause(clt *core.SDKClient, req *tmallgenie.TaobaoTmallgenieHotelplayerpauseRequest, session string) (*tmallgenie.TaobaoTmallgenieHotelplayerpauseAPIResponse, error) {
    var resp tmallgenie.TaobaoTmallgenieHotelplayerpauseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

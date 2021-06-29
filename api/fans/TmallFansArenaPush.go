package fans

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fans"
)

/* 
消息推送 
tmall.fans.arena.push

超级擂台消息推送
*/
func TmallFansArenaPush(clt *core.SDKClient, req *fans.TmallFansArenaPushRequest, session string) (*fans.TmallFansArenaPushAPIResponse, error) {
    var resp fans.TmallFansArenaPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

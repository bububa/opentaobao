package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
推广淘小铺isv 活动到微淘feed 
taobao.weitao.feed.synchronize.new

推广微淘互动应用活动到微淘
*/
func TaobaoWeitaoFeedSynchronizeNew(clt *core.SDKClient, req *interact.TaobaoWeitaoFeedSynchronizeNewAPIRequest, session string) (*interact.TaobaoWeitaoFeedSynchronizeNewAPIResponse, error) {
    var resp interact.TaobaoWeitaoFeedSynchronizeNewAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

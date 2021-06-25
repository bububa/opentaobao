package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
推广淘小铺isv 活动到微淘feed 
taobao.weitao.feed.synchronize

推广淘小铺isv 活动到微淘feed
*/
func TaobaoWeitaoFeedSynchronize(clt *core.SDKClient, req *interact.TaobaoWeitaoFeedSynchronizeRequest, session string) (*interact.TaobaoWeitaoFeedSynchronizeResponse, error) {
    var resp interact.TaobaoWeitaoFeedSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package mtopopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mtopopen"
)

/* 
是否关注 
taobao.weitao.feed.isrelation

判断用户是否关注对应的公共账号
*/
func TaobaoWeitaoFeedIsrelation(clt *core.SDKClient, req *mtopopen.TaobaoWeitaoFeedIsrelationRequest, session string) (*mtopopen.TaobaoWeitaoFeedIsrelationResponse, error) {
    var resp mtopopen.TaobaoWeitaoFeedIsrelationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

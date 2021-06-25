package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
（新）关键词更新相关接口 
taobao.simba.keyword.update

（新）关键词更新相关接口
*/
func TaobaoSimbaKeywordUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordUpdateRequest, session string) (*simba.TaobaoSimbaKeywordUpdateResponse, error) {
    var resp simba.TaobaoSimbaKeywordUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
删除一批关键词 
taobao.simba.keywords.delete

删除一批关键词
*/
func TaobaoSimbaKeywordsDelete(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsDeleteRequest, session string) (*simba.TaobaoSimbaKeywordsDeleteAPIResponse, error) {
    var resp simba.TaobaoSimbaKeywordsDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

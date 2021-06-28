package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
创建一批关键词 
taobao.simba.keywordsvon.add

创建一批关键词
*/
func TaobaoSimbaKeywordsvonAdd(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsvonAddRequest, session string) (*simba.TaobaoSimbaKeywordsvonAddAPIResponse, error) {
    var resp simba.TaobaoSimbaKeywordsvonAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

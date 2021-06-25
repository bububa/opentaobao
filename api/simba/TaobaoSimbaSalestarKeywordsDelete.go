package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
销量明星关键词删除 
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口
*/
func TaobaoSimbaSalestarKeywordsDelete(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarKeywordsDeleteRequest, session string) (*simba.TaobaoSimbaSalestarKeywordsDeleteResponse, error) {
    var resp simba.TaobaoSimbaSalestarKeywordsDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

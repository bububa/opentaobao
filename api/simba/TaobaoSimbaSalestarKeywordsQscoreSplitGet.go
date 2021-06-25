package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
(新)销量明星质量分相关接口 
taobao.simba.salestar.keywords.qscore.split.get

获取关键词新的质量分
*/
func TaobaoSimbaSalestarKeywordsQscoreSplitGet(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest, session string) (*simba.TaobaoSimbaSalestarKeywordsQscoreSplitGetResponse, error) {
    var resp simba.TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

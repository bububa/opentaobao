package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
新质量分服务 
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分
*/
func TaobaoSimbaKeywordsQscoreSplitGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsQscoreSplitGetRequest, session string) (*simba.TaobaoSimbaKeywordsQscoreSplitGetAPIResponse, error) {
    var resp simba.TaobaoSimbaKeywordsQscoreSplitGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
取得一个推广组的推荐关键词列表 
taobao.simba.keywords.recommend.get

取得一个推广组的推荐关键词列表
*/
func TaobaoSimbaKeywordsRecommendGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsRecommendGetRequest, session string) (*simba.TaobaoSimbaKeywordsRecommendGetAPIResponse, error) {
    var resp simba.TaobaoSimbaKeywordsRecommendGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

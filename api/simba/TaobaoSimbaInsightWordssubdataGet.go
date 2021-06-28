package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取关键词按流量细分的数据 
taobao.simba.insight.wordssubdata.get

获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1->PC站内，2->PC站外,4->无线站内 5->无线站外
*/
func TaobaoSimbaInsightWordssubdataGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightWordssubdataGetRequest, session string) (*simba.TaobaoSimbaInsightWordssubdataGetAPIResponse, error) {
    var resp simba.TaobaoSimbaInsightWordssubdataGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

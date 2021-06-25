package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取关键词的大盘数据 
taobao.simba.insight.wordsdata.get

获取关键词的详细数据
*/
func TaobaoSimbaInsightWordsdataGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightWordsdataGetRequest, session string) (*simba.TaobaoSimbaInsightWordsdataGetResponse, error) {
    var resp simba.TaobaoSimbaInsightWordsdataGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

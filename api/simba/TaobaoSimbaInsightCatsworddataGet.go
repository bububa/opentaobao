package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取类目下关键词的数据 
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据
*/
func TaobaoSimbaInsightCatsworddataGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightCatsworddataGetRequest, session string) (*simba.TaobaoSimbaInsightCatsworddataGetAPIResponse, error) {
    var resp simba.TaobaoSimbaInsightCatsworddataGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

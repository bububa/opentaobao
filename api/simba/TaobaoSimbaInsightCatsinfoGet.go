package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
类目信息获取 
taobao.simba.insight.catsinfo.get

获取类目信息，此接口既提供所有顶级类目的查询，又提供给定类目id自身信息和子类目信息的查询，所以可以根据此接口逐层获取所有的类目信息
*/
func TaobaoSimbaInsightCatsinfoGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightCatsinfoGetAPIRequest, session string) (*simba.TaobaoSimbaInsightCatsinfoGetAPIResponse, error) {
    var resp simba.TaobaoSimbaInsightCatsinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
家装业务查询物流公司api 
taobao.wlb.order.jz.query

家装业务查询物流公司api
*/
func TaobaoWlbOrderJzQuery(clt *core.SDKClient, req *logistic.TaobaoWlbOrderJzQueryRequest, session string) (*logistic.TaobaoWlbOrderJzQueryAPIResponse, error) {
    var resp logistic.TaobaoWlbOrderJzQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
服务商资金权益逆向扣回的查询接口 
tmall.servicecenter.tp.funds.recover.query

服务商资金权益逆向扣回的查询接口
*/
func TmallServicecenterTpFundsRecoverQuery(clt *core.SDKClient, req *servicecenter.TmallServicecenterTpFundsRecoverQueryRequest, session string) (*servicecenter.TmallServicecenterTpFundsRecoverQueryResponse, error) {
    var resp servicecenter.TmallServicecenterTpFundsRecoverQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

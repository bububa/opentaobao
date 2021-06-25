package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
服务商资金权益发放的查询接口 
tmall.servicecenter.tp.funds.send.query

服务商资金权益发放结果的查询接口
*/
func TmallServicecenterTpFundsSendQuery(clt *core.SDKClient, req *servicecenter.TmallServicecenterTpFundsSendQueryRequest, session string) (*servicecenter.TmallServicecenterTpFundsSendQueryResponse, error) {
    var resp servicecenter.TmallServicecenterTpFundsSendQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

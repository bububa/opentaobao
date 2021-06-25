package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
物流订单呼叫运力 
tmall.servicecenter.workcard.expressorder.consign

天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
*/
func TmallServicecenterWorkcardExpressorderConsign(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardExpressorderConsignRequest, session string) (*tmallservice.TmallServicecenterWorkcardExpressorderConsignResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardExpressorderConsignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

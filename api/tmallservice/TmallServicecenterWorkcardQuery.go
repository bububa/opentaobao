package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
工单查询接口 
tmall.servicecenter.workcard.query

工单查询接口
*/
func TmallServicecenterWorkcardQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQueryRequest, session string) (*tmallservice.TmallServicecenterWorkcardQueryResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

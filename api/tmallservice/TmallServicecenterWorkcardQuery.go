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
func TmallServicecenterWorkcardQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQueryAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardQueryAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

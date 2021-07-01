package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
工单查询接口（面向商家） 
tmall.servicecenter.workcard.querybyseller

查询工单
*/
func TmallServicecenterWorkcardQuerybyseller(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQuerybysellerAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardQuerybysellerAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardQuerybysellerAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

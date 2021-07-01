package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务商分派工人 
tmall.servicecenter.workcard.assignworker

服务商调用该接口分派工人给具体的工单
*/
func TmallServicecenterWorkcardAssignworker(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardAssignworkerAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardAssignworkerAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardAssignworkerAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

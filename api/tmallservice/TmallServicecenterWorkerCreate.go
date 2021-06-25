package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务商工人信息创建 
tmall.servicecenter.worker.create

服务商工人信息创建
*/
func TmallServicecenterWorkerCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerCreateRequest, session string) (*tmallservice.TmallServicecenterWorkerCreateResponse, error) {
    var resp tmallservice.TmallServicecenterWorkerCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

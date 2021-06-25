package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
查询工人列表 
tmall.servicecenter.worker.querypage

服务商查询工人列表
*/
func TmallServicecenterWorkerQuerypage(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerQuerypageRequest, session string) (*tmallservice.TmallServicecenterWorkerQuerypageResponse, error) {
    var resp tmallservice.TmallServicecenterWorkerQuerypageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

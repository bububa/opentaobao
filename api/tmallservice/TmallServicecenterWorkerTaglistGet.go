package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
获取工人标签 
tmall.servicecenter.worker.taglist.get

服务商获取对应工人的标签
*/
func TmallServicecenterWorkerTaglistGet(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerTaglistGetRequest, session string) (*tmallservice.TmallServicecenterWorkerTaglistGetResponse, error) {
    var resp tmallservice.TmallServicecenterWorkerTaglistGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

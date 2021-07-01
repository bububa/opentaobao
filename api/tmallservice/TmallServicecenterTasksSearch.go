package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
查询任务类工单信息 
tmall.servicecenter.tasks.search

查询任务类工单信息
*/
func TmallServicecenterTasksSearch(clt *core.SDKClient, req *tmallservice.TmallServicecenterTasksSearchAPIRequest, session string) (*tmallservice.TmallServicecenterTasksSearchAPIResponse, error) {
    var resp tmallservice.TmallServicecenterTasksSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

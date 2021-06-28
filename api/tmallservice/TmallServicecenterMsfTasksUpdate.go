package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
喵师傅工人任务批量完成接口 
tmall.servicecenter.msf.tasks.update

喵师傅工人任务批量完成接口
*/
func TmallServicecenterMsfTasksUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterMsfTasksUpdateRequest, session string) (*tmallservice.TmallServicecenterMsfTasksUpdateAPIResponse, error) {
    var resp tmallservice.TmallServicecenterMsfTasksUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

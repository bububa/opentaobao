package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
查询任务类工单是否退款 
tmall.servicecenter.task.queryrefund

查询任务类工单是否退款
*/
func TmallServicecenterTaskQueryrefund(clt *core.SDKClient, req *tmallservice.TmallServicecenterTaskQueryrefundRequest, session string) (*tmallservice.TmallServicecenterTaskQueryrefundAPIResponse, error) {
    var resp tmallservice.TmallServicecenterTaskQueryrefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

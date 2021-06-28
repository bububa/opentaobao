package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
工单预约 
tmall.servicecenter.workcard.reserve

服务工单更新通用接口
*/
func TmallServicecenterWorkcardReserve(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReserveRequest, session string) (*tmallservice.TmallServicecenterWorkcardReserveAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardReserveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

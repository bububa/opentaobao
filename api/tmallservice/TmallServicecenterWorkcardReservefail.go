package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
预约失败 
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败
*/
func TmallServicecenterWorkcardReservefail(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReservefailRequest, session string) (*tmallservice.TmallServicecenterWorkcardReservefailResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardReservefailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

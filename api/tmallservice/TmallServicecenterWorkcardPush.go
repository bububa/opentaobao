package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
推送服务工单信息 
tmall.servicecenter.workcard.push

服务商家推送工单信息到天猫。
*/
func TmallServicecenterWorkcardPush(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardPushRequest, session string) (*tmallservice.TmallServicecenterWorkcardPushResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

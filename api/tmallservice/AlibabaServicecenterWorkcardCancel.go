package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务平台工单取消接口 
alibaba.servicecenter.workcard.cancel

取消服务工单
*/
func AlibabaServicecenterWorkcardCancel(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterWorkcardCancelRequest, session string) (*tmallservice.AlibabaServicecenterWorkcardCancelAPIResponse, error) {
    var resp tmallservice.AlibabaServicecenterWorkcardCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
物流订单创建API 
tmall.servicecenter.workcard.expressorder.create

天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
*/
func TmallServicecenterWorkcardExpressorderCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardExpressorderCreateAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardExpressorderCreateAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardExpressorderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

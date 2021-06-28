package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务平台工单创建接口 
alibaba.servicecenter.workcard.create

创建服务平台工单
*/
func AlibabaServicecenterWorkcardCreate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterWorkcardCreateRequest, session string) (*tmallservice.AlibabaServicecenterWorkcardCreateAPIResponse, error) {
    var resp tmallservice.AlibabaServicecenterWorkcardCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

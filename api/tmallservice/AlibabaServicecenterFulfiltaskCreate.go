package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
合单生成核销单 
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单
*/
func AlibabaServicecenterFulfiltaskCreate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterFulfiltaskCreateRequest, session string) (*tmallservice.AlibabaServicecenterFulfiltaskCreateAPIResponse, error) {
    var resp tmallservice.AlibabaServicecenterFulfiltaskCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

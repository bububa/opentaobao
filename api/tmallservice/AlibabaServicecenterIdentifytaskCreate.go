package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
创建核销单 
alibaba.servicecenter.identifytask.create

创建核销单
*/
func AlibabaServicecenterIdentifytaskCreate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterIdentifytaskCreateAPIRequest, session string) (*tmallservice.AlibabaServicecenterIdentifytaskCreateAPIResponse, error) {
    var resp tmallservice.AlibabaServicecenterIdentifytaskCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

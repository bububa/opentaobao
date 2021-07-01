package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
核销单查询 
alibaba.servicecenter.fulfiltask.query

当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
*/
func AlibabaServicecenterFulfiltaskQuery(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterFulfiltaskQueryAPIRequest, session string) (*tmallservice.AlibabaServicecenterFulfiltaskQueryAPIResponse, error) {
    var resp tmallservice.AlibabaServicecenterFulfiltaskQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

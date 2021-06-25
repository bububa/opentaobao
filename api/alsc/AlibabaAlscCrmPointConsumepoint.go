package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
积分抵现 
alibaba.alsc.crm.point.consumepoint

积分抵现
*/
func AlibabaAlscCrmPointConsumepoint(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointConsumepointRequest, session string) (*alsc.AlibabaAlscCrmPointConsumepointResponse, error) {
    var resp alsc.AlibabaAlscCrmPointConsumepointAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

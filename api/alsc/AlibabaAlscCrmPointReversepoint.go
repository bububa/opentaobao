package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
积分消费回退 
alibaba.alsc.crm.point.reversepoint

积分消费回退
*/
func AlibabaAlscCrmPointReversepoint(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointReversepointRequest, session string) (*alsc.AlibabaAlscCrmPointReversepointResponse, error) {
    var resp alsc.AlibabaAlscCrmPointReversepointAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

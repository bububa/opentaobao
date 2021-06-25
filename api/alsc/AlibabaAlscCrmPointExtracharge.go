package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
积分补录 
alibaba.alsc.crm.point.extracharge

积分补录
*/
func AlibabaAlscCrmPointExtracharge(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointExtrachargeRequest, session string) (*alsc.AlibabaAlscCrmPointExtrachargeResponse, error) {
    var resp alsc.AlibabaAlscCrmPointExtrachargeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
积分补扣 
alibaba.alsc.crm.point.extra.consume

积分补扣
*/
func AlibabaAlscCrmPointExtraConsume(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointExtraConsumeAPIRequest, session string) (*alsc.AlibabaAlscCrmPointExtraConsumeAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmPointExtraConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

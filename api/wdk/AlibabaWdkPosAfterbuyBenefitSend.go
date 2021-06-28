package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
生态pos购后发放权益 
alibaba.wdk.pos.afterbuy.benefit.send

生态pos购后发放权益接口开放
*/
func AlibabaWdkPosAfterbuyBenefitSend(clt *core.SDKClient, req *wdk.AlibabaWdkPosAfterbuyBenefitSendRequest, session string) (*wdk.AlibabaWdkPosAfterbuyBenefitSendAPIResponse, error) {
    var resp wdk.AlibabaWdkPosAfterbuyBenefitSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

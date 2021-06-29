package middleclaims

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/middleclaims"
)

/* 
国际化中台服务域接收理赔账单 
alibaba.middle.claimsbill.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
*/
func AlibabaMiddleClaimsbillReceive(clt *core.SDKClient, req *middleclaims.AlibabaMiddleClaimsbillReceiveRequest, session string) (*middleclaims.AlibabaMiddleClaimsbillReceiveAPIResponse, error) {
    var resp middleclaims.AlibabaMiddleClaimsbillReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

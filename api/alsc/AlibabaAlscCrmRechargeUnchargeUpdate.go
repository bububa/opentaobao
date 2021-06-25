package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
充值退款 
alibaba.alsc.crm.recharge.uncharge.update

充值退款
*/
func AlibabaAlscCrmRechargeUnchargeUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeUnchargeUpdateRequest, session string) (*alsc.AlibabaAlscCrmRechargeUnchargeUpdateResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

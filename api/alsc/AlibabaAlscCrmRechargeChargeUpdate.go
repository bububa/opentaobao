package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
储值充值 
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值
*/
func AlibabaAlscCrmRechargeChargeUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeChargeUpdateAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeChargeUpdateAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeChargeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

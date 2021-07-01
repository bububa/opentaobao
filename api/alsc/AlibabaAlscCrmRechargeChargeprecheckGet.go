package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
储值账户充值前校验 
alibaba.alsc.crm.recharge.chargeprecheck.get

储值账户充值前校验接口
*/
func AlibabaAlscCrmRechargeChargeprecheckGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

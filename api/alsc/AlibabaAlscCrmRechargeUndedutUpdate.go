package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
储值消费退款(逆向) 
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口
*/
func AlibabaAlscCrmRechargeUndedutUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeUndedutUpdateRequest, session string) (*alsc.AlibabaAlscCrmRechargeUndedutUpdateAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeUndedutUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

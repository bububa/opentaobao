package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
更新顾客信息 
alibaba.alsc.crm.customer.update

更新顾客信息
*/
func AlibabaAlscCrmCustomerUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerUpdateRequest, session string) (*alsc.AlibabaAlscCrmCustomerUpdateAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCustomerUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

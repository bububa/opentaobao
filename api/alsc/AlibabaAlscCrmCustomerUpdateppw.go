package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
修改支付密码 
alibaba.alsc.crm.customer.updateppw

修改支付密码
*/
func AlibabaAlscCrmCustomerUpdateppw(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerUpdateppwAPIRequest, session string) (*alsc.AlibabaAlscCrmCustomerUpdateppwAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCustomerUpdateppwAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

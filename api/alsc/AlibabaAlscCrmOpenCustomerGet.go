package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询会员资产 
alibaba.alsc.crm.open.customer.get

查询会员身份，资产等
*/
func AlibabaAlscCrmOpenCustomerGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenCustomerGetRequest, session string) (*alsc.AlibabaAlscCrmOpenCustomerGetAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmOpenCustomerGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

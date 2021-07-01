package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
优惠券状态更改 
alibaba.alsc.crm.voucher.status.change

核销优惠券
*/
func AlibabaAlscCrmVoucherStatusChange(clt *core.SDKClient, req *alsc.AlibabaAlscCrmVoucherStatusChangeAPIRequest, session string) (*alsc.AlibabaAlscCrmVoucherStatusChangeAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmVoucherStatusChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

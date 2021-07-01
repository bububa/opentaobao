package alscmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alscmerchant"
)

/* 
外部核销服务 
alibaba.alsc.merchant.ext.ticketcode.use

外部核销服务
*/
func AlibabaAlscMerchantExtTicketcodeUse(clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketcodeUseAPIRequest, session string) (*alscmerchant.AlibabaAlscMerchantExtTicketcodeUseAPIResponse, error) {
    var resp alscmerchant.AlibabaAlscMerchantExtTicketcodeUseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

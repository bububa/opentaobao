package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
确认预约 
alibaba.alihealth.booking.reserve.confirm

确认预约
*/
func AlibabaAlihealthBookingReserveConfirm(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveConfirmRequest, session string) (*alihealth2.AlibabaAlihealthBookingReserveConfirmAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthBookingReserveConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

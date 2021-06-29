package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
取消预约 
alibaba.alihealth.booking.reserve.cancel

消费医疗统一预约平台，ISV取消预约
*/
func AlibabaAlihealthBookingReserveCancel(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveCancelRequest, session string) (*alihealth2.AlibabaAlihealthBookingReserveCancelAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthBookingReserveCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

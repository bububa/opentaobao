package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
确认到店 
alibaba.alihealth.booking.reserve.checkin

消费医疗统一预约平台，ISV 确认到店
*/
func AlibabaAlihealthBookingReserveCheckin(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveCheckinRequest, session string) (*alihealth2.AlibabaAlihealthBookingReserveCheckinAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthBookingReserveCheckinAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

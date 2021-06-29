package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认到店 APIRequest
alibaba.alihealth.booking.reserve.checkin

消费医疗统一预约平台，ISV 确认到店
*/
type AlibabaAlihealthBookingReserveCheckinRequest struct {
    model.Params

    // check_in
    checkIn   *IsvReserveRequest 

}

func NewAlibabaAlihealthBookingReserveCheckinRequest() *AlibabaAlihealthBookingReserveCheckinRequest{
    return &AlibabaAlihealthBookingReserveCheckinRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBookingReserveCheckinRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.checkin"
}

func (r AlibabaAlihealthBookingReserveCheckinRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBookingReserveCheckinRequest) SetCheckIn(checkIn *IsvReserveRequest) error {
    r.checkIn = checkIn
    r.Set("check_in", checkIn)
    return nil
}

func (r AlibabaAlihealthBookingReserveCheckinRequest) GetCheckIn() *IsvReserveRequest {
    return r.checkIn
}


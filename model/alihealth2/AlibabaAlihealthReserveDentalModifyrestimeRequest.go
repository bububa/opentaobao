package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改预约时间 APIRequest
alibaba.alihealth.reserve.dental.modifyrestime

修改预约时间
*/
type AlibabaAlihealthReserveDentalModifyrestimeRequest struct {
    model.Params

    // 预约单ID
    reserveId   int64 

    // 预约时间
    reserveTime   string 

}

func NewAlibabaAlihealthReserveDentalModifyrestimeRequest() *AlibabaAlihealthReserveDentalModifyrestimeRequest{
    return &AlibabaAlihealthReserveDentalModifyrestimeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.modifyrestime"
}

func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthReserveDentalModifyrestimeRequest) SetReserveId(reserveId int64) error {
    r.reserveId = reserveId
    r.Set("reserve_id", reserveId)
    return nil
}

func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetReserveId() int64 {
    return r.reserveId
}

func (r *AlibabaAlihealthReserveDentalModifyrestimeRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetReserveTime() string {
    return r.reserveTime
}


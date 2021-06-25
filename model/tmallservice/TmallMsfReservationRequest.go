package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅服务预约API APIRequest
tmall.msf.reservation

喵师傅预约api
*/
type TmallMsfReservationRequest struct {
    model.Params

    // 预约内容
    reservInfo   *ReservationDTO 

}

func NewTmallMsfReservationRequest() *TmallMsfReservationRequest{
    return &TmallMsfReservationRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMsfReservationRequest) GetApiMethodName() string {
    return "tmall.msf.reservation"
}

func (r TmallMsfReservationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMsfReservationRequest) SetReservInfo(reservInfo *ReservationDTO) error {
    r.reservInfo = reservInfo
    r.Set("reserv_info", reservInfo)
    return nil
}

func (r TmallMsfReservationRequest) GetReservInfo() *ReservationDTO {
    return r.reservInfo
}


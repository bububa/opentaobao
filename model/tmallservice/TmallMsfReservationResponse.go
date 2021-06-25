package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅服务预约API APIResponse
tmall.msf.reservation

喵师傅预约api
*/
type TmallMsfReservationAPIResponse struct {
    model.CommonResponse
    Response *TmallMsfReservationResponse `json:"tmall_msf_reservation_response,omitempty"`
}

type TmallMsfReservationResponse struct {

    // 预约成功,json
    Result   string `json:"result,omitempty"`

}

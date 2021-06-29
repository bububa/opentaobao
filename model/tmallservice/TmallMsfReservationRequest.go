package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅服务预约API API请求
tmall.msf.reservation

喵师傅预约api
*/
type TmallMsfReservationRequest struct {
    model.Params
    // 预约内容
    _reservInfo   *ReservationDTO
}

// 初始化TmallMsfReservationRequest对象
func NewTmallMsfReservationRequest() *TmallMsfReservationRequest{
    return &TmallMsfReservationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMsfReservationRequest) GetApiMethodName() string {
    return "tmall.msf.reservation"
}

// IRequest interface 方法, 获取API参数
func (r TmallMsfReservationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReservInfo Setter
// 预约内容
func (r *TmallMsfReservationRequest) SetReservInfo(_reservInfo *ReservationDTO) error {
    r._reservInfo = _reservInfo
    r.Set("reserv_info", _reservInfo)
    return nil
}

// ReservInfo Getter
func (r TmallMsfReservationRequest) GetReservInfo() *ReservationDTO {
    return r._reservInfo
}

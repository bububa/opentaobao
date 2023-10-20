package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfReservationAPIRequest 喵师傅服务预约API API请求
// tmall.msf.reservation
//
// 喵师傅预约api
type TmallMsfReservationAPIRequest struct {
	model.Params
	// 预约内容
	_reservInfo *ReservationDto
}

// NewTmallMsfReservationRequest 初始化TmallMsfReservationAPIRequest对象
func NewTmallMsfReservationRequest() *TmallMsfReservationAPIRequest {
	return &TmallMsfReservationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMsfReservationAPIRequest) GetApiMethodName() string {
	return "tmall.msf.reservation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMsfReservationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMsfReservationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReservInfo is ReservInfo Setter
// 预约内容
func (r *TmallMsfReservationAPIRequest) SetReservInfo(_reservInfo *ReservationDto) error {
	r._reservInfo = _reservInfo
	r.Set("reserv_info", _reservInfo)
	return nil
}

// GetReservInfo ReservInfo Getter
func (r TmallMsfReservationAPIRequest) GetReservInfo() *ReservationDto {
	return r._reservInfo
}

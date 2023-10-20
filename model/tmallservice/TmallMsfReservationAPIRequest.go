package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmsfreservationAPIRequest 喵师傅服务预约API API请求
// tmall.msf.reservation
//
// 喵师傅预约api
type TmallmsfreservationAPIRequest struct {
	model.Params
	// 预约内容
	_reservInfo *ReservationDto
}

// NewTmallmsfreservationRequest 初始化TmallmsfreservationAPIRequest对象
func NewTmallmsfreservationRequest() *TmallmsfreservationAPIRequest {
	return &TmallmsfreservationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmsfreservationAPIRequest) GetApiMethodName() string {
	return "tmall.msf.reservation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmsfreservationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmsfreservationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReservInfo is ReservInfo Setter
// 预约内容
func (r *TmallmsfreservationAPIRequest) SetReservInfo(_reservInfo *ReservationDto) error {
	r._reservInfo = _reservInfo
	r.Set("reserv_info", _reservInfo)
	return nil
}

// GetReservInfo ReservInfo Getter
func (r TmallmsfreservationAPIRequest) GetReservInfo() *ReservationDto {
	return r._reservInfo
}

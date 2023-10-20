package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixseatinfoqueryAPIRequest 分销商查询座位信息 API请求
// alibaba.damai.maitix.seat.info.query
//
// 分销查询座位文案信息
type AlibabadamaimaitixseatinfoqueryAPIRequest struct {
	model.Params
	// 入参
	_seatQueryParam *SeatQueryParam
}

// NewAlibabadamaimaitixseatinfoqueryRequest 初始化AlibabadamaimaitixseatinfoqueryAPIRequest对象
func NewAlibabadamaimaitixseatinfoqueryRequest() *AlibabadamaimaitixseatinfoqueryAPIRequest {
	return &AlibabadamaimaitixseatinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixseatinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.seat.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixseatinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixseatinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeatQueryParam is SeatQueryParam Setter
// 入参
func (r *AlibabadamaimaitixseatinfoqueryAPIRequest) SetSeatQueryParam(_seatQueryParam *SeatQueryParam) error {
	r._seatQueryParam = _seatQueryParam
	r.Set("seat_query_param", _seatQueryParam)
	return nil
}

// GetSeatQueryParam SeatQueryParam Getter
func (r AlibabadamaimaitixseatinfoqueryAPIRequest) GetSeatQueryParam() *SeatQueryParam {
	return r._seatQueryParam
}

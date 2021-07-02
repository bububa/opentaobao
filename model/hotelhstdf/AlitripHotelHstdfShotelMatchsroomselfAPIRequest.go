package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelMatchsroomselfAPIRequest 匹配标准房型以及卖家房型 API请求
// alitrip.hotel.hstdf.shotel.matchsroomself
//
// 匹配卖家房型以及标准房型，返回匹配结果
type AlitripHotelHstdfShotelMatchsroomselfAPIRequest struct {
	model.Params
	// SroomTypeMatchParam
	_param0 *SroomTypeMatchParam
}

// NewAlitripHotelHstdfShotelMatchsroomselfRequest 初始化AlitripHotelHstdfShotelMatchsroomselfAPIRequest对象
func NewAlitripHotelHstdfShotelMatchsroomselfRequest() *AlitripHotelHstdfShotelMatchsroomselfAPIRequest {
	return &AlitripHotelHstdfShotelMatchsroomselfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.matchsroomself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// SroomTypeMatchParam
func (r *AlitripHotelHstdfShotelMatchsroomselfAPIRequest) SetParam0(_param0 *SroomTypeMatchParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetParam0() *SroomTypeMatchParam {
	return r._param0
}

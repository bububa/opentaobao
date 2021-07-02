package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfHotelroomstaticGetAPIRequest 根据类型查询静态字段 API请求
// alitrip.hotel.hstdf.hotelroomstatic.get
//
// 根据类型查询分页静态字段
type AlitripHotelHstdfHotelroomstaticGetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetHotelRoomStaticParam *GetHotelRoomStaticParam
}

// NewAlitripHotelHstdfHotelroomstaticGetRequest 初始化AlitripHotelHstdfHotelroomstaticGetAPIRequest对象
func NewAlitripHotelHstdfHotelroomstaticGetRequest() *AlitripHotelHstdfHotelroomstaticGetAPIRequest {
	return &AlitripHotelHstdfHotelroomstaticGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfHotelroomstaticGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.hotelroomstatic.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfHotelroomstaticGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamGetHotelRoomStaticParam is ParamGetHotelRoomStaticParam Setter
// 参数封装
func (r *AlitripHotelHstdfHotelroomstaticGetAPIRequest) SetParamGetHotelRoomStaticParam(_paramGetHotelRoomStaticParam *GetHotelRoomStaticParam) error {
	r._paramGetHotelRoomStaticParam = _paramGetHotelRoomStaticParam
	r.Set("param_get_hotel_room_static_param", _paramGetHotelRoomStaticParam)
	return nil
}

// GetParamGetHotelRoomStaticParam ParamGetHotelRoomStaticParam Getter
func (r AlitripHotelHstdfHotelroomstaticGetAPIRequest) GetParamGetHotelRoomStaticParam() *GetHotelRoomStaticParam {
	return r._paramGetHotelRoomStaticParam
}

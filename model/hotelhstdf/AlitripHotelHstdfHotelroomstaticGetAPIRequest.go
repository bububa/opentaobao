package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfhotelroomstaticgetAPIRequest 根据类型查询静态字段 API请求
// alitrip.hotel.hstdf.hotelroomstatic.get
//
// 根据类型查询分页静态字段
type AlitriphotelhstdfhotelroomstaticgetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetHotelRoomStaticParam *GetHotelRoomStaticParam
}

// NewAlitriphotelhstdfhotelroomstaticgetRequest 初始化AlitriphotelhstdfhotelroomstaticgetAPIRequest对象
func NewAlitriphotelhstdfhotelroomstaticgetRequest() *AlitriphotelhstdfhotelroomstaticgetAPIRequest {
	return &AlitriphotelhstdfhotelroomstaticgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfhotelroomstaticgetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.hotelroomstatic.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfhotelroomstaticgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfhotelroomstaticgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamGetHotelRoomStaticParam is ParamGetHotelRoomStaticParam Setter
// 参数封装
func (r *AlitriphotelhstdfhotelroomstaticgetAPIRequest) SetParamGetHotelRoomStaticParam(_paramGetHotelRoomStaticParam *GetHotelRoomStaticParam) error {
	r._paramGetHotelRoomStaticParam = _paramGetHotelRoomStaticParam
	r.Set("param_get_hotel_room_static_param", _paramGetHotelRoomStaticParam)
	return nil
}

// GetParamGetHotelRoomStaticParam ParamGetHotelRoomStaticParam Getter
func (r AlitriphotelhstdfhotelroomstaticgetAPIRequest) GetParamGetHotelRoomStaticParam() *GetHotelRoomStaticParam {
	return r._paramGetHotelRoomStaticParam
}

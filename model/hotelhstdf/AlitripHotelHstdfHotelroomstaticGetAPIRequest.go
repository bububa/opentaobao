package hotelhstdf

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfHotelroomstaticGetAPIRequest) Reset() {
	r._paramGetHotelRoomStaticParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfHotelroomstaticGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.hotelroomstatic.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfHotelroomstaticGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfHotelroomstaticGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripHotelHstdfHotelroomstaticGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfHotelroomstaticGetRequest()
	},
}

// GetAlitripHotelHstdfHotelroomstaticGetRequest 从 sync.Pool 获取 AlitripHotelHstdfHotelroomstaticGetAPIRequest
func GetAlitripHotelHstdfHotelroomstaticGetAPIRequest() *AlitripHotelHstdfHotelroomstaticGetAPIRequest {
	return poolAlitripHotelHstdfHotelroomstaticGetAPIRequest.Get().(*AlitripHotelHstdfHotelroomstaticGetAPIRequest)
}

// ReleaseAlitripHotelHstdfHotelroomstaticGetAPIRequest 将 AlitripHotelHstdfHotelroomstaticGetAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfHotelroomstaticGetAPIRequest(v *AlitripHotelHstdfHotelroomstaticGetAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfHotelroomstaticGetAPIRequest.Put(v)
}

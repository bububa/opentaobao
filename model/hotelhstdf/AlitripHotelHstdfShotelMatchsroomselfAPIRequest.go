package hotelhstdf

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfShotelMatchsroomselfAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.matchsroomself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// SroomTypeMatchParam
func (r *AlitripHotelHstdfShotelMatchsroomselfAPIRequest) SetParam0(_param0 *SroomTypeMatchParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripHotelHstdfShotelMatchsroomselfAPIRequest) GetParam0() *SroomTypeMatchParam {
	return r._param0
}

var poolAlitripHotelHstdfShotelMatchsroomselfAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfShotelMatchsroomselfRequest()
	},
}

// GetAlitripHotelHstdfShotelMatchsroomselfRequest 从 sync.Pool 获取 AlitripHotelHstdfShotelMatchsroomselfAPIRequest
func GetAlitripHotelHstdfShotelMatchsroomselfAPIRequest() *AlitripHotelHstdfShotelMatchsroomselfAPIRequest {
	return poolAlitripHotelHstdfShotelMatchsroomselfAPIRequest.Get().(*AlitripHotelHstdfShotelMatchsroomselfAPIRequest)
}

// ReleaseAlitripHotelHstdfShotelMatchsroomselfAPIRequest 将 AlitripHotelHstdfShotelMatchsroomselfAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfShotelMatchsroomselfAPIRequest(v *AlitripHotelHstdfShotelMatchsroomselfAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfShotelMatchsroomselfAPIRequest.Put(v)
}

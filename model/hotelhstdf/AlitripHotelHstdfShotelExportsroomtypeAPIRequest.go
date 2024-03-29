package hotelhstdf

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelExportsroomtypeAPIRequest 导出一个卖家房型下的所有标准房型 API请求
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
type AlitripHotelHstdfShotelExportsroomtypeAPIRequest struct {
	model.Params
	// 卖家酒店id
	_hid int64
}

// NewAlitripHotelHstdfShotelExportsroomtypeRequest 初始化AlitripHotelHstdfShotelExportsroomtypeAPIRequest对象
func NewAlitripHotelHstdfShotelExportsroomtypeRequest() *AlitripHotelHstdfShotelExportsroomtypeAPIRequest {
	return &AlitripHotelHstdfShotelExportsroomtypeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfShotelExportsroomtypeAPIRequest) Reset() {
	r._hid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelExportsroomtypeAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.exportsroomtype"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelExportsroomtypeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfShotelExportsroomtypeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// 卖家酒店id
func (r *AlitripHotelHstdfShotelExportsroomtypeAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitripHotelHstdfShotelExportsroomtypeAPIRequest) GetHid() int64 {
	return r._hid
}

var poolAlitripHotelHstdfShotelExportsroomtypeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfShotelExportsroomtypeRequest()
	},
}

// GetAlitripHotelHstdfShotelExportsroomtypeRequest 从 sync.Pool 获取 AlitripHotelHstdfShotelExportsroomtypeAPIRequest
func GetAlitripHotelHstdfShotelExportsroomtypeAPIRequest() *AlitripHotelHstdfShotelExportsroomtypeAPIRequest {
	return poolAlitripHotelHstdfShotelExportsroomtypeAPIRequest.Get().(*AlitripHotelHstdfShotelExportsroomtypeAPIRequest)
}

// ReleaseAlitripHotelHstdfShotelExportsroomtypeAPIRequest 将 AlitripHotelHstdfShotelExportsroomtypeAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfShotelExportsroomtypeAPIRequest(v *AlitripHotelHstdfShotelExportsroomtypeAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfShotelExportsroomtypeAPIRequest.Put(v)
}

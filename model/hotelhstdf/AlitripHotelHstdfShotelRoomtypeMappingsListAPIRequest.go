package hotelhstdf

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest 根据HID获取所有卖家房型匹配关系 API请求
// alitrip.hotel.hstdf.shotel.roomtype.mappings.list
//
// 根据HID获取所有卖家房型匹配关系
type AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest struct {
	model.Params
	// HID
	_hid int64
}

// NewAlitripHotelHstdfShotelRoomtypeMappingsListRequest 初始化AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest对象
func NewAlitripHotelHstdfShotelRoomtypeMappingsListRequest() *AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest {
	return &AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) Reset() {
	r._hid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.roomtype.mappings.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// HID
func (r *AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetHid() int64 {
	return r._hid
}

var poolAlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfShotelRoomtypeMappingsListRequest()
	},
}

// GetAlitripHotelHstdfShotelRoomtypeMappingsListRequest 从 sync.Pool 获取 AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest
func GetAlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest() *AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest {
	return poolAlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest.Get().(*AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest)
}

// ReleaseAlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest 将 AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest(v *AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest.Put(v)
}

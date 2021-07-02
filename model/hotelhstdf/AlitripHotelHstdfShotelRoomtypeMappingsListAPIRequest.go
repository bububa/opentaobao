package hotelhstdf

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.roomtype.mappings.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Hid Setter
// HID
func (r *AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest) GetHid() int64 {
	return r._hid
}

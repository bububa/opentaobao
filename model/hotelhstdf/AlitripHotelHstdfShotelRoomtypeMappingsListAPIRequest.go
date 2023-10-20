package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelroomtypemappingslistAPIRequest 根据HID获取所有卖家房型匹配关系 API请求
// alitrip.hotel.hstdf.shotel.roomtype.mappings.list
//
// 根据HID获取所有卖家房型匹配关系
type AlitriphotelhstdfshotelroomtypemappingslistAPIRequest struct {
	model.Params
	// HID
	_hid int64
}

// NewAlitriphotelhstdfshotelroomtypemappingslistRequest 初始化AlitriphotelhstdfshotelroomtypemappingslistAPIRequest对象
func NewAlitriphotelhstdfshotelroomtypemappingslistRequest() *AlitriphotelhstdfshotelroomtypemappingslistAPIRequest {
	return &AlitriphotelhstdfshotelroomtypemappingslistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfshotelroomtypemappingslistAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.roomtype.mappings.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfshotelroomtypemappingslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfshotelroomtypemappingslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// HID
func (r *AlitriphotelhstdfshotelroomtypemappingslistAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitriphotelhstdfshotelroomtypemappingslistAPIRequest) GetHid() int64 {
	return r._hid
}

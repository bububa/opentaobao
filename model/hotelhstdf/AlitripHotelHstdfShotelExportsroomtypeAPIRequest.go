package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelexportsroomtypeAPIRequest 导出一个卖家房型下的所有标准房型 API请求
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
type AlitriphotelhstdfshotelexportsroomtypeAPIRequest struct {
	model.Params
	// 卖家酒店id
	_hid int64
}

// NewAlitriphotelhstdfshotelexportsroomtypeRequest 初始化AlitriphotelhstdfshotelexportsroomtypeAPIRequest对象
func NewAlitriphotelhstdfshotelexportsroomtypeRequest() *AlitriphotelhstdfshotelexportsroomtypeAPIRequest {
	return &AlitriphotelhstdfshotelexportsroomtypeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfshotelexportsroomtypeAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.exportsroomtype"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfshotelexportsroomtypeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfshotelexportsroomtypeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// 卖家酒店id
func (r *AlitriphotelhstdfshotelexportsroomtypeAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitriphotelhstdfshotelexportsroomtypeAPIRequest) GetHid() int64 {
	return r._hid
}

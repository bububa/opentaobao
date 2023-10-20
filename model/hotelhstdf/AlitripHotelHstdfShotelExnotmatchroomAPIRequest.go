package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelexnotmatchroomAPIRequest 导出一个hid下所有未匹配rid的接口 API请求
// alitrip.hotel.hstdf.shotel.exnotmatchroom
//
// 导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
type AlitriphotelhstdfshotelexnotmatchroomAPIRequest struct {
	model.Params
	// 卖家酒店hid
	_hid int64
}

// NewAlitriphotelhstdfshotelexnotmatchroomRequest 初始化AlitriphotelhstdfshotelexnotmatchroomAPIRequest对象
func NewAlitriphotelhstdfshotelexnotmatchroomRequest() *AlitriphotelhstdfshotelexnotmatchroomAPIRequest {
	return &AlitriphotelhstdfshotelexnotmatchroomAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfshotelexnotmatchroomAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.exnotmatchroom"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfshotelexnotmatchroomAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfshotelexnotmatchroomAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// 卖家酒店hid
func (r *AlitriphotelhstdfshotelexnotmatchroomAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitriphotelhstdfshotelexnotmatchroomAPIRequest) GetHid() int64 {
	return r._hid
}

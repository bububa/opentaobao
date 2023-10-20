package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelExnotmatchroomAPIRequest 导出一个hid下所有未匹配rid的接口 API请求
// alitrip.hotel.hstdf.shotel.exnotmatchroom
//
// 导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
type AlitripHotelHstdfShotelExnotmatchroomAPIRequest struct {
	model.Params
	// 卖家酒店hid
	_hid int64
}

// NewAlitripHotelHstdfShotelExnotmatchroomRequest 初始化AlitripHotelHstdfShotelExnotmatchroomAPIRequest对象
func NewAlitripHotelHstdfShotelExnotmatchroomRequest() *AlitripHotelHstdfShotelExnotmatchroomAPIRequest {
	return &AlitripHotelHstdfShotelExnotmatchroomAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelExnotmatchroomAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.exnotmatchroom"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelExnotmatchroomAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfShotelExnotmatchroomAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// 卖家酒店hid
func (r *AlitripHotelHstdfShotelExnotmatchroomAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitripHotelHstdfShotelExnotmatchroomAPIRequest) GetHid() int64 {
	return r._hid
}

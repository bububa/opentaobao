package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformpoirawfeedAPIRequest 存储poi原始数据 API请求
// alitrip.platform.poi.raw.feed
//
// 对接外部数据源，外部数据推送poi数据到飞猪
type AlitripplatformpoirawfeedAPIRequest struct {
	model.Params
	// poi存储参数
	_param0 *TripPoiRawSaveParam
}

// NewAlitripplatformpoirawfeedRequest 初始化AlitripplatformpoirawfeedAPIRequest对象
func NewAlitripplatformpoirawfeedRequest() *AlitripplatformpoirawfeedAPIRequest {
	return &AlitripplatformpoirawfeedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripplatformpoirawfeedAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.feed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripplatformpoirawfeedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripplatformpoirawfeedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// poi存储参数
func (r *AlitripplatformpoirawfeedAPIRequest) SetParam0(_param0 *TripPoiRawSaveParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripplatformpoirawfeedAPIRequest) GetParam0() *TripPoiRawSaveParam {
	return r._param0
}

package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawFeedAPIRequest 存储poi原始数据 API请求
// alitrip.platform.poi.raw.feed
//
// 对接外部数据源，外部数据推送poi数据到飞猪
type AlitripPlatformPoiRawFeedAPIRequest struct {
	model.Params
	// poi存储参数
	_param0 *TripPoiRawSaveParam
}

// NewAlitripPlatformPoiRawFeedRequest 初始化AlitripPlatformPoiRawFeedAPIRequest对象
func NewAlitripPlatformPoiRawFeedRequest() *AlitripPlatformPoiRawFeedAPIRequest {
	return &AlitripPlatformPoiRawFeedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawFeedAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.feed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawFeedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// poi存储参数
func (r *AlitripPlatformPoiRawFeedAPIRequest) SetParam0(_param0 *TripPoiRawSaveParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlitripPlatformPoiRawFeedAPIRequest) GetParam0() *TripPoiRawSaveParam {
	return r._param0
}

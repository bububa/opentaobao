package alitrippoi

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPlatformPoiRawFeedAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawFeedAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.feed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawFeedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPlatformPoiRawFeedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// poi存储参数
func (r *AlitripPlatformPoiRawFeedAPIRequest) SetParam0(_param0 *TripPoiRawSaveParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripPlatformPoiRawFeedAPIRequest) GetParam0() *TripPoiRawSaveParam {
	return r._param0
}

var poolAlitripPlatformPoiRawFeedAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPlatformPoiRawFeedRequest()
	},
}

// GetAlitripPlatformPoiRawFeedRequest 从 sync.Pool 获取 AlitripPlatformPoiRawFeedAPIRequest
func GetAlitripPlatformPoiRawFeedAPIRequest() *AlitripPlatformPoiRawFeedAPIRequest {
	return poolAlitripPlatformPoiRawFeedAPIRequest.Get().(*AlitripPlatformPoiRawFeedAPIRequest)
}

// ReleaseAlitripPlatformPoiRawFeedAPIRequest 将 AlitripPlatformPoiRawFeedAPIRequest 放入 sync.Pool
func ReleaseAlitripPlatformPoiRawFeedAPIRequest(v *AlitripPlatformPoiRawFeedAPIRequest) {
	v.Reset()
	poolAlitripPlatformPoiRawFeedAPIRequest.Put(v)
}

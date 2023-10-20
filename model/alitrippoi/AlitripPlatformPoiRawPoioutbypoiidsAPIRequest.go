package alitrippoi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawPoioutbypoiidsAPIRequest 根据poiId输出飞猪poi数据 API请求
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
type AlitripPlatformPoiRawPoioutbypoiidsAPIRequest struct {
	model.Params
	// 查询参数
	_fliggyPoiidParam *FliggyPoiIdParam
}

// NewAlitripPlatformPoiRawPoioutbypoiidsRequest 初始化AlitripPlatformPoiRawPoioutbypoiidsAPIRequest对象
func NewAlitripPlatformPoiRawPoioutbypoiidsRequest() *AlitripPlatformPoiRawPoioutbypoiidsAPIRequest {
	return &AlitripPlatformPoiRawPoioutbypoiidsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) Reset() {
	r._fliggyPoiidParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.poioutbypoiids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFliggyPoiidParam is FliggyPoiidParam Setter
// 查询参数
func (r *AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) SetFliggyPoiidParam(_fliggyPoiidParam *FliggyPoiIdParam) error {
	r._fliggyPoiidParam = _fliggyPoiidParam
	r.Set("fliggy_poiid_param", _fliggyPoiidParam)
	return nil
}

// GetFliggyPoiidParam FliggyPoiidParam Getter
func (r AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) GetFliggyPoiidParam() *FliggyPoiIdParam {
	return r._fliggyPoiidParam
}

var poolAlitripPlatformPoiRawPoioutbypoiidsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPlatformPoiRawPoioutbypoiidsRequest()
	},
}

// GetAlitripPlatformPoiRawPoioutbypoiidsRequest 从 sync.Pool 获取 AlitripPlatformPoiRawPoioutbypoiidsAPIRequest
func GetAlitripPlatformPoiRawPoioutbypoiidsAPIRequest() *AlitripPlatformPoiRawPoioutbypoiidsAPIRequest {
	return poolAlitripPlatformPoiRawPoioutbypoiidsAPIRequest.Get().(*AlitripPlatformPoiRawPoioutbypoiidsAPIRequest)
}

// ReleaseAlitripPlatformPoiRawPoioutbypoiidsAPIRequest 将 AlitripPlatformPoiRawPoioutbypoiidsAPIRequest 放入 sync.Pool
func ReleaseAlitripPlatformPoiRawPoioutbypoiidsAPIRequest(v *AlitripPlatformPoiRawPoioutbypoiidsAPIRequest) {
	v.Reset()
	poolAlitripPlatformPoiRawPoioutbypoiidsAPIRequest.Put(v)
}

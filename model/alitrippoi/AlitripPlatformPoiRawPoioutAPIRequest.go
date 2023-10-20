package alitrippoi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawPoioutAPIRequest 飞猪poi输出 API请求
// alitrip.platform.poi.raw.poiout
//
// 输出指定城市poi指定信息
type AlitripPlatformPoiRawPoioutAPIRequest struct {
	model.Params
	// 查询参数
	_fliggyPoiOutParam *FliggyPoiOutParam
}

// NewAlitripPlatformPoiRawPoioutRequest 初始化AlitripPlatformPoiRawPoioutAPIRequest对象
func NewAlitripPlatformPoiRawPoioutRequest() *AlitripPlatformPoiRawPoioutAPIRequest {
	return &AlitripPlatformPoiRawPoioutAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPlatformPoiRawPoioutAPIRequest) Reset() {
	r._fliggyPoiOutParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawPoioutAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.poiout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawPoioutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPlatformPoiRawPoioutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFliggyPoiOutParam is FliggyPoiOutParam Setter
// 查询参数
func (r *AlitripPlatformPoiRawPoioutAPIRequest) SetFliggyPoiOutParam(_fliggyPoiOutParam *FliggyPoiOutParam) error {
	r._fliggyPoiOutParam = _fliggyPoiOutParam
	r.Set("fliggy_poi_out_param", _fliggyPoiOutParam)
	return nil
}

// GetFliggyPoiOutParam FliggyPoiOutParam Getter
func (r AlitripPlatformPoiRawPoioutAPIRequest) GetFliggyPoiOutParam() *FliggyPoiOutParam {
	return r._fliggyPoiOutParam
}

var poolAlitripPlatformPoiRawPoioutAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPlatformPoiRawPoioutRequest()
	},
}

// GetAlitripPlatformPoiRawPoioutRequest 从 sync.Pool 获取 AlitripPlatformPoiRawPoioutAPIRequest
func GetAlitripPlatformPoiRawPoioutAPIRequest() *AlitripPlatformPoiRawPoioutAPIRequest {
	return poolAlitripPlatformPoiRawPoioutAPIRequest.Get().(*AlitripPlatformPoiRawPoioutAPIRequest)
}

// ReleaseAlitripPlatformPoiRawPoioutAPIRequest 将 AlitripPlatformPoiRawPoioutAPIRequest 放入 sync.Pool
func ReleaseAlitripPlatformPoiRawPoioutAPIRequest(v *AlitripPlatformPoiRawPoioutAPIRequest) {
	v.Reset()
	poolAlitripPlatformPoiRawPoioutAPIRequest.Put(v)
}

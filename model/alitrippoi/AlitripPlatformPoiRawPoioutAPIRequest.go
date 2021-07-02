package alitrippoi

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawPoioutAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.poiout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawPoioutAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

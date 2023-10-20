package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformpoirawpoioutAPIRequest 飞猪poi输出 API请求
// alitrip.platform.poi.raw.poiout
//
// 输出指定城市poi指定信息
type AlitripplatformpoirawpoioutAPIRequest struct {
	model.Params
	// 查询参数
	_fliggyPoiOutParam *FliggyPoiOutParam
}

// NewAlitripplatformpoirawpoioutRequest 初始化AlitripplatformpoirawpoioutAPIRequest对象
func NewAlitripplatformpoirawpoioutRequest() *AlitripplatformpoirawpoioutAPIRequest {
	return &AlitripplatformpoirawpoioutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripplatformpoirawpoioutAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.poiout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripplatformpoirawpoioutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripplatformpoirawpoioutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFliggyPoiOutParam is FliggyPoiOutParam Setter
// 查询参数
func (r *AlitripplatformpoirawpoioutAPIRequest) SetFliggyPoiOutParam(_fliggyPoiOutParam *FliggyPoiOutParam) error {
	r._fliggyPoiOutParam = _fliggyPoiOutParam
	r.Set("fliggy_poi_out_param", _fliggyPoiOutParam)
	return nil
}

// GetFliggyPoiOutParam FliggyPoiOutParam Getter
func (r AlitripplatformpoirawpoioutAPIRequest) GetFliggyPoiOutParam() *FliggyPoiOutParam {
	return r._fliggyPoiOutParam
}

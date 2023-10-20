package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformpoirawpoioutbypoiidsAPIRequest 根据poiId输出飞猪poi数据 API请求
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
type AlitripplatformpoirawpoioutbypoiidsAPIRequest struct {
	model.Params
	// 查询参数
	_fliggyPoiidParam *FliggyPoiIdParam
}

// NewAlitripplatformpoirawpoioutbypoiidsRequest 初始化AlitripplatformpoirawpoioutbypoiidsAPIRequest对象
func NewAlitripplatformpoirawpoioutbypoiidsRequest() *AlitripplatformpoirawpoioutbypoiidsAPIRequest {
	return &AlitripplatformpoirawpoioutbypoiidsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripplatformpoirawpoioutbypoiidsAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.poioutbypoiids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripplatformpoirawpoioutbypoiidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripplatformpoirawpoioutbypoiidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFliggyPoiidParam is FliggyPoiidParam Setter
// 查询参数
func (r *AlitripplatformpoirawpoioutbypoiidsAPIRequest) SetFliggyPoiidParam(_fliggyPoiidParam *FliggyPoiIdParam) error {
	r._fliggyPoiidParam = _fliggyPoiidParam
	r.Set("fliggy_poiid_param", _fliggyPoiidParam)
	return nil
}

// GetFliggyPoiidParam FliggyPoiidParam Getter
func (r AlitripplatformpoirawpoioutbypoiidsAPIRequest) GetFliggyPoiidParam() *FliggyPoiIdParam {
	return r._fliggyPoiidParam
}

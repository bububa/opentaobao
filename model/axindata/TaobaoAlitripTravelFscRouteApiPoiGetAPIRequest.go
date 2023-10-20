package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest 获取景点（POI）信息 API请求
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
type TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest struct {
	model.Params
	// fscPoiQueryRequest
	_fscPoiQueryRequest *FscPoiQueryRequest
}

// NewTaobaoAlitripTravelFscRouteApiPoiGetRequest 初始化TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiPoiGetRequest() *TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.poi.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscPoiQueryRequest is FscPoiQueryRequest Setter
// fscPoiQueryRequest
func (r *TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest) SetFscPoiQueryRequest(_fscPoiQueryRequest *FscPoiQueryRequest) error {
	r._fscPoiQueryRequest = _fscPoiQueryRequest
	r.Set("fsc_poi_query_request", _fscPoiQueryRequest)
	return nil
}

// GetFscPoiQueryRequest FscPoiQueryRequest Getter
func (r TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest) GetFscPoiQueryRequest() *FscPoiQueryRequest {
	return r._fscPoiQueryRequest
}

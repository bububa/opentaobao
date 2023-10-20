package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapipoigetAPIRequest 获取景点（POI）信息 API请求
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
type TaobaoalitriptravelfscrouteapipoigetAPIRequest struct {
	model.Params
	// fscPoiQueryRequest
	_fscPoiQueryRequest *FscPoiQueryRequest
}

// NewTaobaoalitriptravelfscrouteapipoigetRequest 初始化TaobaoalitriptravelfscrouteapipoigetAPIRequest对象
func NewTaobaoalitriptravelfscrouteapipoigetRequest() *TaobaoalitriptravelfscrouteapipoigetAPIRequest {
	return &TaobaoalitriptravelfscrouteapipoigetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapipoigetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.poi.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapipoigetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapipoigetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscPoiQueryRequest is FscPoiQueryRequest Setter
// fscPoiQueryRequest
func (r *TaobaoalitriptravelfscrouteapipoigetAPIRequest) SetFscPoiQueryRequest(_fscPoiQueryRequest *FscPoiQueryRequest) error {
	r._fscPoiQueryRequest = _fscPoiQueryRequest
	r.Set("fsc_poi_query_request", _fscPoiQueryRequest)
	return nil
}

// GetFscPoiQueryRequest FscPoiQueryRequest Getter
func (r TaobaoalitriptravelfscrouteapipoigetAPIRequest) GetFscPoiQueryRequest() *FscPoiQueryRequest {
	return r._fscPoiQueryRequest
}

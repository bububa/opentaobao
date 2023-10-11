package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest 获取国家城市信息 API请求
// taobao.alitrip.travel.fsc.route.api.division.get
//
// 获取国家城市信息
type TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest struct {
	model.Params
	// fscDivisionQueryRequest
	_fscDivisionQueryRequest *FscDivisionQueryRequest
}

// NewTaobaoAlitripTravelFscRouteApiDivisionGetRequest 初始化TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiDivisionGetRequest() *TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.division.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscDivisionQueryRequest is FscDivisionQueryRequest Setter
// fscDivisionQueryRequest
func (r *TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest) SetFscDivisionQueryRequest(_fscDivisionQueryRequest *FscDivisionQueryRequest) error {
	r._fscDivisionQueryRequest = _fscDivisionQueryRequest
	r.Set("fsc_division_query_request", _fscDivisionQueryRequest)
	return nil
}

// GetFscDivisionQueryRequest FscDivisionQueryRequest Getter
func (r TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest) GetFscDivisionQueryRequest() *FscDivisionQueryRequest {
	return r._fscDivisionQueryRequest
}

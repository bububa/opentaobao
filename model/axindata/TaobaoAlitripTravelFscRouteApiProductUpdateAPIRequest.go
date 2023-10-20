package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest 更新线路产品基本信息 API请求
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
type TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest struct {
	model.Params
	// fscRouteProductUpdateRequest
	_fscRouteProductUpdateRequest *FscRouteProductUpdateRequest
}

// NewTaobaoAlitripTravelFscRouteApiProductUpdateRequest 初始化TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest对象
func NewTaobaoAlitripTravelFscRouteApiProductUpdateRequest() *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest {
	return &TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscRouteProductUpdateRequest is FscRouteProductUpdateRequest Setter
// fscRouteProductUpdateRequest
func (r *TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) SetFscRouteProductUpdateRequest(_fscRouteProductUpdateRequest *FscRouteProductUpdateRequest) error {
	r._fscRouteProductUpdateRequest = _fscRouteProductUpdateRequest
	r.Set("fsc_route_product_update_request", _fscRouteProductUpdateRequest)
	return nil
}

// GetFscRouteProductUpdateRequest FscRouteProductUpdateRequest Getter
func (r TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest) GetFscRouteProductUpdateRequest() *FscRouteProductUpdateRequest {
	return r._fscRouteProductUpdateRequest
}

package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductupdateAPIRequest 更新线路产品基本信息 API请求
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
type TaobaoalitriptravelfscrouteapiproductupdateAPIRequest struct {
	model.Params
	// fscRouteProductUpdateRequest
	_fscRouteProductUpdateRequest *FscRouteProductUpdateRequest
}

// NewTaobaoalitriptravelfscrouteapiproductupdateRequest 初始化TaobaoalitriptravelfscrouteapiproductupdateAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiproductupdateRequest() *TaobaoalitriptravelfscrouteapiproductupdateAPIRequest {
	return &TaobaoalitriptravelfscrouteapiproductupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiproductupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiproductupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiproductupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscRouteProductUpdateRequest is FscRouteProductUpdateRequest Setter
// fscRouteProductUpdateRequest
func (r *TaobaoalitriptravelfscrouteapiproductupdateAPIRequest) SetFscRouteProductUpdateRequest(_fscRouteProductUpdateRequest *FscRouteProductUpdateRequest) error {
	r._fscRouteProductUpdateRequest = _fscRouteProductUpdateRequest
	r.Set("fsc_route_product_update_request", _fscRouteProductUpdateRequest)
	return nil
}

// GetFscRouteProductUpdateRequest FscRouteProductUpdateRequest Getter
func (r TaobaoalitriptravelfscrouteapiproductupdateAPIRequest) GetFscRouteProductUpdateRequest() *FscRouteProductUpdateRequest {
	return r._fscRouteProductUpdateRequest
}

package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductaddAPIRequest 新增线路产品基本信息 API请求
// taobao.alitrip.travel.fsc.route.api.product.add
//
// 新增线路产品基本信息
type TaobaoalitriptravelfscrouteapiproductaddAPIRequest struct {
	model.Params
	// fscRouteProductAddRequest
	_fscRouteProductAddRequest *FscRouteProductAddRequest
}

// NewTaobaoalitriptravelfscrouteapiproductaddRequest 初始化TaobaoalitriptravelfscrouteapiproductaddAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiproductaddRequest() *TaobaoalitriptravelfscrouteapiproductaddAPIRequest {
	return &TaobaoalitriptravelfscrouteapiproductaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiproductaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiproductaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiproductaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscRouteProductAddRequest is FscRouteProductAddRequest Setter
// fscRouteProductAddRequest
func (r *TaobaoalitriptravelfscrouteapiproductaddAPIRequest) SetFscRouteProductAddRequest(_fscRouteProductAddRequest *FscRouteProductAddRequest) error {
	r._fscRouteProductAddRequest = _fscRouteProductAddRequest
	r.Set("fsc_route_product_add_request", _fscRouteProductAddRequest)
	return nil
}

// GetFscRouteProductAddRequest FscRouteProductAddRequest Getter
func (r TaobaoalitriptravelfscrouteapiproductaddAPIRequest) GetFscRouteProductAddRequest() *FscRouteProductAddRequest {
	return r._fscRouteProductAddRequest
}

package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductonlineAPIRequest 产品上线 API请求
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
type TaobaoalitriptravelfscrouteapiproductonlineAPIRequest struct {
	model.Params
	// fscProductOnlineRequest
	_fscProductOnlineRequest *FscProductOnlineRequest
}

// NewTaobaoalitriptravelfscrouteapiproductonlineRequest 初始化TaobaoalitriptravelfscrouteapiproductonlineAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiproductonlineRequest() *TaobaoalitriptravelfscrouteapiproductonlineAPIRequest {
	return &TaobaoalitriptravelfscrouteapiproductonlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiproductonlineAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiproductonlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiproductonlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProductOnlineRequest is FscProductOnlineRequest Setter
// fscProductOnlineRequest
func (r *TaobaoalitriptravelfscrouteapiproductonlineAPIRequest) SetFscProductOnlineRequest(_fscProductOnlineRequest *FscProductOnlineRequest) error {
	r._fscProductOnlineRequest = _fscProductOnlineRequest
	r.Set("fsc_product_online_request", _fscProductOnlineRequest)
	return nil
}

// GetFscProductOnlineRequest FscProductOnlineRequest Getter
func (r TaobaoalitriptravelfscrouteapiproductonlineAPIRequest) GetFscProductOnlineRequest() *FscProductOnlineRequest {
	return r._fscProductOnlineRequest
}

package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductofflineAPIRequest 产品下线 API请求
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
type TaobaoalitriptravelfscrouteapiproductofflineAPIRequest struct {
	model.Params
	// fscProductOfflineRequest
	_fscProductOfflineRequest *FscProductOfflineRequest
}

// NewTaobaoalitriptravelfscrouteapiproductofflineRequest 初始化TaobaoalitriptravelfscrouteapiproductofflineAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiproductofflineRequest() *TaobaoalitriptravelfscrouteapiproductofflineAPIRequest {
	return &TaobaoalitriptravelfscrouteapiproductofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiproductofflineAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiproductofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiproductofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProductOfflineRequest is FscProductOfflineRequest Setter
// fscProductOfflineRequest
func (r *TaobaoalitriptravelfscrouteapiproductofflineAPIRequest) SetFscProductOfflineRequest(_fscProductOfflineRequest *FscProductOfflineRequest) error {
	r._fscProductOfflineRequest = _fscProductOfflineRequest
	r.Set("fsc_product_offline_request", _fscProductOfflineRequest)
	return nil
}

// GetFscProductOfflineRequest FscProductOfflineRequest Getter
func (r TaobaoalitriptravelfscrouteapiproductofflineAPIRequest) GetFscProductOfflineRequest() *FscProductOfflineRequest {
	return r._fscProductOfflineRequest
}

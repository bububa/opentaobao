package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest 获取线路主题 API请求
// taobao.alitrip.travel.fsc.route.api.product.label.get
//
// 获取线路主题
type TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest struct {
	model.Params
	// fscProductLabelQueryRequest
	_fscProductLabelQueryRequest *FscProductLabelQueryRequest
}

// NewTaobaoalitriptravelfscrouteapiproductlabelgetRequest 初始化TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiproductlabelgetRequest() *TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest {
	return &TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.product.label.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProductLabelQueryRequest is FscProductLabelQueryRequest Setter
// fscProductLabelQueryRequest
func (r *TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest) SetFscProductLabelQueryRequest(_fscProductLabelQueryRequest *FscProductLabelQueryRequest) error {
	r._fscProductLabelQueryRequest = _fscProductLabelQueryRequest
	r.Set("fsc_product_label_query_request", _fscProductLabelQueryRequest)
	return nil
}

// GetFscProductLabelQueryRequest FscProductLabelQueryRequest Getter
func (r TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest) GetFscProductLabelQueryRequest() *FscProductLabelQueryRequest {
	return r._fscProductLabelQueryRequest
}

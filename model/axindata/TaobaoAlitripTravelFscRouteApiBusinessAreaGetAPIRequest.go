package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest 获取业务区域 API请求
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
type TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest struct {
	model.Params
	// fscBusinessAreaQueryRequest
	_fscBusinessAreaQueryRequest *FscBusinessAreaQueryRequest
}

// NewTaobaoalitriptravelfscrouteapibusinessareagetRequest 初始化TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest对象
func NewTaobaoalitriptravelfscrouteapibusinessareagetRequest() *TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest {
	return &TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.business.area.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscBusinessAreaQueryRequest is FscBusinessAreaQueryRequest Setter
// fscBusinessAreaQueryRequest
func (r *TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest) SetFscBusinessAreaQueryRequest(_fscBusinessAreaQueryRequest *FscBusinessAreaQueryRequest) error {
	r._fscBusinessAreaQueryRequest = _fscBusinessAreaQueryRequest
	r.Set("fsc_business_area_query_request", _fscBusinessAreaQueryRequest)
	return nil
}

// GetFscBusinessAreaQueryRequest FscBusinessAreaQueryRequest Getter
func (r TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest) GetFscBusinessAreaQueryRequest() *FscBusinessAreaQueryRequest {
	return r._fscBusinessAreaQueryRequest
}

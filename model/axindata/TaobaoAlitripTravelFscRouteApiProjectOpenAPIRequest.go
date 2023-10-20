package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectopenAPIRequest 打开团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
type TaobaoalitriptravelfscrouteapiprojectopenAPIRequest struct {
	model.Params
	// fscProjectOpenRequest
	_fscProjectOpenRequest *FscProjectOpenRequest
}

// NewTaobaoalitriptravelfscrouteapiprojectopenRequest 初始化TaobaoalitriptravelfscrouteapiprojectopenAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiprojectopenRequest() *TaobaoalitriptravelfscrouteapiprojectopenAPIRequest {
	return &TaobaoalitriptravelfscrouteapiprojectopenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiprojectopenAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiprojectopenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiprojectopenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectOpenRequest is FscProjectOpenRequest Setter
// fscProjectOpenRequest
func (r *TaobaoalitriptravelfscrouteapiprojectopenAPIRequest) SetFscProjectOpenRequest(_fscProjectOpenRequest *FscProjectOpenRequest) error {
	r._fscProjectOpenRequest = _fscProjectOpenRequest
	r.Set("fsc_project_open_request", _fscProjectOpenRequest)
	return nil
}

// GetFscProjectOpenRequest FscProjectOpenRequest Getter
func (r TaobaoalitriptravelfscrouteapiprojectopenAPIRequest) GetFscProjectOpenRequest() *FscProjectOpenRequest {
	return r._fscProjectOpenRequest
}

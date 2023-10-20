package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest 关闭团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
type TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest struct {
	model.Params
	// fscProjectCloseRequest
	_fscProjectCloseRequest *FscProjectCloseRequest
}

// NewTaobaoalitriptravelfscrouteapiprojectcloseRequest 初始化TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiprojectcloseRequest() *TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest {
	return &TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectCloseRequest is FscProjectCloseRequest Setter
// fscProjectCloseRequest
func (r *TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest) SetFscProjectCloseRequest(_fscProjectCloseRequest *FscProjectCloseRequest) error {
	r._fscProjectCloseRequest = _fscProjectCloseRequest
	r.Set("fsc_project_close_request", _fscProjectCloseRequest)
	return nil
}

// GetFscProjectCloseRequest FscProjectCloseRequest Getter
func (r TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest) GetFscProjectCloseRequest() *FscProjectCloseRequest {
	return r._fscProjectCloseRequest
}
